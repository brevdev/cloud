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
	byteCount := big.NewInt(int64(value))
	unitByteCount := big.NewInt(0).Set(unit.byteCount)
	byteCount = byteCount.Mul(byteCount, unitByteCount)

	return Bytes{
		value:     value,
		unit:      unit,
		byteCount: byteCount,
	}
}

type (
	BytesValue int64
)

// Bytes represents a number of some unit of bytes
type Bytes struct {
	value     BytesValue
	unit      BytesUnit
	byteCount *big.Int
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
	// Return a copy of the byte count to avoid mutating the original value
	return big.NewInt(0).Set(b.byteCount)
}

// ByteCountInUnit is the number of bytes in the Bytes of the given unit. For example, if
// the Bytes is 1000 MB, then:
//
//	ByteCountInUnit(Byte)     = 1000000
//	ByteCountInUnit(Kilobyte) = 1000
//	ByteCountInUnit(Megabyte) = 1
func (b Bytes) ByteCountInUnit(unit BytesUnit) *big.Int {
	return big.NewInt(0).Div(b.byteCount, unit.byteCount)
}

// ByteCountInUnitInt64 attempts to convert the result of ByteCountInUnit to an int64. If this conversion would
// result in an overflow, it returns an ErrBytesNotAnInt64 error.
func (b Bytes) ByteCountInUnitInt64(unit BytesUnit) (int64, error) {
	byteCount := b.ByteCountInUnit(unit)
	if !byteCount.IsInt64() {
		return 0, errors.WrapAndTrace(errors.Join(ErrBytesNotAnInt64, fmt.Errorf("byte count %v is not an int64", byteCount)))
	}
	return byteCount.Int64(), nil
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
	return b.byteCount.Cmp(other.byteCount) < 0
}

// GreaterThan returns true if the Bytes is greater than the other Bytes
func (b Bytes) GreaterThan(other Bytes) bool {
	return b.byteCount.Cmp(other.byteCount) > 0
}

// Equal returns true if the Bytes is equal to the other Bytes
func (b Bytes) Equal(other Bytes) bool {
	return b.byteCount.Cmp(other.byteCount) == 0
}

// BytesUnit is a unit of measurement for bytes. Note for maintainers: this is defined as a struct rather than a
// type alias to ensure stronger compile-time type checking and to avoid the need for a validation function.
type BytesUnit struct {
	name      string
	byteCount *big.Int
}

// ByteCountInUnit is the number of bytes in the BytesUnit in the given unit
func (u BytesUnit) ByteCount() *big.Int {
	// Return a copy of the byte count to avoid mutating the original value
	return big.NewInt(0).Set(u.byteCount)
}

// String returns the string representation of the BytesUnit
func (u BytesUnit) String() string {
	return u.name
}

var (
	Byte = BytesUnit{name: "B", byteCount: big.NewInt(1)}

	// Base 10
	Kilobyte = BytesUnit{name: "KB", byteCount: big.NewInt(1000)}
	Megabyte = BytesUnit{name: "MB", byteCount: big.NewInt(1000 * 1000)}
	Gigabyte = BytesUnit{name: "GB", byteCount: big.NewInt(1000 * 1000 * 1000)}
	Terabyte = BytesUnit{name: "TB", byteCount: big.NewInt(1000 * 1000 * 1000 * 1000)}
	Petabyte = BytesUnit{name: "PB", byteCount: big.NewInt(1000 * 1000 * 1000 * 1000 * 1000)}
	Exabyte  = BytesUnit{name: "EB", byteCount: big.NewInt(1000 * 1000 * 1000 * 1000 * 1000 * 1000)}

	// Base 2
	Kibibyte = BytesUnit{name: "KiB", byteCount: big.NewInt(1024)}
	Mebibyte = BytesUnit{name: "MiB", byteCount: big.NewInt(1024 * 1024)}
	Gibibyte = BytesUnit{name: "GiB", byteCount: big.NewInt(1024 * 1024 * 1024)}
	Tebibyte = BytesUnit{name: "TiB", byteCount: big.NewInt(1024 * 1024 * 1024 * 1024)}
	Pebibyte = BytesUnit{name: "PiB", byteCount: big.NewInt(1024 * 1024 * 1024 * 1024 * 1024)}
	Exbibyte = BytesUnit{name: "EiB", byteCount: big.NewInt(1024 * 1024 * 1024 * 1024 * 1024 * 1024)}
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
