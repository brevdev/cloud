package v1

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	sfcnodes "github.com/sfcompute/nodes-go"
	"github.com/sfcompute/nodes-go/packages/param"
)

// define function to convert string to b64
func toBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// define function to add ssh key to cloud init
func sshKeyCloudInit(sshKey string) string {
	return toBase64(fmt.Sprintf("#cloud-config\nssh_authorized_keys:\n  - %s", sshKey))
}

func (c *SFCClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	resp, err := c.client.Nodes.New(ctx, sfcnodes.NodeNewParams{
		CreateNodesRequest: sfcnodes.CreateNodesRequestParam{
			DesiredCount:        1,
			MaxPricePerNodeHour: 1000,
			Zone:                attrs.Location,
			ImageID:             param.Opt[string]{Value: attrs.ImageID},                    //this needs to point to a valid image
			CloudInitUserData:   param.Opt[string]{Value: sshKeyCloudInit(attrs.PublicKey)}, // encode ssh key to b64-wrapped cloud-init script
		},
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no nodes returned")
	}
	node := resp.Data[0]

	inst := &v1.Instance{
		Name:           attrs.Name,
		RefID:          attrs.RefID,
		CloudCredRefID: c.refID,
		CloudID:        v1.CloudProviderInstanceID(node.ID), // SFC ID
		ImageID:        attrs.ImageID,
		InstanceType:   attrs.InstanceType,
		Location:       attrs.Location,
		CreatedAt:      time.Now(),
		Status:         v1.Status{LifecycleStatus: v1.LifecycleStatusPending}, // or map from SDK status
		InstanceTypeID: v1.InstanceTypeID(node.GPUType),
	}

	return inst, nil
}

func (c *SFCClient) GetInstance(ctx context.Context, id v1.CloudProviderInstanceID) (*v1.Instance, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *SFCClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *SFCClient) TerminateInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return fmt.Errorf("not implemented")
}

// Optional if supported:
func (c *SFCClient) RebootInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return fmt.Errorf("not implemented")
}
func (c *SFCClient) StopInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return fmt.Errorf("not implemented")
}
func (c *SFCClient) StartInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return fmt.Errorf("not implemented")
}

// Merge strategies (pass-through is acceptable baseline).
func (c *SFCClient) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}
func (c *SFCClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newIt v1.InstanceType) v1.InstanceType {
	return newIt
}
