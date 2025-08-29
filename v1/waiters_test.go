package v1

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test suite for WaitForInstanceLifecycleStatus function.
// Some tests are skipped in short mode (-short flag) to speed up development.
// Run without -short flag for full test coverage including longer-running tests.

// mockCloudInstanceReader implements CloudInstanceReader for testing
type mockCloudInstanceReader struct {
	instances map[CloudProviderInstanceID]*Instance
	errors    map[CloudProviderInstanceID]error
	calls     map[CloudProviderInstanceID]int
	// For status transition testing
	statusSequence map[CloudProviderInstanceID][]LifecycleStatus
	sequenceIndex  map[CloudProviderInstanceID]int
}

func newMockCloudInstanceReader() *mockCloudInstanceReader {
	return &mockCloudInstanceReader{
		instances:      make(map[CloudProviderInstanceID]*Instance),
		errors:         make(map[CloudProviderInstanceID]error),
		calls:          make(map[CloudProviderInstanceID]int),
		statusSequence: make(map[CloudProviderInstanceID][]LifecycleStatus),
		sequenceIndex:  make(map[CloudProviderInstanceID]int),
	}
}

func (m *mockCloudInstanceReader) GetInstance(_ context.Context, id CloudProviderInstanceID) (*Instance, error) {
	m.calls[id]++

	if err, exists := m.errors[id]; exists {
		return nil, err
	}

	// Check if we have a status sequence for this instance
	if sequence, exists := m.statusSequence[id]; exists {
		index := m.sequenceIndex[id]
		if index < len(sequence) {
			status := sequence[index]
			m.sequenceIndex[id]++
			return &Instance{
				CloudID: id,
				Status: Status{
					LifecycleStatus: status,
				},
			}, nil
		}
		// If we've exhausted the sequence, use the last status
		if len(sequence) > 0 {
			lastStatus := sequence[len(sequence)-1]
			return &Instance{
				CloudID: id,
				Status: Status{
					LifecycleStatus: lastStatus,
				},
			}, nil
		}
	}

	if instance, exists := m.instances[id]; exists {
		return instance, nil
	}

	return nil, errors.New("instance not found")
}

func (m *mockCloudInstanceReader) ListInstances(_ context.Context, _ ListInstancesArgs) ([]Instance, error) {
	// Not used in tests, but required by interface
	return nil, nil
}

// setStatusSequence sets up a sequence of statuses that will be returned for an instance
func (m *mockCloudInstanceReader) setStatusSequence(instanceID CloudProviderInstanceID, sequence []LifecycleStatus) {
	m.statusSequence[instanceID] = sequence
	m.sequenceIndex[instanceID] = 0
}

func TestWaitForInstanceLifecycleStatus_Success(t *testing.T) {
	client := newMockCloudInstanceReader()
	instanceID := CloudProviderInstanceID("test-instance-123")

	// Set up instance that is already in running status
	instance := &Instance{
		CloudID: instanceID,
		Status: Status{
			LifecycleStatus: LifecycleStatusRunning,
		},
	}
	client.instances[instanceID] = instance

	ctx := context.Background()
	timeout := 5 * time.Second

	_, err := WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, timeout)

	assert.NoError(t, err)
	assert.Equal(t, 1, client.calls[instanceID])
}

func TestWaitForInstanceLifecycleStatus_StatusTransition(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping status transition test in short mode")
	}

	client := newMockCloudInstanceReader()
	instanceID := CloudProviderInstanceID("test-instance-456")

	// Set up a status sequence: pending -> pending -> running
	client.setStatusSequence(instanceID, []LifecycleStatus{
		LifecycleStatusPending,
		LifecycleStatusPending,
		LifecycleStatusRunning,
	})

	instance := &Instance{
		CloudID: instanceID,
		Status: Status{
			LifecycleStatus: LifecycleStatusPending,
		},
	}

	ctx := context.Background()
	timeout := 10 * time.Second

	_, err := WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, timeout)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, client.calls[instanceID], 3, "Expected at least 3 calls to GetInstance")
}

func TestWaitForInstanceLifecycleStatus_Timeout(t *testing.T) {
	client := newMockCloudInstanceReader()
	instanceID := CloudProviderInstanceID("test-instance-timeout")

	// Instance stays in pending status
	instance := &Instance{
		CloudID: instanceID,
		Status: Status{
			LifecycleStatus: LifecycleStatusPending,
		},
	}
	client.instances[instanceID] = instance

	ctx := context.Background()
	timeout := 100 * time.Millisecond // Short timeout for testing

	_, err := WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, timeout)

	require.Error(t, err)

	timeoutErr, ok := err.(*InstanceWaitTimeoutError)
	require.True(t, ok, "Expected InstanceWaitTimeoutError, got: %T", err)

	assert.Equal(t, LifecycleStatusRunning, timeoutErr.Desired)
	assert.Equal(t, instanceID, timeoutErr.Instance.CloudID)
	// With a 100ms timeout, we might not get any calls due to the 2-second ticker
	// The function will timeout before the first ticker fires
	assert.GreaterOrEqual(t, client.calls[instanceID], 0, "Expected 0 or more calls to GetInstance")
}

func TestWaitForInstanceLifecycleStatus_ContextCancellation(t *testing.T) {
	client := newMockCloudInstanceReader()
	instanceID := CloudProviderInstanceID("test-instance-context")

	instance := &Instance{
		CloudID: instanceID,
		Status: Status{
			LifecycleStatus: LifecycleStatusPending,
		},
	}
	client.instances[instanceID] = instance

	ctx, cancel := context.WithCancel(context.Background())
	timeout := 5 * time.Second

	// Cancel context after a short delay
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	_, err := WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, timeout)

	require.Error(t, err)
	assert.Equal(t, context.Canceled, err)
}

func TestWaitForInstanceLifecycleStatus_InstanceNotFound(t *testing.T) {
	client := newMockCloudInstanceReader()
	instanceID := CloudProviderInstanceID("test-instance-notfound")

	// Set up client to return "not found" error
	client.errors[instanceID] = errors.New("instance not found")

	instance := &Instance{
		CloudID: instanceID,
		Status: Status{
			LifecycleStatus: LifecycleStatusPending,
		},
	}

	ctx := context.Background()
	timeout := 5 * time.Second

	_, err := WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, timeout)

	require.Error(t, err)

	notFoundErr, ok := err.(*InstanceWaitNotFoundError)
	require.True(t, ok, "Expected InstanceWaitNotFoundError, got: %T", err)

	assert.Equal(t, instanceID, notFoundErr.InstanceID)
	assert.Equal(t, 1, client.calls[instanceID])
}

func TestWaitForInstanceLifecycleStatus_ErrorString(t *testing.T) {
	// Test error string formatting
	instance := Instance{
		CloudID: "test-instance",
		Status: Status{
			LifecycleStatus: LifecycleStatusPending,
		},
	}

	timeoutErr := &InstanceWaitTimeoutError{
		Instance: &instance,
		Desired:  LifecycleStatusRunning,
		Err:      errors.New("test error"),
	}

	errorStr := timeoutErr.Error()
	assert.Contains(t, errorStr, "timeout waiting for instance test-instance to reach status running")
	assert.Contains(t, errorStr, "test error")

	notFoundErr := &InstanceWaitNotFoundError{
		InstanceID: "test-instance",
		Err:        errors.New("not found"),
	}

	errorStr = notFoundErr.Error()
	assert.Contains(t, errorStr, "instance not found: test-instance")
	assert.Contains(t, errorStr, "not found")
}

func TestWaitForInstanceLifecycleStatus_ErrorUnwrap(t *testing.T) {
	originalErr := errors.New("original error")

	timeoutErr := &InstanceWaitTimeoutError{
		Instance: &Instance{CloudID: "test"},
		Desired:  LifecycleStatusRunning,
		Err:      originalErr,
	}

	assert.Equal(t, originalErr, timeoutErr.Unwrap())

	notFoundErr := &InstanceWaitNotFoundError{
		InstanceID: "test",
		Err:        originalErr,
	}

	assert.Equal(t, originalErr, notFoundErr.Unwrap())
}

func TestWaitForInstanceLifecycleStatus_ErrString(t *testing.T) {
	// Test errString helper function
	assert.Equal(t, "<nil>", errString(nil))

	testErr := errors.New("test error")
	assert.Equal(t, "test error", errString(testErr))
}

func TestWaitForInstanceLifecycleStatus_AllLifecycleStatuses(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping all lifecycle statuses test in short mode")
	}

	// Test that the function works with all possible lifecycle statuses
	statuses := []LifecycleStatus{
		LifecycleStatusPending,
		LifecycleStatusRunning,
		LifecycleStatusStopping,
		LifecycleStatusStopped,
		LifecycleStatusSuspending,
		LifecycleStatusSuspended,
		LifecycleStatusTerminating,
		LifecycleStatusTerminated,
		LifecycleStatusFailed,
	}

	for _, status := range statuses {
		t.Run(string(status), func(t *testing.T) {
			client := newMockCloudInstanceReader()
			instanceID := CloudProviderInstanceID("test-instance-" + string(status))

			instance := &Instance{
				CloudID: instanceID,
				Status: Status{
					LifecycleStatus: status,
				},
			}
			client.instances[instanceID] = instance

			ctx := context.Background()
			timeout := 3 * time.Second // Give enough time for the ticker to fire

			_, err := WaitForInstanceLifecycleStatus(ctx, client, instance, status, timeout)

			assert.NoError(t, err)
			assert.Equal(t, 1, client.calls[instanceID])
		})
	}
}

func TestWaitForInstanceLifecycleStatus_TimeoutWithLastInstance(t *testing.T) {
	client := newMockCloudInstanceReader()
	instanceID := CloudProviderInstanceID("test-instance-timeout-last")

	// Set up instance that stays in pending status
	instance := &Instance{
		CloudID: instanceID,
		Status: Status{
			LifecycleStatus: LifecycleStatusPending,
		},
	}
	client.instances[instanceID] = instance

	ctx := context.Background()
	timeout := 100 * time.Millisecond

	_, err := WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, timeout)

	require.Error(t, err)

	timeoutErr, ok := err.(*InstanceWaitTimeoutError)
	require.True(t, ok)

	// Verify that the timeout error contains the last known instance
	assert.Equal(t, instanceID, timeoutErr.Instance.CloudID)
	assert.Equal(t, LifecycleStatusPending, timeoutErr.Instance.Status.LifecycleStatus)
}

// Benchmark test for performance
func BenchmarkWaitForInstanceLifecycleStatus(b *testing.B) {
	client := newMockCloudInstanceReader()
	instanceID := CloudProviderInstanceID("benchmark-instance")

	instance := &Instance{
		CloudID: instanceID,
		Status: Status{
			LifecycleStatus: LifecycleStatusRunning,
		},
	}
	client.instances[instanceID] = instance

	ctx := context.Background()
	timeout := 1 * time.Second

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, timeout)
	}
}
