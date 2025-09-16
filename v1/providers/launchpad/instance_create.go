package v1

import (
	"context"
	"net"

	"github.com/brevdev/cloud/internal/collections"
	"github.com/brevdev/cloud/internal/errors"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	brevExperienceID    = "43123766-35ec-4eb4-a5ba-3f2945228445"
	minimalExperienceID = "a5d93f56-bbdb-44db-a1ae-6b1ad7d3c6df"
)

func (c *LaunchpadClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	createAttrs := launchpadCreateAttrs(attrs)
	err := createAttrs.validate()
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	ipAllowlist := c.getLaunchpadIPAllowlist(ctx, attrs.FirewallRules)
	instanceTypeInfo, err := getInstanceTypeInfo(createAttrs.InstanceType)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	var region openapi.NullableString
	if createAttrs.Location != "" {
		region = *openapi.NewNullableString(collections.Ptr(createAttrs.Location))
	} else {
		region = *openapi.NewNullableString(nil)
	}

	publicKey := openapi.NewNullableString(collections.Ptr(attrs.PublicKey))
	createAttrs.Tags = createAttrs.generateTags(c.GetReferenceID())

	deployment := openapi.Deployment{
		ExperienceId:   collections.Ptr(instanceTypeToLaunchpadExperience(instanceTypeInfo)), // TODO: necessary for now, but ideally we would use a Brev reservation/pool
		ExpiresAt:      *openapi.NewNullableTime(nil),
		GpuModel:       *openapi.NewNullableString(collections.Ptr(instanceTypeInfo.gpuName)),
		GpuCount:       *openapi.NewNullableInt32(collections.Ptr(instanceTypeInfo.gpuCount)),
		IpAllowlist:    ipAllowlist,
		OrgName:        "brev", // TODO: use org name
		ProviderName:   *openapi.NewNullableString(collections.Ptr(instanceTypeInfo.cloud)),
		PublicKey:      *publicKey,
		RequestId:      collections.Ptr(attrs.RefID),
		RequesterEmail: "brev@nvidia.com", // TODO: use user email address
		RequesterName:  "brev.dev",        // TODO: use user name
		Tags:           createAttrs.Tags,
		Region:         region,
	}

	// The workshop ID should only be set if it was determined to be not empty
	if instanceTypeInfo.workshopID != "" {
		deployment.Workshop = *openapi.NewNullableBool(collections.Ptr(true))
		deployment.WorkshopId = *openapi.NewNullableString(collections.Ptr(instanceTypeInfo.workshopID))
	} else {
		deployment.Workshop = *openapi.NewNullableBool(nil)
		deployment.WorkshopId = *openapi.NewNullableString(nil)
	}

	createDeployment, resp, err := c.client.CatalogDeploymentsAPI.CatalogDeploymentsCreate(c.makeAuthContext(ctx)).
		Deployment(deployment).
		Execute()
	if resp != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				c.logger.Error(ctx, errors.Wrap(err, "cloud/launchpad error closing response body"))
			}
		}()
	}

	if err != nil {
		if resp == nil {
			return nil, errors.WrapAndTrace(err)
		}

		lpErr := c.handleLaunchpadAPIErr(ctx, resp, err)
		if !wasCreateDeploymentRequestAlreadyMade(lpErr) { // handleLaunchpadAPIErr casts error to include error message
			return nil, errors.WrapAndTrace(lpErr)
		}

		// Error was "create request already made", so we can continue
	}

	inst, err := c.GetInstance(ctx, v1.CloudProviderInstanceID(createDeployment.Id))
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return inst, nil
}

func instanceTypeToLaunchpadExperience(instanceTypeInfo instanceTypeInfo) string {
	if instanceTypeInfo.cloud == "nebius" {
		return minimalExperienceID
	}
	return brevExperienceID
}

type launchpadCreateAttrs v1.CreateInstanceAttrs

func (l launchpadCreateAttrs) validate() error {
	err := validation.ValidateStruct(
		&l,
		validation.Field(&l.RefID, validation.Required),
		validation.Field(&l.Name, validation.Required),
		validation.Field(&l.PublicKey, validation.Required),
		validation.Field(&l.InstanceType, validation.Required),
	)
	if err != nil {
		return errors.WrapAndTrace(err)
	}
	return nil
}

func (l launchpadCreateAttrs) generateTags(cloudCredRefID string) v1.Tags {
	tags := map[string]string{}

	for key, value := range l.Tags {
		tags[key] = value
	}
	tags["RefID"] = l.RefID
	tags["Name"] = l.Name
	tags["CloudCredRefID"] = cloudCredRefID
	return tags
}

func (c *LaunchpadClient) getLaunchpadIPAllowlist(ctx context.Context, firewallRules v1.FirewallRules) []string {
	if len(firewallRules.EgressRules) > 0 {
		c.logger.Info(ctx, "cloud/launchpad egress rules not supported", v1.LogField("egressRules", firewallRules.EgressRules))
	}
	ips := []string{}
	for _, rule := range firewallRules.IngressRules {
		for _, cidr := range rule.IPRanges {
			_, ipNet, err := net.ParseCIDR(cidr)
			if err != nil {
				c.logger.Error(ctx, errors.Wrap(err, "cloud/launchpad error parsing cidr"), v1.LogField("cidr", cidr))
				continue
			}
			ones, bits := ipNet.Mask.Size()
			// check if mask is 32 bits
			if ones == bits {
				ips = append(ips, ipNet.IP.String())
			} else {
				c.logger.Error(ctx, errors.New("cloud/launchpad only accept singular IPs"), v1.LogField("cidr", cidr))
			}
		}
	}

	return ips
}

func wasCreateDeploymentRequestAlreadyMade(err error) bool {
	return errors.ErrorContains(err, "request_id already exist")
}
