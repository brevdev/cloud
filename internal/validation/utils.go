package validation

import (
	"context"
	"fmt"
	"time"
)

// WaitForKubernetesClusterPredicate waits for the Kubernetes cluster to satisfy the predicate function. If the predicate returns true, the loop breaks.
type WaitForResourcePredicateOpts[T any] struct {
	GetResource func() (T, error)
	Predicate   func(resource T) bool
	Timeout     time.Duration
	Interval    time.Duration
}

func WaitForResourcePredicate[T any](ctx context.Context, opts WaitForResourcePredicateOpts[T]) error {
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()

	ticker := time.NewTicker(opts.Interval)
	defer ticker.Stop()

	fmt.Printf("Entering WaitForResourcePredicate, timeout: %s, interval: %s\n", opts.Timeout.String(), opts.Interval.String())
	for {
		resource, err := opts.GetResource()
		if err != nil {
			return err
		}

		if opts.Predicate(resource) {
			fmt.Println("Resource satisfies predicate")
			break
		}
		fmt.Printf("Waiting %s for resource to satisfy predicate\n", opts.Interval.String())
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for cluster to satisfy predicate")
		case <-ticker.C:
			continue
		}
	}
	return nil
}
