package v1

import "context"

type CloudQuota interface {
	GetInstanceTypeQuotas(ctx context.Context, args GetInstanceTypeQuotasArgs) (Quota, error)
}

type GetInstanceTypeQuotasArgs struct {
	InstanceType string
}

type Quota struct {
	ID      string
	Name    string
	Maximum int    // maximum number of units
	Current int    // current number of units being used
	Unit    string // if gpu, will trigger https://github.com/brevdev/brev-deploy/blob/e8ad711e641794e0b2f7b582e301535826aca8b3/internal/routes/instance_types.go#L408
}
