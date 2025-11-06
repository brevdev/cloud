package v1

var zeroBytes = Bytes{value: 0, unit: Byte}

// NewBytes creates a new Bytes with the given value and unit
func NewBytes(value BytesValue, unit BytesUnit) Bytes {
	if value < 0 {
		return zeroBytes
	}
	return Bytes{
		value: value,
		unit:  unit,
	}
}

type (
	BytesValue int64
	BytesUnit  string
)

// Bytes represents a number of some unit of bytes
type Bytes struct {
	value BytesValue
	unit  BytesUnit
}

// Value is the whole non-negative number of bytes of the specified unit
func (b Bytes) Value() BytesValue {
	return b.value
}

// Unit is the unit of the byte value
func (b Bytes) Unit() BytesUnit {
	return b.unit
}

// ByteUnit is a unit of measurement for bytes
const (
	Byte BytesUnit = "B"

	// Base 10
	Kilobyte BytesUnit = "KB"
	Megabyte BytesUnit = "MB"
	Gigabyte BytesUnit = "GB"
	Terabyte BytesUnit = "TB"
	Petabyte BytesUnit = "PB"
	Exabyte  BytesUnit = "EB"

	// Base 2
	Kibibyte BytesUnit = "KiB"
	Mebibyte BytesUnit = "MiB"
	Gibibyte BytesUnit = "GiB"
	Tebibyte BytesUnit = "TiB"
	Pebibyte BytesUnit = "PiB"
	Exbibyte BytesUnit = "EiB"
)
