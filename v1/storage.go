package v1

import (
	"context"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
)

type Storage struct {
	Count                   int32
	Size                    units.Base2Bytes
	Type                    string
	MinSize                 *units.Base2Bytes
	MaxSize                 *units.Base2Bytes
	PricePerGBHr            *currency.Amount
	IsEphemeral             bool
	IsAdditionalDisk        bool
	RequiresVolumeMountPath bool
	IsElastic               bool
}

type Disk struct {
	Size      units.Base2Bytes
	Type      string
	MountPath string
}

type CloudResizeInstanceVolume interface {
	ResizeInstanceVolume(ctx context.Context, args ResizeInstanceVolumeArgs) error
}

type ResizeInstanceVolumeArgs struct {
	InstanceID        CloudProviderInstanceID
	Size              units.Base2Bytes
	WaitForOptimizing bool
}
