package validation

import (
	"context"
	"errors"
	"fmt"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/cenkalti/backoff/v4"
)

// CIRunIDLabel re-exports v1.CIRunIDLabel so the sweeper and create paths share one value.
const CIRunIDLabel = v1.CIRunIDLabel

// SweepOpts configures SweepOrphanedInstances.
type SweepOpts struct {
	// RunID is the CIRunIDLabel label value to match. It is REQUIRED and must be unique per CI run.
	RunID string
	// Logf, if set, receives human-readable progress lines. Defaults to fmt.Printf.
	Logf func(format string, args ...any)
}

// SweepResult reports what a sweep scanned and did.
type SweepResult struct {
	Scanned int
	Matched []v1.Instance
	Deleted []v1.CloudProviderInstanceID
	Failed  map[v1.CloudProviderInstanceID]error
}

// SweepOrphanedInstances lists every instance visible to client and terminates exactly the ones
// whose CIRunIDLabel label equals opts.RunID — i.e. the VMs created by this CI run.
func SweepOrphanedInstances(ctx context.Context, client v1.CloudClient, opts SweepOpts) (SweepResult, error) {
	res := SweepResult{Failed: map[v1.CloudProviderInstanceID]error{}}
	logf := opts.Logf
	if logf == nil {
		logf = func(format string, args ...any) { fmt.Printf(format+"\n", args...) }
	}

	// A run ID is mandatory: without it there is nothing safe to match on.
	if opts.RunID == "" {
		return res, errors.New("SweepOrphanedInstances: RunID is required")
	}

	instances, err := client.ListInstances(ctx, v1.ListInstancesArgs{})
	if err != nil {
		return res, fmt.Errorf("failed to list instances: %w", err)
	}
	res.Scanned = len(instances)

	for i := range instances {
		inst := instances[i]
		if inst.Tags[CIRunIDLabel] != opts.RunID {
			continue
		}
		res.Matched = append(res.Matched, inst)
	}

	logf("[sweep] scanned=%d matched=%d runID=%q", res.Scanned, len(res.Matched), opts.RunID)

	for i := range res.Matched {
		inst := res.Matched[i]
		logf("[sweep] terminating cloudID=%s name=%q owner=%s location=%s",
			inst.CloudID, inst.Name, inst.CloudCredRefID, inst.Location)
		if err := terminateWithRetry(ctx, client, inst.CloudID); err != nil {
			res.Failed[inst.CloudID] = err
			logf("[sweep] FAILED to terminate cloudID=%s: %v", inst.CloudID, err)
			continue
		}
		res.Deleted = append(res.Deleted, inst.CloudID)
	}
	return res, nil
}

// terminateWithRetry terminates an instance with a bounded exponential backoff, treating a
// not-found result as success (the instance is already gone).
func terminateWithRetry(ctx context.Context, client v1.CloudCreateTerminateInstance, id v1.CloudProviderInstanceID) error {
	op := func() error {
		termCtx, cancel := context.WithTimeout(ctx, 10*time.Minute)
		defer cancel()
		err := client.TerminateInstance(termCtx, id)
		if err == nil || errors.Is(err, v1.ErrInstanceNotFound) || errors.Is(err, v1.ErrResourceNotFound) {
			return nil
		}
		return err
	}
	return backoff.Retry(op, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 4))
}
