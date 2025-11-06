package v1

import (
	"context"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
)

type Storage struct {
	Count                   int32
	Size                    units.Base2Bytes // TODO: deprecate in favor of SizeByteValue
	SizeBytes               Bytes
	Type                    string
	MinSize                 *units.Base2Bytes // TODO: deprecate in favor of MinSizeByteValue
	MinSizeBytes            *Bytes
	MaxSize                 *units.Base2Bytes // TODO: deprecate in favor of MaxSizeByteValue
	MaxSizeBytes            *Bytes
	PricePerGBHr            *currency.Amount
	IsEphemeral             bool
	IsAdditionalDisk        bool
	RequiresVolumeMountPath bool
	IsElastic               bool
}

type Disk struct {
	Size      units.Base2Bytes // TODO: deprecate in favor of SizeByteValue
	SizeBytes Bytes
	Type      string
	MountPath string
}

type CloudResizeInstanceVolume interface {
	ResizeInstanceVolume(ctx context.Context, args ResizeInstanceVolumeArgs) error
}

type ResizeInstanceVolumeArgs struct {
	InstanceID        CloudProviderInstanceID
	Size              units.Base2Bytes // TODO: deprecate in favor of SizeByteValue
	SizeBytes         Bytes
	WaitForOptimizing bool
}
