//go:build scripts
// +build scripts

package scripts

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/ssh"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/google/uuid"
)

func TestGetInstanceTypes(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	credential := NewSFCCredential("validation-test", apiKey)
	client, err := credential.MakeClient(context.Background(), "richmond")
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	locations, err := client.GetLocations(context.Background(), v1.GetLocationsArgs{
		IncludeUnavailable: true,
	})
	if err != nil {
		t.Fatalf("failed to get locations: %v", err)
	}

	t.Logf("locations: %v", locations)

	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{
		Locations: v1.LocationsFilter{"all"},
	})
	if err != nil {
		t.Fatalf("failed to get instance types: %v", err)
	}

	t.Logf("instance types: %v", instanceTypes)
}

func TestCreateInstance(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	credential := NewSFCCredential("validation-test", apiKey)
	client, err := credential.MakeClient(context.Background(), "richmond")
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	id := uuid.New().String()

	instance, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		Name:         fmt.Print("test-%s", id),
		RefID:        id,
		PublicKey:    ssh.GetTestPublicKey(),
		InstanceType: "h100",
		Location:     "richmond",
	})
	if err != nil {
		t.Fatalf("failed to create instance: %v", err)
	}

	t.Logf("instance: %v", instance)
}

func TestGetInstance(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	credential := NewSFCCredential("validation-test", apiKey)
	client, err := credential.MakeClient(context.Background(), "")
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	instance, err := client.GetInstance(context.Background(), "6c7a3ade-1e59-4e04-af6e-365046995a81_test")
	if err != nil {
		t.Fatalf("failed to get instance: %v", err)
	}

	t.Logf("instance: %v", instance)

	// status
	t.Logf("status: %v", instance.Status)

	// ssh details
	t.Logf("ssh details: %v,%v,%v", instance.SSHUser, instance.SSHPort, instance.PublicIP)
}

func TestSSHInstance(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	credential := NewSFCCredential("validation-test", apiKey)
	client, err := credential.MakeClient(context.Background(), "")
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	instance, err := client.GetInstance(context.Background(), "6c7a3ade-1e59-4e04-af6e-365046995a81_test")
	if err != nil {
		t.Fatalf("failed to get instance: %v", err)
	}

	t.Logf("instance: %v", instance)

	// ssh details
	t.Logf("ssh details: %v,%v,%v", instance.SSHUser, instance.SSHPort, instance.PublicIP)

	// ssh to instance
	err = ssh.WaitForSSH(context.Background(), ssh.ConnectionConfig{
		User:     "root",
		HostPort: fmt.Sprintf("%s:%d", instance.PublicIP, instance.SSHPort),
		PrivKey:  ssh.GetTestPrivateKey(),
	}, ssh.WaitForSSHOptions{
		Timeout: 10 * time.Second,
	})
	if err != nil {
		t.Fatalf("failed to wait for SSH: %v", err)
	}

	t.Logf("SSH connection validated successfully for %s@%s:%d", instance.SSHUser, instance.PublicIP, instance.SSHPort)
}

func TestListInstances(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	credential := NewSFCCredential("validation-test", apiKey)
	client, err := credential.MakeClient(context.Background(), "")
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	instances, err := client.ListInstances(context.Background(), v1.ListInstancesArgs{
		TagFilters: map[string][]string{
			"dev-plane-managedBy": {"dev-plane"},
		},
		Locations: v1.All,
	})
	if err != nil {
		t.Fatalf("failed to list instances: %v", err)
	}

	t.Logf("instances: %v", instances)
}

func TestTerminateInstance(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	credential := NewSFCCredential("validation-test", apiKey)
	client, err := credential.MakeClient(context.Background(), "")
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	err = client.TerminateInstance(context.Background(), "6c7a3ade-1e59-4e04-af6e-365046995a81_test")
	if err != nil {
		t.Fatalf("failed to terminate instance: %v", err)
	}
}
