package v1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/brevdev/cloud/internal/collections"
	"github.com/brevdev/cloud/internal/ssh"
	"github.com/google/uuid"
)

func ValidateCreateInstance(ctx context.Context, client CloudCreateTerminateInstance, attrs CreateInstanceAttrs, selectedType InstanceType) (*Instance, error) { //nolint:gocyclo // ok
	t0 := time.Now().Add(-time.Minute)
	attrs.RefID = uuid.New().String()
	name, err := makeDebuggableName(attrs.Name)
	if err != nil {
		return nil, err
	}
	attrs.Name = name
	i, err := client.CreateInstance(ctx, attrs)
	if err != nil {
		return nil, err
	}
	var validationErr error
	t1 := time.Now().Add(1 * time.Minute)
	diff := t1.Sub(t0)
	if diff > 3*time.Minute {
		validationErr = errors.Join(validationErr, fmt.Errorf("create instance took too long: %s", diff))
	}
	if i.CreatedAt.Before(t0) {
		validationErr = errors.Join(validationErr, fmt.Errorf("createdAt is before t0: %s", i.CreatedAt))
	}
	if i.CreatedAt.After(t1) {
		validationErr = errors.Join(validationErr, fmt.Errorf("createdAt is after t1: %s", i.CreatedAt))
	}
	if i.Name != name {
		fmt.Printf("name mismatch: %s != %s, input name does not mean return name will be stable\n", i.Name, name)
	}
	if i.RefID != attrs.RefID {
		validationErr = errors.Join(validationErr, fmt.Errorf("refID mismatch: %s != %s", i.RefID, attrs.RefID))
	}
	if attrs.Location != "" && attrs.Location != i.Location {
		validationErr = errors.Join(validationErr, fmt.Errorf("location mismatch: %s != %s", attrs.Location, i.Location))
	}
	if attrs.SubLocation != "" && attrs.SubLocation != i.SubLocation {
		validationErr = errors.Join(validationErr, fmt.Errorf("subLocation mismatch: %s != %s", attrs.SubLocation, i.SubLocation))
	}
	if attrs.InstanceType != "" && attrs.InstanceType != i.InstanceType {
		validationErr = errors.Join(validationErr, fmt.Errorf("instanceType mismatch: %s != %s", attrs.InstanceType, i.InstanceType))
	}
	if selectedType.ID != "" && selectedType.ID != i.InstanceTypeID {
		validationErr = errors.Join(validationErr, fmt.Errorf("instanceTypeID mismatch: %s != %s", selectedType.ID, i.InstanceTypeID))
	}

	return i, validationErr
}

func ValidateListCreatedInstance(ctx context.Context, client CloudCreateTerminateInstance, i *Instance) error {
	// List instances by location and search for the instance by CloudID
	ins, err := client.ListInstances(ctx, ListInstancesArgs{
		Locations: []string{i.Location},
	})
	if err != nil {
		return err
	}
	if len(ins) == 0 {
		return fmt.Errorf("no instances found")
	}
	foundInstance := collections.Find(ins, func(inst Instance) bool {
		return inst.CloudID == i.CloudID
	})
	err = validateInstance(i, foundInstance)
	if err != nil {
		return err
	}

	// List instances by instance ID and search for the instance by CloudID
	ins, err = client.ListInstances(ctx, ListInstancesArgs{
		InstanceIDs: []CloudProviderInstanceID{i.CloudID},
	})
	if err != nil {
		return err
	}
	if len(ins) == 0 {
		return fmt.Errorf("instance not found: %s", i.CloudID)
	}

	foundInstance = collections.Find(ins, func(inst Instance) bool {
		return inst.CloudID == i.CloudID
	})
	err = validateInstance(i, foundInstance)
	if err != nil {
		return err
	}
	return nil
}

func validateInstance(i *Instance, foundInstance *Instance) error {
	var validationErr error
	if foundInstance == nil {
		validationErr = errors.Join(validationErr, fmt.Errorf("instance not found: %s", i.CloudID))
		return validationErr
	}
	if foundInstance.Location != i.Location { //nolint:gocritic // fine
		validationErr = errors.Join(validationErr, fmt.Errorf("location mismatch: %s != %s", foundInstance.Location, i.Location))
	} else if foundInstance.RefID == "" {
		validationErr = errors.Join(validationErr, fmt.Errorf("refID is empty"))
	} else if foundInstance.RefID != i.RefID {
		validationErr = errors.Join(validationErr, fmt.Errorf("refID mismatch: %s != %s", foundInstance.RefID, i.RefID))
	} else if foundInstance.CloudCredRefID == "" {
		validationErr = errors.Join(validationErr, fmt.Errorf("cloudCredRefID is empty"))
	} else if foundInstance.CloudCredRefID != i.CloudCredRefID {
		validationErr = errors.Join(validationErr, fmt.Errorf("cloudCredRefID mismatch: %s != %s", foundInstance.CloudCredRefID, i.CloudCredRefID))
	}
	return validationErr
}

func ValidateTerminateInstance(ctx context.Context, client CloudCreateTerminateInstance, instance *Instance) error {
	err := client.TerminateInstance(ctx, instance.CloudID)
	if err != nil {
		return err
	}
	// TODO wait for instance to go into terminating state
	return nil
}

func ValidateStopStartInstance(ctx context.Context, client CloudStopStartInstance, instance *Instance) error {
	err := client.StopInstance(ctx, instance.CloudID)
	if err != nil {
		return err
	}
	// TODO wait for stopped
	err = client.StartInstance(ctx, instance.CloudID)
	if err != nil {
		return err
	}
	// TODO wait for running
	return nil
}

func ValidateMergeInstanceForUpdate(client UpdateHandler, currInst Instance, newInst Instance) error {
	mergedInst := client.MergeInstanceForUpdate(currInst, newInst)

	var validationErr error
	if currInst.Name != mergedInst.Name {
		validationErr = errors.Join(validationErr, fmt.Errorf("name mismatch: %s != %s", currInst.Name, mergedInst.Name))
	}
	if currInst.RefID != mergedInst.RefID {
		validationErr = errors.Join(validationErr, fmt.Errorf("refID mismatch: %s != %s", currInst.RefID, mergedInst.RefID))
	}
	if currInst.Location != mergedInst.Location {
		validationErr = errors.Join(validationErr, fmt.Errorf("location mismatch: %s != %s", currInst.Location, newInst.Location))
	}
	if currInst.SubLocation != mergedInst.SubLocation {
		validationErr = errors.Join(validationErr, fmt.Errorf("subLocation mismatch: %s != %s", currInst.SubLocation, mergedInst.SubLocation))
	}
	if currInst.InstanceType != "" && currInst.InstanceType != mergedInst.InstanceType {
		validationErr = errors.Join(validationErr, fmt.Errorf("instanceType mismatch: %s != %s", currInst.InstanceType, mergedInst.InstanceType))
	}
	if currInst.InstanceTypeID != "" && currInst.InstanceTypeID != mergedInst.InstanceTypeID {
		validationErr = errors.Join(validationErr, fmt.Errorf("instanceTypeID mismatch: %s != %s", currInst.InstanceTypeID, mergedInst.InstanceTypeID))
	}
	if currInst.CloudCredRefID != mergedInst.CloudCredRefID {
		validationErr = errors.Join(validationErr, fmt.Errorf("cloudCredRefID mismatch: %s != %s", currInst.CloudCredRefID, mergedInst.CloudCredRefID))
	}
	if currInst.VolumeType != "" && currInst.VolumeType != mergedInst.VolumeType {
		validationErr = errors.Join(validationErr, fmt.Errorf("volumeType mismatch: %s != %s", currInst.VolumeType, mergedInst.VolumeType))
	}
	if currInst.Spot != mergedInst.Spot {
		validationErr = errors.Join(validationErr, fmt.Errorf("spot mismatch: %v != %v", currInst.Spot, mergedInst.Spot))
	}
	return validationErr
}

func ValidateInstanceSSHAccessible(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string) error {
	var err error
	instance, err = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, PendingToRunningTimeout)
	if err != nil {
		return err
	}
	sshUser := instance.SSHUser
	sshPort := instance.SSHPort
	publicIP := instance.PublicIP
	// Validate that we have the required SSH connection details
	if sshUser == "" {
		return fmt.Errorf("SSH user is not set for instance %s", instance.CloudID)
	}
	if sshPort == 0 {
		return fmt.Errorf("SSH port is not set for instance %s", instance.CloudID)
	}
	if publicIP == "" {
		return fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	err = ssh.WaitForSSH(ctx, ssh.ConnectionConfig{
		User:     sshUser,
		HostPort: fmt.Sprintf("%s:%d", publicIP, sshPort),
		PrivKey:  privateKey,
	}, ssh.WaitForSSHOptions{
		Timeout: RunningSSHTimeout,
	})
	if err != nil {
		return err
	}

	fmt.Printf("SSH connection validated successfully for %s@%s:%d\n", sshUser, publicIP, sshPort)

	return nil
}
