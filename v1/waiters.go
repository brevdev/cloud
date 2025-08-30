package v1

import (
	"context"
	"time"
)

func WaitForInstanceLifecycleStatus(ctx context.Context,
	client CloudInstanceReader,
	instance *Instance,
	status LifecycleStatus,
	timeout time.Duration,
) (*Instance, error) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timeoutCh := time.After(timeout)
	var lastInstance *Instance
	var lastErr error

	for {
		select {
		case <-ctx.Done():
			return instance, ctx.Err()
		case <-timeoutCh:
			if lastInstance != nil {
				return lastInstance, &InstanceWaitTimeoutError{
					Instance: lastInstance,
					Desired:  status,
					Err:      lastErr,
				}
			}
			return instance, &InstanceWaitTimeoutError{
				Instance: instance,
				Desired:  status,
				Err:      lastErr,
			}
		case <-ticker.C:
			inst, err := client.GetInstance(ctx, instance.CloudID)
			if err != nil {
				// If instance is not found, return error immediately
				return inst, &InstanceWaitNotFoundError{
					InstanceID: instance.CloudID,
					Err:        err,
				}
			}
			lastInstance = inst
			if inst.Status.LifecycleStatus == status {
				return inst, nil
			}
		}
	}
}

// InstanceWaitTimeoutError is returned when waiting for an instance times out.
type InstanceWaitTimeoutError struct {
	Instance *Instance
	Desired  LifecycleStatus
	Err      error
}

func (e *InstanceWaitTimeoutError) Error() string {
	return "timeout waiting for instance " + string(e.Instance.CloudID) +
		" to reach status " + string(e.Desired) +
		", last known status: " + string(e.Instance.Status.LifecycleStatus) +
		", last error: " + errString(e.Err)
}

func (e *InstanceWaitTimeoutError) Unwrap() error {
	return e.Err
}

// InstanceWaitNotFoundError is returned when the instance is not found during wait.
type InstanceWaitNotFoundError struct {
	InstanceID CloudProviderInstanceID
	Err        error
}

func (e *InstanceWaitNotFoundError) Error() string {
	return "instance not found: " + string(e.InstanceID) + ", error: " + errString(e.Err)
}

func (e *InstanceWaitNotFoundError) Unwrap() error {
	return e.Err
}

func errString(err error) string {
	if err == nil {
		return "<nil>"
	}
	return err.Error()
}
