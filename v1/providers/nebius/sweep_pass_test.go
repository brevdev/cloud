package v1

import (
	"context"
	"errors"
	"reflect"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

const (
	sweepTestRunID = "run-1"
	testDiskID     = "d1"
)

var (
	errListFailed   = errors.New("list failed")
	errDeleteFailed = errors.New("delete failed")
)

// fakeItem is a minimal element for exercising resourceSweepPass without the Nebius SDK.
type fakeItem struct {
	id    string
	match bool
}

// fakePage is one page returned by a fake list callback; err makes that page return an error.
type fakePage struct {
	items []fakeItem
	err   bool
}

// listFromPages returns a list callback that serves pages in order, one per call.
func listFromPages(pages []fakePage) func(context.Context, string, string) ([]fakeItem, string, error) {
	call := 0
	return func(_ context.Context, _, _ string) ([]fakeItem, string, error) {
		if call >= len(pages) {
			return nil, "", nil
		}
		p := pages[call]
		call++
		if p.err {
			return nil, "", errListFailed
		}
		next := ""
		if call < len(pages) {
			next = "more"
		}
		return p.items, next, nil
	}
}

// TestResourceSweepPass covers keep filtering, pagination, and failure counting: a list error or a
// real delete error counts as a failure, while a not-found delete is success.
func TestResourceSweepPass(t *testing.T) {
	match := fakeItem{id: "a", match: true}
	other := fakeItem{id: "c", match: true}
	skip := fakeItem{id: "b", match: false}

	tests := []struct {
		name         string
		pages        []fakePage
		delErr       error
		wantDeleted  int
		wantFailures int
		wantIDs      []string
	}{
		{
			name:        "deletes only matching items across two pages",
			pages:       []fakePage{{items: []fakeItem{match, skip}}, {items: []fakeItem{other}}},
			wantDeleted: 2,
			wantIDs:     []string{"a", "c"},
		},
		{
			name:         "list error counts as a failure and deletes nothing",
			pages:        []fakePage{{err: true}},
			wantFailures: 1,
		},
		{
			name:        "not-found on delete is treated as success",
			pages:       []fakePage{{items: []fakeItem{match}}},
			delErr:      v1.ErrResourceNotFound,
			wantDeleted: 1,
		},
		{
			name:         "real delete error counts as a failure",
			pages:        []fakePage{{items: []fakeItem{match}}},
			delErr:       errDeleteFailed,
			wantFailures: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var deleted []string
			del := func(_ context.Context, id string) error {
				if tt.delErr != nil {
					return tt.delErr
				}
				deleted = append(deleted, id)
				return nil
			}
			keep := func(it fakeItem) (string, string, bool) { return it.id, it.id, it.match }

			gotDeleted, gotFailures := resourceSweepPass(
				context.Background(), "fake", []string{"proj-1"}, t.Logf,
				listFromPages(tt.pages), keep, del)

			if gotDeleted != tt.wantDeleted {
				t.Errorf("deleted = %d, want %d", gotDeleted, tt.wantDeleted)
			}
			if gotFailures != tt.wantFailures {
				t.Errorf("failures = %d, want %d", gotFailures, tt.wantFailures)
			}
			if tt.wantIDs != nil && !reflect.DeepEqual(deleted, tt.wantIDs) {
				t.Errorf("deleted ids = %v, want %v", deleted, tt.wantIDs)
			}
		})
	}
}

// TestResourceSweepPass_PerProjectFailureIsolation verifies a list error in one project counts as a
// failure but still lets another project be swept.
func TestResourceSweepPass_PerProjectFailureIsolation(t *testing.T) {
	list := func(_ context.Context, projectID, _ string) ([]fakeItem, string, error) {
		if projectID == "bad" {
			return nil, "", errListFailed
		}
		return []fakeItem{{id: "g", match: true}}, "", nil
	}
	var deleted []string
	del := func(_ context.Context, id string) error {
		deleted = append(deleted, id)
		return nil
	}
	keep := func(it fakeItem) (string, string, bool) { return it.id, it.id, it.match }

	gotDeleted, gotFailures := resourceSweepPass(
		context.Background(), "fake", []string{"bad", "good"}, t.Logf, list, keep, del)

	if gotFailures != 1 {
		t.Errorf("failures = %d, want 1 (the failing project)", gotFailures)
	}
	if gotDeleted != 1 || !reflect.DeepEqual(deleted, []string{"g"}) {
		t.Errorf("good project not swept: deleted=%d ids=%v", gotDeleted, deleted)
	}
}

// diskMeta builds resource metadata with the given labels for the disk keep-predicate tests.
func diskMeta(labels map[string]string) *common.ResourceMetadata {
	return &common.ResourceMetadata{Id: testDiskID, Name: "disk-1", Labels: labels}
}

// TestDiskKeepPredicates verifies keepByLabel matches only this run's label (nil metadata safe) and
// keepUnattachedDisk also skips an attached disk.
func TestDiskKeepPredicates(t *testing.T) {
	matching := map[string]string{CIRunIDLabel: sweepTestRunID}
	byLabel := keepByLabel[*compute.Disk](sweepTestRunID)
	unattached := keepUnattachedDisk(sweepTestRunID)

	tests := []struct {
		name   string
		keep   func(*compute.Disk) (string, string, bool)
		disk   *compute.Disk
		wantOK bool
	}{
		{"label match", byLabel, &compute.Disk{Metadata: diskMeta(matching)}, true},
		{"label mismatch", byLabel, &compute.Disk{Metadata: diskMeta(map[string]string{CIRunIDLabel: "run-2"})}, false},
		{"nil metadata", byLabel, &compute.Disk{}, false},
		{"unattached kept", unattached, &compute.Disk{Metadata: diskMeta(matching)}, true},
		{"read-write attached skipped", unattached, &compute.Disk{Metadata: diskMeta(matching), Status: &compute.DiskStatus{ReadWriteAttachment: "vm-1"}}, false},
		{"read-only attached skipped", unattached, &compute.Disk{Metadata: diskMeta(matching), Status: &compute.DiskStatus{ReadOnlyAttachments: []string{"vm-1"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, ok := tt.keep(tt.disk)
			if ok != tt.wantOK {
				t.Errorf("ok = %v, want %v", ok, tt.wantOK)
			}
		})
	}
}
