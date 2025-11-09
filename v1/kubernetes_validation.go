package v1

import (
	"context"
	"fmt"
)

// ValidateCreateKubernetesCluster validates that the CreateCluster functionality works correctly.
func ValidateCreateKubernetesCluster(ctx context.Context, client CloudMaintainKubernetes, attrs CreateClusterArgs) (*Cluster, error) {
	cluster, err := client.CreateCluster(ctx, attrs)
	if err != nil {
		return nil, err
	}

	if cluster.GetName() != attrs.Name {
		return nil, fmt.Errorf("cluster name does not match create args: '%s' != '%s'", cluster.GetName(), attrs.Name)
	}
	if cluster.GetRefID() != attrs.RefID {
		return nil, fmt.Errorf("cluster refID does not match create args: '%s' != '%s'", cluster.GetRefID(), attrs.RefID)
	}
	if cluster.GetKubernetesVersion() != attrs.KubernetesVersion {
		return nil, fmt.Errorf("cluster KubernetesVersion does not match create args: '%s' != '%s'", cluster.GetKubernetesVersion(), attrs.KubernetesVersion)
	}
	if cluster.GetVPCID() != attrs.VPCID {
		return nil, fmt.Errorf("cluster VPCID does not match create args: '%s' != '%s'", cluster.GetVPCID(), attrs.VPCID)
	}
	if len(cluster.GetSubnetIDs()) != len(attrs.SubnetIDs) {
		return nil, fmt.Errorf("cluster subnetIDs does not match create args: '%d' != '%d'", len(cluster.GetSubnetIDs()), len(attrs.SubnetIDs))
	}
	for key, value := range attrs.Tags {
		tagValue, ok := cluster.GetTags()[key]
		if !ok {
			return nil, fmt.Errorf("cluster tag does not match create args: '%s' not found", key)
		}
		if tagValue != value {
			return nil, fmt.Errorf("cluster tag does not match create args: '%s' != '%s'", key, value)
		}
	}
	return cluster, nil
}

// ValidateGetKubernetesCluster validates that the GetCluster functionality works correctly.
func ValidateGetKubernetesCluster(ctx context.Context, client CloudMaintainKubernetes, attrs GetClusterArgs) (*Cluster, error) {
	cluster, err := client.GetCluster(ctx, attrs)
	if err != nil {
		return nil, err
	}

	if cluster.GetID() != attrs.ID {
		return nil, fmt.Errorf("cluster ID does not match get args: '%s' != '%s'", cluster.GetID(), attrs.ID)
	}

	return cluster, nil
}

// ValidateSetKubernetesClusterUser validates that the SetClusterUser functionality works correctly.
func ValidateSetKubernetesClusterUser(ctx context.Context, client CloudMaintainKubernetes, attrs SetClusterUserArgs) (*ClusterUser, error) {
	clusterUser, err := client.SetClusterUser(ctx, attrs)
	if err != nil {
		return nil, err
	}
	return clusterUser, nil
}

// ValidateCreateKubernetesNodeGroup validates that the CreateNodeGroup functionality works correctly.
func ValidateCreateKubernetesNodeGroup(ctx context.Context, client CloudMaintainKubernetes, attrs CreateNodeGroupArgs) (*NodeGroup, error) {
	nodeGroup, err := client.CreateNodeGroup(ctx, attrs)
	if err != nil {
		return nil, err
	}

	if nodeGroup.GetName() != attrs.Name {
		return nil, fmt.Errorf("node group name does not match create args: '%s' != '%s'", nodeGroup.GetName(), attrs.Name)
	}
	if nodeGroup.GetRefID() != attrs.RefID {
		return nil, fmt.Errorf("node group refID does not match create args: '%s' != '%s'", nodeGroup.GetRefID(), attrs.RefID)
	}
	if nodeGroup.GetMinNodeCount() != attrs.MinNodeCount {
		return nil, fmt.Errorf("node group minNodeCount does not match create args: '%d' != '%d'", nodeGroup.GetMinNodeCount(), attrs.MinNodeCount)
	}
	if nodeGroup.GetMaxNodeCount() != attrs.MaxNodeCount {
		return nil, fmt.Errorf("node group maxNodeCount does not match create args: '%d' != '%d'", nodeGroup.GetMaxNodeCount(), attrs.MaxNodeCount)
	}
	if nodeGroup.GetInstanceType() != attrs.InstanceType {
		return nil, fmt.Errorf("node group instanceType does not match create args: '%s' != '%s'", nodeGroup.GetInstanceType(), attrs.InstanceType)
	}
	if nodeGroup.GetDiskSize() != attrs.DiskSize {
		return nil, fmt.Errorf("node group diskSize does not match create args: '%s' != '%s'", nodeGroup.GetDiskSize(), attrs.DiskSize)
	}

	return nodeGroup, nil
}

// ValidateClusterNodeGroups validates that the GetCluster functionality works correctly.
func ValidateClusterNodeGroups(ctx context.Context, client CloudMaintainKubernetes, attrs GetClusterArgs, nodeGroup NodeGroup) error {
	cluster, err := client.GetCluster(ctx, attrs)
	if err != nil {
		return err
	}

	if len(cluster.GetNodeGroups()) != 1 {
		return fmt.Errorf("cluster node groups does not match create args: '%d' != '%d'", len(cluster.GetNodeGroups()), 1)
	}

	clusterNodeGroup := cluster.GetNodeGroups()[0]
	if clusterNodeGroup.GetID() != nodeGroup.GetID() {
		return fmt.Errorf("cluster node group ID does not match create args: '%s' != '%s'", clusterNodeGroup.GetID(), nodeGroup.GetID())
	}
	if clusterNodeGroup.GetName() != nodeGroup.GetName() {
		return fmt.Errorf("cluster node group name does not match create args: '%s' != '%s'", clusterNodeGroup.GetName(), nodeGroup.GetName())
	}
	if clusterNodeGroup.GetRefID() != nodeGroup.GetRefID() {
		return fmt.Errorf("cluster node group refID does not match create args: '%s' != '%s'", clusterNodeGroup.GetRefID(), nodeGroup.GetRefID())
	}
	if clusterNodeGroup.GetMinNodeCount() != nodeGroup.GetMinNodeCount() {
		return fmt.Errorf("cluster node group minNodeCount does not match create args: '%d' != '%d'", clusterNodeGroup.GetMinNodeCount(), nodeGroup.GetMinNodeCount())
	}
	if clusterNodeGroup.GetMaxNodeCount() != nodeGroup.GetMaxNodeCount() {
		return fmt.Errorf("cluster node group maxNodeCount does not match create args: '%d' != '%d'", clusterNodeGroup.GetMaxNodeCount(), nodeGroup.GetMaxNodeCount())
	}
	if clusterNodeGroup.GetInstanceType() != nodeGroup.GetInstanceType() {
		return fmt.Errorf("cluster node group instanceType does not match create args: '%s' != '%s'", clusterNodeGroup.GetInstanceType(), nodeGroup.GetInstanceType())
	}
	if clusterNodeGroup.GetDiskSize() != nodeGroup.GetDiskSize() {
		return fmt.Errorf("cluster node group diskSize does not match create args: '%s' != '%s'", clusterNodeGroup.GetDiskSize(), nodeGroup.GetDiskSize())
	}
	for key, value := range nodeGroup.GetTags() {
		tagValue, ok := clusterNodeGroup.GetTags()[key]
		if !ok {
			return fmt.Errorf("cluster node group tag does not match create args: '%s' not found", key)
		}
		if tagValue != value {
			return fmt.Errorf("cluster node group tag does not match create args: '%s' != '%s'", key, value)
		}
	}

	return nil
}

// ValidateModifyKubernetesNodeGroup validates that the ModifyNodeGroup functionality works correctly.
func ValidateModifyKubernetesNodeGroup(ctx context.Context, client CloudMaintainKubernetes, attrs ModifyNodeGroupArgs) error {
	err := client.ModifyNodeGroup(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}

// ValidateDeleteKubernetesNodeGroup validates that the DeleteNodeGroup functionality works correctly.
func ValidateDeleteKubernetesNodeGroup(ctx context.Context, client CloudMaintainKubernetes, attrs DeleteNodeGroupArgs) error {
	err := client.DeleteNodeGroup(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}

// ValidateDeleteKubernetesCluster validates that the DeleteCluster functionality works correctly.
func ValidateDeleteKubernetesCluster(ctx context.Context, client CloudMaintainKubernetes, attrs DeleteClusterArgs) error {
	err := client.DeleteCluster(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}
