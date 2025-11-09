package v1

import (
	"encoding/json"
	"fmt"
	"math/bits"

	"github.com/brevdev/cloud/internal/errors"
)

var (
	zeroBytes = Bytes{value: 0, unit: Byte}

	ErrBytesNegativeValue = errors.New("value must be non-negative")
	ErrBytesEmptyUnit     = errors.New("unit must be set")
	ErrBytesInvalidUnit   = errors.New("invalid unit")
)

// NewBytes creates a new Bytes with the given value and unit.
// Returns zeroBytes if value is negative or if the byte count calculation would overflow.
func NewBytes(value BytesValue, unit BytesUnit) Bytes {
	if value < 0 {
		return zeroBytes
	}

	// Check for multiplication overflow using bits.Mul64. If more bits are needed than the lower 64
	// bits, then we know the full result doesn't fit in uint64.
	hi, lo := bits.Mul64(uint64(value), unit.byteCount) //nolint:gosec // 'value' can safely be converted to uint64
	if hi != 0 {
		return zeroBytes
	}
	return Bytes{
		value:     value,
		unit:      unit,
		byteCount: lo,
	}
}

type (
	BytesValue int64
)

// Bytes represents a number of some unit of bytes
type Bytes struct {
	value     BytesValue
	unit      BytesUnit
	byteCount uint64 // TODO: consider using "https://pkg.go.dev/math/big" for this
}

// bytesJSON is the JSON representation of a Bytes. This struct is maintained separately from the core Bytes
// struct to allow for unexported fields to be used in the MarshalJSON and UnmarshalJSON methods.
type bytesJSON struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}

// Value is the whole non-negative number of bytes of the specified unit
func (b Bytes) Value() BytesValue {
	return b.value
}

// Unit is the unit of the byte value
func (b Bytes) Unit() BytesUnit {
	return b.unit
}

// ByteCount is the number of bytes
func (b Bytes) ByteCount() uint64 {
	return b.byteCount
}

// String returns the string representation of the Bytes
func (b Bytes) String() string {
	return fmt.Sprintf("%d %s", b.value, b.unit)
}

// MarshalJSON implements the json.Marshaler interface
func (b Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(bytesJSON{
		Value: int64(b.value),
		Unit:  b.unit.name,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (b *Bytes) UnmarshalJSON(data []byte) error {
	var bytesJSON bytesJSON
	if err := json.Unmarshal(data, &bytesJSON); err != nil {
		return errors.WrapAndTrace(err)
	}

	if bytesJSON.Value < 0 {
		return errors.WrapAndTrace(ErrBytesNegativeValue)
	}

	if bytesJSON.Unit == "" {
		return errors.WrapAndTrace(ErrBytesEmptyUnit)
	}

	unit, err := stringToBytesUnit(bytesJSON.Unit)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	newBytes := NewBytes(BytesValue(bytesJSON.Value), unit)
	*b = newBytes
	return nil
}

// LessThan returns true if the Bytes is less than the other Bytes
func (b Bytes) LessThan(other Bytes) bool {
	return b.byteCount < other.byteCount
}

// GreaterThan returns true if the Bytes is greater than the other Bytes
func (b Bytes) GreaterThan(other Bytes) bool {
	return b.byteCount > other.byteCount
}

// BytesUnit is a unit of measurement for bytes. Note for maintainers: this is defined as a struct rather than a
// type alias to ensure stronger compile-time type checking and to avoid the need for a validation function.
type BytesUnit struct {
	name      string
	byteCount uint64
}

// ByteCountInUnit is the number of bytes in the BytesUnit in the given unit
func (u BytesUnit) ByteCount() uint64 {
	return u.byteCount
}

// String returns the string representation of the BytesUnit
func (u BytesUnit) String() string {
	return u.name
}

var (
	Byte = BytesUnit{name: "B", byteCount: 1}

	// Base 10
	Kilobyte = BytesUnit{name: "KB", byteCount: 1000}
	Megabyte = BytesUnit{name: "MB", byteCount: 1000 * 1000}
	Gigabyte = BytesUnit{name: "GB", byteCount: 1000 * 1000 * 1000}
	Terabyte = BytesUnit{name: "TB", byteCount: 1000 * 1000 * 1000 * 1000}
	Petabyte = BytesUnit{name: "PB", byteCount: 1000 * 1000 * 1000 * 1000 * 1000}

	// Base 2
	Kibibyte = BytesUnit{name: "KiB", byteCount: 1024}
	Mebibyte = BytesUnit{name: "MiB", byteCount: 1024 * 1024}
	Gibibyte = BytesUnit{name: "GiB", byteCount: 1024 * 1024 * 1024}
	Tebibyte = BytesUnit{name: "TiB", byteCount: 1024 * 1024 * 1024 * 1024}
	Pebibyte = BytesUnit{name: "PiB", byteCount: 1024 * 1024 * 1024 * 1024 * 1024}
)

func stringToBytesUnit(unit string) (BytesUnit, error) {
	switch unit {
	case Byte.name:
		return Byte, nil
	case Kilobyte.name:
		return Kilobyte, nil
	case Megabyte.name:
		return Megabyte, nil
	case Gigabyte.name:
		return Gigabyte, nil
	case Terabyte.name:
		return Terabyte, nil
	case Petabyte.name:
		return Petabyte, nil
	case Kibibyte.name:
		return Kibibyte, nil
	case Mebibyte.name:
		return Mebibyte, nil
	case Gibibyte.name:
		return Gibibyte, nil
	case Tebibyte.name:
		return Tebibyte, nil
	case Pebibyte.name:
		return Pebibyte, nil
	}
	return BytesUnit{}, errors.WrapAndTrace(errors.Join(ErrBytesInvalidUnit, fmt.Errorf("invalid unit: %s", unit)))
}
