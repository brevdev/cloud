package v1

import (
	"context"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
)

type Storage struct {
	Count                   int32
	Size                    units.Base2Bytes // TODO: deprecate in favor of SizeByteValue
	SizeByteValue           ByteValue
	Type                    string
	MinSize                 *units.Base2Bytes // TODO: deprecate in favor of MinSizeByteValue
	MinSizeByteValue        *ByteValue
	MaxSize                 *units.Base2Bytes // TODO: deprecate in favor of MaxSizeByteValue
	MaxSizeByteValue        *ByteValue
	PricePerGBHr            *currency.Amount
	IsEphemeral             bool
	IsAdditionalDisk        bool
	RequiresVolumeMountPath bool
	IsElastic               bool
}

type Disk struct {
	Size          units.Base2Bytes // TODO: deprecate in favor of SizeByteValue
	SizeByteValue ByteValue
	Type          string
	MountPath     string
}

type CloudResizeInstanceVolume interface {
	ResizeInstanceVolume(ctx context.Context, args ResizeInstanceVolumeArgs) error
}

type ResizeInstanceVolumeArgs struct {
	InstanceID        CloudProviderInstanceID
	Size              units.Base2Bytes // TODO: deprecate in favor of SizeByteValue
	SizeByteValue     ByteValue
	WaitForOptimizing bool
}
