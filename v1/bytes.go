package v1

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	zeroBytes = Bytes{value: 0, unit: Byte}

	ErrNegativeValue = errors.New("value must be non-negative")
	ErrEmptyUnit     = errors.New("unit must be set")
	ErrInvalidUnit   = errors.New("invalid unit")
)

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

// String returns the string representation of the Bytes
func (b Bytes) String() string {
	return fmt.Sprintf("%d %s", b.value, b.unit)
}

// MarshalJSON implements the json.Marshaler interface
func (b Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(bytesJSON{
		Value: int64(b.value),
		Unit:  b.unit.value,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (b *Bytes) UnmarshalJSON(data []byte) error {
	var bytesJSON bytesJSON
	if err := json.Unmarshal(data, &bytesJSON); err != nil {
		return err
	}

	if bytesJSON.Value < 0 {
		return ErrNegativeValue
	}

	if bytesJSON.Unit == "" {
		return ErrEmptyUnit
	}

	unit, err := stringToBytesUnit(bytesJSON.Unit)
	if err != nil {
		return err
	}

	b.value = BytesValue(bytesJSON.Value)
	b.unit = unit
	return nil
}

// BytesUnit is a unit of measurement for bytes. Note for implementers: this is defined as a struct rather than a
// type alias to ensure stronger compile-time type checking and to avoid the need for a validation function.
type BytesUnit struct {
	value string
}

// String returns the string representation of the BytesUnit
func (u BytesUnit) String() string {
	return u.value
}

var (
	Byte = BytesUnit{value: "B"}

	// Base 10
	Kilobyte = BytesUnit{value: "KB"}
	Megabyte = BytesUnit{value: "MB"}
	Gigabyte = BytesUnit{value: "GB"}
	Terabyte = BytesUnit{value: "TB"}
	Petabyte = BytesUnit{value: "PB"}
	Exabyte  = BytesUnit{value: "EB"}

	// Base 2
	Kibibyte = BytesUnit{value: "KiB"}
	Mebibyte = BytesUnit{value: "MiB"}
	Gibibyte = BytesUnit{value: "GiB"}
	Tebibyte = BytesUnit{value: "TiB"}
	Pebibyte = BytesUnit{value: "PiB"}
	Exbibyte = BytesUnit{value: "EiB"}
)

func stringToBytesUnit(unit string) (BytesUnit, error) {
	switch unit {
	case "B":
		return Byte, nil
	case "KB":
		return Kilobyte, nil
	case "MB":
		return Megabyte, nil
	case "GB":
		return Gigabyte, nil
	case "TB":
		return Terabyte, nil
	case "PB":
		return Petabyte, nil
	case "EB":
		return Exabyte, nil
	case "KiB":
		return Kibibyte, nil
	case "MiB":
		return Mebibyte, nil
	case "GiB":
		return Gibibyte, nil
	case "TiB":
		return Tebibyte, nil
	case "PiB":
		return Pebibyte, nil
	case "EiB":
		return Exbibyte, nil
	}
	return BytesUnit{}, errors.Join(ErrInvalidUnit, fmt.Errorf("invalid unit: %s", unit))
}
