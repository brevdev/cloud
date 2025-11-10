package v1

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"github.com/brevdev/cloud/internal/errors"
)

var (
	zeroBytes = Bytes{value: 0, unit: Byte}

	ErrBytesInvalidUnit = errors.New("invalid unit")
	ErrBytesNotAnInt64  = errors.New("byte count is not an int64")
	ErrBytesNotAnInt32  = errors.New("byte count is not an int32")
)

// NewBytes creates a new Bytes with the given value and unit.
func NewBytes(value BytesValue, unit BytesUnit) Bytes {
	return Bytes{value: value, unit: unit}
}

type (
	BytesValue int64
)

// Bytes represents a number of some unit of bytes
type Bytes struct {
	value BytesValue
	unit  BytesUnit
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

// ByteCount is the total number of bytes in the Bytes
func (b Bytes) ByteCount() *big.Int {
	bytesByteCount := big.NewInt(0).SetInt64(int64(b.value))
	unitByteCount := big.NewInt(0).SetUint64(b.unit.byteCount)

	return big.NewInt(0).Mul(bytesByteCount, unitByteCount)
}

// ByteCountInUnit is the number of bytes in the Bytes of the given unit. For example, if
// the Bytes is 1000 MB, then:
//
//		1000 MB -> B  = 1000000
//		1000 MB -> KB = 1000
//		1000 MB -> MB = 1
//	    1000 MB -> GB = .001
//
// etc.
func (b Bytes) ByteCountInUnit(unit BytesUnit) *big.Float {
	if b.unit == unit {
		// If the units are the same, return the value as a float
		return big.NewFloat(0).SetInt64(int64(b.value))
	}

	bytesByteCount := big.NewFloat(0).SetInt(b.ByteCount())
	unitByteCount := big.NewFloat(0).SetUint64(unit.byteCount)

	return big.NewFloat(0).Quo(bytesByteCount, unitByteCount)
}

// ByteCountInUnitInt64 attempts to convert the result of ByteCountInUnit to an int64. If this conversion would
// result in an overflow, it returns an ErrBytesNotAnInt64 error. If the byte count is not an integer, the value
// is truncated towards zero.
func (b Bytes) ByteCountInUnitInt64(unit BytesUnit) (int64, error) {
	byteCount := b.ByteCountInUnit(unit)

	byteCountInt64, accuracy := byteCount.Int64()
	if byteCountInt64 == math.MaxInt64 && accuracy == big.Below {
		return 0, errors.WrapAndTrace(errors.Join(ErrBytesNotAnInt64, fmt.Errorf("byte count %v is greater than %d", byteCount, math.MaxInt64)))
	}
	if byteCountInt64 == math.MinInt64 && accuracy == big.Above {
		return 0, errors.WrapAndTrace(errors.Join(ErrBytesNotAnInt64, fmt.Errorf("byte count %v is less than %d", byteCount, math.MinInt64)))
	}
	return byteCountInt64, nil
}

// ByteCountInUnitInt32 attempts to convert the result of ByteCountInUnit to an int32. If this conversion would
// result in an overflow, it returns an ErrBytesNotAnInt32 error.
func (b Bytes) ByteCountInUnitInt32(unit BytesUnit) (int32, error) {
	byteCountInt64, err := b.ByteCountInUnitInt64(unit)
	if err != nil {
		return 0, errors.WrapAndTrace(err)
	}
	if byteCountInt64 > math.MaxInt32 {
		return 0, errors.WrapAndTrace(errors.Join(ErrBytesNotAnInt32, fmt.Errorf("byte count %v is greater than %d", byteCountInt64, math.MaxInt32)))
	}
	return int32(byteCountInt64), nil //nolint:gosec // checked above
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
	return b.ByteCount().Cmp(other.ByteCount()) < 0
}

// GreaterThan returns true if the Bytes is greater than the other Bytes
func (b Bytes) GreaterThan(other Bytes) bool {
	return b.ByteCount().Cmp(other.ByteCount()) > 0
}

// Equal returns true if the Bytes is equal to the other Bytes
func (b Bytes) Equal(other Bytes) bool {
	return b.ByteCount().Cmp(other.ByteCount()) == 0
}

// BytesUnit is a unit of measurement for bytes. Note for maintainers: this is defined as a struct rather than a
// type alias to ensure stronger compile-time type checking and to avoid the need for a validation function.
type BytesUnit struct {
	name      string
	byteCount uint64
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
	Exabyte  = BytesUnit{name: "EB", byteCount: 1000 * 1000 * 1000 * 1000 * 1000 * 1000}

	// Base 2
	Kibibyte = BytesUnit{name: "KiB", byteCount: 1024}
	Mebibyte = BytesUnit{name: "MiB", byteCount: 1024 * 1024}
	Gibibyte = BytesUnit{name: "GiB", byteCount: 1024 * 1024 * 1024}
	Tebibyte = BytesUnit{name: "TiB", byteCount: 1024 * 1024 * 1024 * 1024}
	Pebibyte = BytesUnit{name: "PiB", byteCount: 1024 * 1024 * 1024 * 1024 * 1024}
	Exbibyte = BytesUnit{name: "EiB", byteCount: 1024 * 1024 * 1024 * 1024 * 1024 * 1024}
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
