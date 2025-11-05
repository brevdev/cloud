package v1

var ZeroBytes = NewByteValue(0, Byte)

type ByteValue struct {
	// Value is the whole non-negative number of bytes of the specified unit
	value uint64

	// Unit is the unit of the byte value
	unit ByteUnit
}

func NewByteValue(value uint64, unit ByteUnit) ByteValue {
	return ByteValue{
		value: value,
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
