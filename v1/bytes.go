package v1

var zeroBytes = byteValue{value: 0, unit: Byte}

// NewByteValue creates a new ByteValue with the given value and unit
func NewByteValue(value int32, unit ByteUnit) ByteValue {
	if value < 0 {
		return zeroBytes
	}
	return byteValue{
		value: uint32(value),
		unit:  unit,
	}
}

// ByteValue is a value that represents a number of bytes
type ByteValue interface {
	// Value is the whole non-negative number of bytes of the specified unit
	Value() uint32

	// Unit is the unit of the byte value
	Unit() ByteUnit
}

type byteValue struct {
	value uint32
	unit  ByteUnit
}

func (b byteValue) Value() uint32 {
	return b.value
}

func (b byteValue) Unit() ByteUnit {
	return b.unit
}

// ByteUnit is a unit of measurement for bytes
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
