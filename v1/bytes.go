package v1

var zeroBytes = ByteValue{value: 0, unit: Byte}

type ByteValue struct {
	// Value is the whole non-negative number of bytes of the specified unit
	value uint32

	// Unit is the unit of the byte value
	unit ByteUnit
}

func NewByteValue(value int32, unit ByteUnit) ByteValue {
	if value < 0 {
		return zeroBytes
	}
	return ByteValue{
		value: uint32(value),
		unit:  unit,
	}
}

type ByteUnit string

const (
	Byte ByteUnit = "B"

	// Base 10
	Kilobyte ByteUnit = "KB"
	Megabyte ByteUnit = "MB"
	Gigabyte ByteUnit = "GB"
	Terabyte ByteUnit = "TB"
	Petabyte ByteUnit = "PB"
	Exabyte  ByteUnit = "EB"

	// Base 2
	Kibibyte ByteUnit = "KiB"
	Mebibyte ByteUnit = "MiB"
	Gibibyte ByteUnit = "GiB"
	Tebibyte ByteUnit = "TiB"
	Pebibyte ByteUnit = "PiB"
	Exbibyte ByteUnit = "EiB"
)
