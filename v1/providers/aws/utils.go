package v1

import (
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

const (
	tagBrevRefID      = "brev-ref-id"
	tagBrevVPCID      = "brev-vpc-id"
	tagBrevClusterID  = "brev-cluster-id"
	tagBrevSubnetType = "brev-subnet-type"
	tagBrevCloudSDK   = "brev-cloud-sdk"
	tagCreatedBy      = "CreatedBy"
	tagName           = "Name"
)

type awsClient struct {
	eksClient *eks.Client
	iamClient *iam.Client
}
