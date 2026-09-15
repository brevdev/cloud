package validation

import (
	"context"
	"testing"
	"time"

	v1 "github.com/brevdev/cloud/v1"
)

// fakeSweepClient is a CloudClient whose ListInstances returns a canned set and whose
// TerminateInstance records the IDs it was asked to delete. All other methods come from
// NotImplCloudClient and return ErrNotImplemented (the sweeper never calls them).
type fakeSweepClient struct {
	v1.NotImplCloudClient
	instances  []v1.Instance
	terminated []v1.CloudProviderInstanceID
}

var _ v1.CloudClient = (*fakeSweepClient)(nil)

func (c *fakeSweepClient) ListInstances(context.Context, v1.ListInstancesArgs) ([]v1.Instance, error) {
	return c.instances, nil
}

func (c *fakeSweepClient) TerminateInstance(_ context.Context, id v1.CloudProviderInstanceID) error {
	c.terminated = append(c.terminated, id)
	return nil
}

func matchedIDs(res SweepResult) []v1.CloudProviderInstanceID {
	ids := make([]v1.CloudProviderInstanceID, 0, len(res.Matched))
	for _, m := range res.Matched {
		ids = append(ids, m.CloudID)
	}
	return ids
}

func assertSameSet(t *testing.T, got, want []v1.CloudProviderInstanceID) {
	t.Helper()
	gm := map[v1.CloudProviderInstanceID]bool{}
	for _, g := range got {
		gm[g] = true
	}
	wm := map[v1.CloudProviderInstanceID]bool{}
	for _, w := range want {
		wm[w] = true
	}
	for w := range wm {
		if !gm[w] {
			t.Fatalf("missing %q: got %v want %v", w, got, want)
		}
	}
	for g := range gm {
		if !wm[g] {
			t.Fatalf("unexpected %q: got %v want %v", g, got, want)
		}
	}
}

// TestSweepOrphanedInstances_OnlyMatchesRunID is the safety-critical test: the sweep terminates
// exactly the VMs carrying this run's ci-run-id label and nothing else.
func TestSweepOrphanedInstances_OnlyMatchesRunID(t *testing.T) {
	old := time.Now().Add(-10 * 24 * time.Hour)
	fresh := time.Now().Add(-1 * time.Minute)

	instances := []v1.Instance{
		// This run's VMs — must be deleted, regardless of age.
		{CloudID: "A-thisrun", CloudCredRefID: "validation-test", CreatedAt: fresh, Tags: v1.Tags{CIRunIDLabel: "run-1"}},
		{CloudID: "B-thisrun-int", CloudCredRefID: "integration-test-ref", CreatedAt: fresh, Tags: v1.Tags{CIRunIDLabel: "run-1"}},
		// Another run's VM — must NOT be deleted (different run id).
		{CloudID: "C-otherrun", CloudCredRefID: "validation-test", CreatedAt: old, Tags: v1.Tags{CIRunIDLabel: "run-2"}},
		// Legacy CI VM with no run label (predates this feature) — must NOT be deleted (manual cleanup).
		{CloudID: "D-legacy-nolabel", CloudCredRefID: "validation-test", CreatedAt: old},
		// Production VM that reused a CI credential ref but has no run label — must NOT be deleted.
		{CloudID: "E-prod-reused-cred", CloudCredRefID: "validation-test", CreatedAt: old},
		// Ordinary production VM — must NOT be deleted.
		{CloudID: "F-prod", CloudCredRefID: "cloud-cred-abc123", CreatedAt: old},
	}

	fc := &fakeSweepClient{instances: instances}
	res, err := SweepOrphanedInstances(context.Background(), fc, SweepOpts{RunID: "run-1", Logf: t.Logf})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []v1.CloudProviderInstanceID{"A-thisrun", "B-thisrun-int"}
	assertSameSet(t, matchedIDs(res), want)
	assertSameSet(t, res.Deleted, want)
	assertSameSet(t, fc.terminated, want)
}

// TestSweepOrphanedInstances_RequiresRunID pins the safety invariant: without a run ID the sweep
// refuses to run and terminates nothing.
func TestSweepOrphanedInstances_RequiresRunID(t *testing.T) {
	fc := &fakeSweepClient{instances: []v1.Instance{
		{CloudID: "x", CloudCredRefID: "validation-test", Tags: v1.Tags{CIRunIDLabel: "run-1"}},
	}}
	if _, err := SweepOrphanedInstances(context.Background(), fc, SweepOpts{RunID: ""}); err == nil {
		t.Error("expected an error when RunID is empty")
	}
	if len(fc.terminated) != 0 {
		t.Errorf("empty RunID must terminate nothing, terminated: %v", fc.terminated)
	}
}
