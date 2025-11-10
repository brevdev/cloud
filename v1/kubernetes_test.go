package v1

import (
	"encoding/json"
	"errors"
	"regexp"
	"testing"
)

var reWhitespace = regexp.MustCompile(`\s+`)

func TestClusterMarshalJSON(t *testing.T) {
	cluster := &Cluster{
		id:                         "test-id",
		name:                       "test-name",
		refID:                      "test-refID",
		provider:                   "test-provider",
		cloud:                      "test-cloud",
		location:                   "test-location",
		vpcID:                      "test-vpcID",
		subnetIDs:                  []CloudProviderResourceID{"test-subnetID"},
		kubernetesVersion:          "test-kubernetesVersion",
		status:                     ClusterStatusAvailable,
		apiEndpoint:                "test-apiEndpoint",
		clusterCACertificateBase64: "test-clusterCACertificateBase64",
		nodeGroups: []*NodeGroup{
			{
				name:         "test-nodeGroupName",
				refID:        "test-nodeGroupRefID",
				id:           "test-nodeGroupID",
				minNodeCount: 1,
				maxNodeCount: 2,
				instanceType: "test-instanceType",
				diskSize:     NewBytes(10, Gibibyte),
				status:       NodeGroupStatusAvailable,
				tags: Tags{
					"test-nodeGroupTagName": "test-nodeGroupTagValue",
				},
			},
		},
		tags: Tags{
			"test-clusterTagName": "test-clusterTagValue",
		},
	}

	expectedJSON := `{
		"id": "test-id",
		"name": "test-name",
		"refID": "test-refID",
		"provider": "test-provider",
		"cloud": "test-cloud",
		"location": "test-location",
		"vpcID": "test-vpcID",
		"subnetIDs": ["test-subnetID"],
		"kubernetesVersion": "test-kubernetesVersion",
		"status": "available",
		"apiEndpoint": "test-apiEndpoint",
		"clusterCACertificateBase64": "test-clusterCACertificateBase64",
		"nodeGroups": [
			{
				"name": "test-nodeGroupName",
				"refID": "test-nodeGroupRefID",
				"id": "test-nodeGroupID",
				"minNodeCount": 1,
				"maxNodeCount": 2,
				"instanceType": "test-instanceType",
				"diskSize": {
					"value": 10,
					"unit": "GiB"
				},
				"status": "available",
				"tags": {
					"test-nodeGroupTagName": "test-nodeGroupTagValue"
				}
			}
		],
		"tags": {
			"test-clusterTagName": "test-clusterTagValue"
		}
	}`
	expectedJSON = reWhitespace.ReplaceAllString(expectedJSON, "")

	clusterJSON, err := cluster.MarshalJSON()
	if err != nil {
		t.Fatalf("Failed to marshal node group: %v", err)
	}

	if string(clusterJSON) != expectedJSON {
		t.Fatalf("Cluster JSON = %s, want %s", string(clusterJSON), expectedJSON)
	}
}

func TestClusterUnmarshalJSON_InvalidJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr error
	}{
		{name: "invalid status", json: `{"status":"invalid"}`, wantErr: ErrClusterInvalidStatus},
		{name: "refID is required", json: `{"name":"test-name", "status":"available"}`, wantErr: ErrRefIDRequired},
		{name: "name is required", json: `{"refID":"test-refID", "status":"available"}`, wantErr: ErrNameRequired},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var cluster Cluster
			err := json.Unmarshal([]byte(test.json), &cluster)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("Expected error, got %v", err)
			}
		})
	}
}

func TestClusterUnmarshalJSON(t *testing.T) { //nolint:gocyclo,funlen // test ok
	clusterJSON := `{
		"id": "test-id",
		"name": "test-name",
		"refID": "test-refID",
		"provider": "test-provider",
		"cloud": "test-cloud",
		"location": "test-location",
		"vpcID": "test-vpcID",
		"subnetIDs": ["test-subnetID"],
		"kubernetesVersion": "test-kubernetesVersion",
		"status": "available",
		"apiEndpoint": "test-apiEndpoint",
		"clusterCACertificateBase64": "test-clusterCACertificateBase64",
		"nodeGroups": [
			{
				"name": "test-nodeGroupName",
				"refID": "test-nodeGroupRefID",
				"id": "test-nodeGroupID",
				"minNodeCount": 1,
				"maxNodeCount": 2,
				"instanceType": "test-instanceType",
				"diskSize": {
					"value": 10,
					"unit": "GiB"
				},
				"status": "available",
				"tags": {
					"test-nodeGroupTagName": "test-nodeGroupTagValue"
				}
			}
		],
		"tags": {
			"test-clusterTagName": "test-clusterTagValue"
		}
	}`
	var cluster Cluster
	if err := json.Unmarshal([]byte(clusterJSON), &cluster); err != nil {
		t.Fatalf("Failed to unmarshal cluster: %v", err)
	}

	if cluster.id != "test-id" {
		t.Fatalf("Cluster ID = %s, want %s", cluster.id, "test-id")
	}
	if cluster.name != "test-name" {
		t.Fatalf("Cluster Name = %s, want %s", cluster.name, "test-name")
	}
	if cluster.refID != "test-refID" {
		t.Fatalf("Cluster RefID = %s, want %s", cluster.refID, "test-refID")
	}
	if cluster.provider != "test-provider" {
		t.Fatalf("Cluster Provider = %s, want %s", cluster.provider, "test-provider")
	}
	if cluster.cloud != "test-cloud" {
		t.Fatalf("Cluster Cloud = %s, want %s", cluster.cloud, "test-cloud")
	}
	if cluster.location != "test-location" {
		t.Fatalf("Cluster Location = %s, want %s", cluster.location, "test-location")
	}
	if cluster.vpcID != "test-vpcID" {
		t.Fatalf("Cluster VPCID = %s, want %s", cluster.vpcID, "test-vpcID")
	}
	if cluster.kubernetesVersion != "test-kubernetesVersion" {
		t.Fatalf("Cluster KubernetesVersion = %s, want %s", cluster.kubernetesVersion, "test-kubernetesVersion")
	}
	if cluster.status != ClusterStatusAvailable {
		t.Fatalf("Cluster Status = %s, want %s", cluster.status, "available")
	}
	if cluster.apiEndpoint != "test-apiEndpoint" {
		t.Fatalf("Cluster APIEndpoint = %s, want %s", cluster.apiEndpoint, "test-apiEndpoint")
	}
	if cluster.clusterCACertificateBase64 != "test-clusterCACertificateBase64" {
		t.Fatalf("Cluster ClusterCACertificateBase64 = %s, want %s", cluster.clusterCACertificateBase64, "test-clusterCACertificateBase64")
	}
	if len(cluster.nodeGroups) != 1 {
		t.Fatalf("Cluster NodeGroups = %d, want %d", len(cluster.nodeGroups), 1)
	}
	if cluster.nodeGroups[0].name != "test-nodeGroupName" {
		t.Fatalf("Cluster NodeGroup Name = %s, want %s", cluster.nodeGroups[0].name, "test-nodeGroupName")
	}
	if cluster.nodeGroups[0].refID != "test-nodeGroupRefID" {
		t.Fatalf("Cluster NodeGroup RefID = %s, want %s", cluster.nodeGroups[0].refID, "test-nodeGroupRefID")
	}
	if cluster.nodeGroups[0].id != "test-nodeGroupID" {
		t.Fatalf("Cluster NodeGroup ID = %s, want %s", cluster.nodeGroups[0].id, "test-nodeGroupID")
	}
	if cluster.nodeGroups[0].minNodeCount != 1 {
		t.Fatalf("Cluster NodeGroup MinNodeCount = %d, want %d", cluster.nodeGroups[0].minNodeCount, 1)
	}
	if cluster.nodeGroups[0].maxNodeCount != 2 {
		t.Fatalf("Cluster NodeGroup MaxNodeCount = %d, want %d", cluster.nodeGroups[0].maxNodeCount, 2)
	}
	if cluster.nodeGroups[0].instanceType != "test-instanceType" {
		t.Fatalf("Cluster NodeGroup InstanceType = %s, want %s", cluster.nodeGroups[0].instanceType, "test-instanceType")
	}
	if !cluster.nodeGroups[0].diskSize.Equal(NewBytes(10, Gibibyte)) {
		t.Fatalf("Cluster NodeGroup DiskSize = %s, want %s", cluster.nodeGroups[0].diskSize, "10 GiB")
	}
	if len(cluster.nodeGroups[0].tags) != 1 {
		t.Fatalf("Cluster NodeGroup Tags = %d, want %d", len(cluster.nodeGroups[0].tags), 1)
	}
	if cluster.nodeGroups[0].tags["test-nodeGroupTagName"] != "test-nodeGroupTagValue" {
		t.Fatalf("Cluster NodeGroup Tag = %s, want %s", cluster.nodeGroups[0].tags["test-nodeGroupTagName"], "test-nodeGroupTagValue")
	}
	if len(cluster.tags) != 1 {
		t.Fatalf("Cluster Tags = %d, want %d", len(cluster.tags), 1)
	}
	if cluster.tags["test-clusterTagName"] != "test-clusterTagValue" {
		t.Fatalf("Cluster Tag = %s, want %s", cluster.tags["test-clusterTagName"], "test-clusterTagValue")
	}
}

func TestNodeGroupMarshalJSON(t *testing.T) {
	nodeGroup := &NodeGroup{
		name:         "test-nodeGroupName",
		refID:        "test-nodeGroupRefID",
		id:           "test-nodeGroupID",
		minNodeCount: 1,
		maxNodeCount: 2,
		instanceType: "test-instanceType",
		diskSize:     NewBytes(10, Gibibyte),
		status:       NodeGroupStatusAvailable,
		tags: Tags{
			"test-tagName": "test-tagValue",
		},
	}
	expectedJSON := `{
		"name": "test-nodeGroupName",
		"refID": "test-nodeGroupRefID",
		"id": "test-nodeGroupID",
		"minNodeCount": 1,
		"maxNodeCount": 2,
		"instanceType": "test-instanceType",
		"diskSize": {
			"value": 10,
			"unit": "GiB"
		},
		"status": "available",
		"tags": {
			"test-tagName": "test-tagValue"
		}
	}`
	expectedJSON = reWhitespace.ReplaceAllString(expectedJSON, "")

	nodeGroupJSON, err := nodeGroup.MarshalJSON()
	if err != nil {
		t.Fatalf("Failed to marshal node group: %v", err)
	}

	if string(nodeGroupJSON) != expectedJSON {
		t.Fatalf("NodeGroup JSON = %s, want %s", string(nodeGroupJSON), expectedJSON)
	}
}

func TestNodeGroupUnmarshalJSON(t *testing.T) {
	nodeGroupJSON := `{
		"name": "test-nodeGroupName",
		"refID": "test-nodeGroupRefID",
		"id": "test-nodeGroupID",
		"minNodeCount": 1,
		"maxNodeCount": 2,
		"instanceType": "test-instanceType",
		"diskSize": {
			"value": 10,
			"unit": "GiB"
		},
		"status": "available",
		"tags": {
			"test-tagName": "test-tagValue"
		}
	}`
	var nodeGroup NodeGroup
	if err := json.Unmarshal([]byte(nodeGroupJSON), &nodeGroup); err != nil {
		t.Fatalf("Failed to unmarshal node group: %v", err)
	}

	if nodeGroup.name != "test-nodeGroupName" {
		t.Fatalf("NodeGroup Name = %s, want %s", nodeGroup.name, "test-nodeGroupName")
	}
	if nodeGroup.refID != "test-nodeGroupRefID" {
		t.Fatalf("NodeGroup RefID = %s, want %s", nodeGroup.refID, "test-nodeGroupRefID")
	}
	if nodeGroup.id != "test-nodeGroupID" {
		t.Fatalf("NodeGroup ID = %s, want %s", nodeGroup.id, "test-nodeGroupID")
	}
	if nodeGroup.minNodeCount != 1 {
		t.Fatalf("NodeGroup MinNodeCount = %d, want %d", nodeGroup.minNodeCount, 1)
	}
	if nodeGroup.maxNodeCount != 2 {
		t.Fatalf("NodeGroup MaxNodeCount = %d, want %d", nodeGroup.maxNodeCount, 2)
	}
	if nodeGroup.instanceType != "test-instanceType" {
		t.Fatalf("NodeGroup InstanceType = %s, want %s", nodeGroup.instanceType, "test-instanceType")
	}
	if nodeGroup.diskSize != NewBytes(10, Gibibyte) {
		t.Fatalf("NodeGroup DiskSize = %s, want %s", nodeGroup.diskSize, "10 GiB")
	}
	if nodeGroup.status != NodeGroupStatusAvailable {
		t.Fatalf("NodeGroup Status = %s, want %s", nodeGroup.status, "available")
	}
	if len(nodeGroup.tags) != 1 {
		t.Fatalf("NodeGroup Tags = %d, want %d", len(nodeGroup.tags), 1)
	}
	if nodeGroup.tags["test-tagName"] != "test-tagValue" {
		t.Fatalf("NodeGroup Tag = %s, want %s", nodeGroup.tags["test-tagName"], "test-tagValue")
	}
}

func TestNodeGroupUnmarshalJSON_InvalidJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr error
	}{
		{name: "invalid status", json: `{"status":"invalid"}`, wantErr: ErrNodeGroupInvalidStatus},
		{name: "refID is required", json: `{"name":"test-nodeGroupName", "status":"available"}`, wantErr: ErrRefIDRequired},
		{name: "name is required", json: `{"refID":"test-nodeGroupRefID", "status":"available"}`, wantErr: ErrNameRequired},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var nodeGroup NodeGroup
			err := json.Unmarshal([]byte(test.json), &nodeGroup)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("Expected error, got %v", err)
			}
		})
	}
}
