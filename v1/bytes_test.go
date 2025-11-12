package v1

import (
	"encoding/json"
	"errors"
	"math/big"
	"testing"
)

func TestNewBytes(t *testing.T) {
	tests := []struct {
		name    string
		value   BytesValue
		unit    BytesUnit
		want    Bytes
		wantErr error
	}{
		{name: "1000 B", value: 1000, unit: Byte, want: NewBytes(1000, Byte), wantErr: nil},
		{name: "1000 KB", value: 1000, unit: Kilobyte, want: NewBytes(1000, Kilobyte), wantErr: nil},
		{name: "1000 MB", value: 1000, unit: Megabyte, want: NewBytes(1000, Megabyte), wantErr: nil},
		{name: "1000 GB", value: 1000, unit: Gigabyte, want: NewBytes(1000, Gigabyte), wantErr: nil},
		{name: "1000 TB", value: 1000, unit: Terabyte, want: NewBytes(1000, Terabyte), wantErr: nil},
		{name: "1000 PB", value: 1000, unit: Petabyte, want: NewBytes(1000, Petabyte), wantErr: nil},
		{name: "1000 KiB", value: 1000, unit: Kibibyte, want: NewBytes(1000, Kibibyte), wantErr: nil},
		{name: "1000 MiB", value: 1000, unit: Mebibyte, want: NewBytes(1000, Mebibyte), wantErr: nil},
		{name: "1000 GiB", value: 1000, unit: Gibibyte, want: NewBytes(1000, Gibibyte), wantErr: nil},
		{name: "1000 TiB", value: 1000, unit: Tebibyte, want: NewBytes(1000, Tebibyte), wantErr: nil},
		{name: "1000 PiB", value: 1000, unit: Pebibyte, want: NewBytes(1000, Pebibyte), wantErr: nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bytes := NewBytes(test.value, test.unit)
			if !bytes.Equal(test.want) {
				t.Errorf("NewBytes() = %v, want %v", bytes, test.want)
			}
		})
	}
}

func TestBytesMarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		bytes   Bytes
		want    string
		wantErr error
	}{
		{name: "1000 B", bytes: NewBytes(1000, Byte), want: `{"value":1000,"unit":"B"}`, wantErr: nil},
		{name: "1000 KB", bytes: NewBytes(1000, Kilobyte), want: `{"value":1000,"unit":"KB"}`, wantErr: nil},
		{name: "1000 MB", bytes: NewBytes(1000, Megabyte), want: `{"value":1000,"unit":"MB"}`, wantErr: nil},
		{name: "1000 GB", bytes: NewBytes(1000, Gigabyte), want: `{"value":1000,"unit":"GB"}`, wantErr: nil},
		{name: "1000 TB", bytes: NewBytes(1000, Terabyte), want: `{"value":1000,"unit":"TB"}`, wantErr: nil},
		{name: "1000 PB", bytes: NewBytes(1000, Petabyte), want: `{"value":1000,"unit":"PB"}`, wantErr: nil},
		{name: "1000 KiB", bytes: NewBytes(1000, Kibibyte), want: `{"value":1000,"unit":"KiB"}`, wantErr: nil},
		{name: "1000 MiB", bytes: NewBytes(1000, Mebibyte), want: `{"value":1000,"unit":"MiB"}`, wantErr: nil},
		{name: "1000 GiB", bytes: NewBytes(1000, Gibibyte), want: `{"value":1000,"unit":"GiB"}`, wantErr: nil},
		{name: "1000 TiB", bytes: NewBytes(1000, Tebibyte), want: `{"value":1000,"unit":"TiB"}`, wantErr: nil},
		{name: "1000 PiB", bytes: NewBytes(1000, Pebibyte), want: `{"value":1000,"unit":"PiB"}`, wantErr: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			json, err := json.Marshal(test.bytes)
			if err != nil {
				if test.wantErr == nil {
					t.Fatalf("json.Marshal() error = %v, want nil", err)
				}
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("json.Marshal() error = %v, want %v", err, test.wantErr)
				}
			} else if string(json) != test.want {
				t.Errorf("json.Marshal() = %v, want %v", string(json), test.want)
			}
		})
	}
}

func TestBytesUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		want    Bytes
		wantErr error
	}{
		{name: "Empty bytes", json: `{"value":0,"unit":""}`, want: zeroBytes, wantErr: nil},
		{name: "1000 B", json: `{"value":1000,"unit":"B"}`, want: NewBytes(1000, Byte), wantErr: nil},
		{name: "1000 KB", json: `{"value":1000,"unit":"KB"}`, want: NewBytes(1000, Kilobyte), wantErr: nil},
		{name: "1000 MB", json: `{"value":1000,"unit":"MB"}`, want: NewBytes(1000, Megabyte), wantErr: nil},
		{name: "1000 GB", json: `{"value":1000,"unit":"GB"}`, want: NewBytes(1000, Gigabyte), wantErr: nil},
		{name: "1000 TB", json: `{"value":1000,"unit":"TB"}`, want: NewBytes(1000, Terabyte), wantErr: nil},
		{name: "1000 PB", json: `{"value":1000,"unit":"PB"}`, want: NewBytes(1000, Petabyte), wantErr: nil},
		{name: "1000 KiB", json: `{"value":1000,"unit":"KiB"}`, want: NewBytes(1000, Kibibyte), wantErr: nil},
		{name: "1000 MiB", json: `{"value":1000,"unit":"MiB"}`, want: NewBytes(1000, Mebibyte), wantErr: nil},
		{name: "1000 GiB", json: `{"value":1000,"unit":"GiB"}`, want: NewBytes(1000, Gibibyte), wantErr: nil},
		{name: "1000 TiB", json: `{"value":1000,"unit":"TiB"}`, want: NewBytes(1000, Tebibyte), wantErr: nil},
		{name: "1000 PiB", json: `{"value":1000,"unit":"PiB"}`, want: NewBytes(1000, Pebibyte), wantErr: nil},

		{name: "Empty unit", json: `{"value":1000,"unit":""}`, want: zeroBytes, wantErr: ErrBytesInvalidUnit},
		{name: "Invalid unit", json: `{"value":1000,"unit":"invalid"}`, want: zeroBytes, wantErr: ErrBytesInvalidUnit},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var bytes Bytes
			err := json.Unmarshal([]byte(test.json), &bytes)
			if err != nil {
				if test.wantErr == nil {
					t.Fatalf("json.Unmarshal() error = %v, want nil", err)
				}
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("json.Unmarshal() error = %v, want %v", err, test.wantErr)
				}
			} else if !bytes.Equal(test.want) {
				t.Errorf("json.Unmarshal() = %v, want %v", bytes, test.want)
			}
		})
	}
}

func TestBytesString(t *testing.T) {
	tests := []struct {
		name  string
		bytes Bytes
		want  string
	}{
		{name: "Byte", bytes: NewBytes(1000, Byte), want: "1000 B"},
		{name: "Kilobyte", bytes: NewBytes(1000, Kilobyte), want: "1000 KB"},
		{name: "Megabyte", bytes: NewBytes(1000, Megabyte), want: "1000 MB"},
		{name: "Gigabyte", bytes: NewBytes(1000, Gigabyte), want: "1000 GB"},
		{name: "Terabyte", bytes: NewBytes(1000, Terabyte), want: "1000 TB"},
		{name: "Petabyte", bytes: NewBytes(1000, Petabyte), want: "1000 PB"},
		{name: "Kibibyte", bytes: NewBytes(1000, Kibibyte), want: "1000 KiB"},
		{name: "Mebibyte", bytes: NewBytes(1000, Mebibyte), want: "1000 MiB"},
		{name: "Gibibyte", bytes: NewBytes(1000, Gibibyte), want: "1000 GiB"},
		{name: "Tebibyte", bytes: NewBytes(1000, Tebibyte), want: "1000 TiB"},
		{name: "Pebibyte", bytes: NewBytes(1000, Pebibyte), want: "1000 PiB"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.bytes.String()
			if got != test.want {
				t.Errorf("Bytes.String() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestBytesGetters(t *testing.T) {
	// Test that Value() and Unit() work correctly and that Unit() can be used
	b := NewBytes(1024, Gigabyte)

	if b.Value() != 1024 {
		t.Errorf("Value() = %v, want 1024", b.Value())
	}

	if b.Unit() != Gigabyte {
		t.Errorf("Unit() = %v, want Gigabyte", b.Unit())
	}

	// Test that we can get the string value from Unit()
	if b.Unit().String() != "GB" {
		t.Errorf("Unit().String() = %v, want GB", b.Unit().String())
	}
}

func TestBytesEqual(t *testing.T) {
	tests := []struct {
		name  string
		bytes Bytes
		other Bytes
		want  bool
	}{
		{name: "1000 B == 1000 B", bytes: NewBytes(1000, Byte), other: NewBytes(1000, Byte), want: true},
		{name: "1000 B != 1000 KB", bytes: NewBytes(1000, Byte), other: NewBytes(1000, Kilobyte), want: false},
		{name: "1000 KB != 1000 B", bytes: NewBytes(1000, Kilobyte), other: NewBytes(1000, Byte), want: false},
		{name: "1000 KB == 1000 KB", bytes: NewBytes(1000, Kilobyte), other: NewBytes(1000, Kilobyte), want: true},
		{name: "1000 KB != 1000 MB", bytes: NewBytes(1000, Kilobyte), other: NewBytes(1000, Megabyte), want: false},
		{name: "1000 MB != 1000 KB", bytes: NewBytes(1000, Megabyte), other: NewBytes(1000, Kilobyte), want: false},
		{name: "1000 MB == 1000 MB", bytes: NewBytes(1000, Megabyte), other: NewBytes(1000, Megabyte), want: true},
		{name: "1000 MB != 1000 GB", bytes: NewBytes(1000, Megabyte), other: NewBytes(1000, Gigabyte), want: false},
		{name: "1000 GB != 1000 MB", bytes: NewBytes(1000, Gigabyte), other: NewBytes(1000, Megabyte), want: false},
		{name: "1000 GB == 1000 GB", bytes: NewBytes(1000, Gigabyte), other: NewBytes(1000, Gigabyte), want: true},
		{name: "1000 GB != 1000 TB", bytes: NewBytes(1000, Gigabyte), other: NewBytes(1000, Terabyte), want: false},
		{name: "1000 TB != 1000 GB", bytes: NewBytes(1000, Terabyte), other: NewBytes(1000, Gigabyte), want: false},
		{name: "1000 TB == 1000 TB", bytes: NewBytes(1000, Terabyte), other: NewBytes(1000, Terabyte), want: true},
		{name: "1000 TB != 1000 PB", bytes: NewBytes(1000, Terabyte), other: NewBytes(1000, Petabyte), want: false},
		{name: "1000 PB != 1000 TB", bytes: NewBytes(1000, Petabyte), other: NewBytes(1000, Terabyte), want: false},
		{name: "1000 PB == 1000 PB", bytes: NewBytes(1000, Petabyte), other: NewBytes(1000, Petabyte), want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test the Equal() method
			got := test.bytes.Equal(test.other)
			if got != test.want {
				t.Errorf("Bytes.Equal() = %v, want %v", got, test.want)
			}
			// Test the == operator
			got = (test.bytes == test.other)
			if got != test.want {
				t.Errorf("Bytes == Bytes = %v, want %v", got, test.want)
			}
		})
	}
}

func TestBytesLessThan(t *testing.T) { //nolint:dupl // test ok
	tests := []struct {
		name  string
		bytes Bytes
		other Bytes
		want  bool
	}{
		{name: "1000 B < 1000 KB", bytes: NewBytes(1000, Byte), other: NewBytes(1000, Kilobyte), want: true},
		{name: "1000 KB < 1000 B", bytes: NewBytes(1000, Kilobyte), other: NewBytes(1000, Byte), want: false},

		{name: "1000 KB < 1000 MB", bytes: NewBytes(1000, Kilobyte), other: NewBytes(1000, Megabyte), want: true},
		{name: "1000 MB < 1000 KB", bytes: NewBytes(1000, Megabyte), other: NewBytes(1000, Kilobyte), want: false},

		{name: "1000 MB < 1000 GB", bytes: NewBytes(1000, Megabyte), other: NewBytes(1000, Gigabyte), want: true},
		{name: "1000 GB < 1000 MB", bytes: NewBytes(1000, Gigabyte), other: NewBytes(1000, Megabyte), want: false},

		{name: "1000 GB < 1000 TB", bytes: NewBytes(1000, Gigabyte), other: NewBytes(1000, Terabyte), want: true},
		{name: "1000 TB < 1000 GB", bytes: NewBytes(1000, Terabyte), other: NewBytes(1000, Gigabyte), want: false},

		{name: "1000 TB < 1000 PB", bytes: NewBytes(1000, Terabyte), other: NewBytes(1000, Petabyte), want: true},
		{name: "1000 PB < 1000 TB", bytes: NewBytes(1000, Petabyte), other: NewBytes(1000, Terabyte), want: false},

		{name: "1000 B < 1000 KiB", bytes: NewBytes(1000, Byte), other: NewBytes(1000, Kibibyte), want: true},
		{name: "1000 KiB < 1000 B", bytes: NewBytes(1000, Kibibyte), other: NewBytes(1000, Byte), want: false},

		{name: "1000 KiB < 1000 MiB", bytes: NewBytes(1000, Kibibyte), other: NewBytes(1000, Mebibyte), want: true},
		{name: "1000 MiB < 1000 KiB", bytes: NewBytes(1000, Mebibyte), other: NewBytes(1000, Kibibyte), want: false},

		{name: "1000 MiB < 1000 GiB", bytes: NewBytes(1000, Mebibyte), other: NewBytes(1000, Gibibyte), want: true},
		{name: "1000 GiB < 1000 MiB", bytes: NewBytes(1000, Gibibyte), other: NewBytes(1000, Mebibyte), want: false},

		{name: "1000 GiB < 1000 TiB", bytes: NewBytes(1000, Gibibyte), other: NewBytes(1000, Tebibyte), want: true},
		{name: "1000 TiB < 1000 GiB", bytes: NewBytes(1000, Tebibyte), other: NewBytes(1000, Gibibyte), want: false},

		{name: "1000 TiB < 1000 PiB", bytes: NewBytes(1000, Tebibyte), other: NewBytes(1000, Pebibyte), want: true},
		{name: "1000 PiB < 1000 TiB", bytes: NewBytes(1000, Pebibyte), other: NewBytes(1000, Tebibyte), want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.bytes.LessThan(test.other)
			if got != test.want {
				t.Errorf("Bytes.LessThan() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestBytesGreaterThan(t *testing.T) { //nolint:dupl // test ok
	tests := []struct {
		name  string
		bytes Bytes
		other Bytes
		want  bool
	}{
		{name: "1000 B > 1000 KB", bytes: NewBytes(1000, Byte), other: NewBytes(1000, Kilobyte), want: false},
		{name: "1000 KB > 1000 B", bytes: NewBytes(1000, Kilobyte), other: NewBytes(1000, Byte), want: true},

		{name: "1000 KB > 1000 MB", bytes: NewBytes(1000, Kilobyte), other: NewBytes(1000, Megabyte), want: false},
		{name: "1000 MB > 1000 KB", bytes: NewBytes(1000, Megabyte), other: NewBytes(1000, Kilobyte), want: true},

		{name: "1000 MB > 1000 GB", bytes: NewBytes(1000, Megabyte), other: NewBytes(1000, Gigabyte), want: false},
		{name: "1000 GB > 1000 MB", bytes: NewBytes(1000, Gigabyte), other: NewBytes(1000, Megabyte), want: true},

		{name: "1000 GB > 1000 TB", bytes: NewBytes(1000, Gigabyte), other: NewBytes(1000, Terabyte), want: false},
		{name: "1000 TB > 1000 GB", bytes: NewBytes(1000, Terabyte), other: NewBytes(1000, Gigabyte), want: true},

		{name: "1000 TB > 1000 PB", bytes: NewBytes(1000, Terabyte), other: NewBytes(1000, Petabyte), want: false},
		{name: "1000 PB > 1000 TB", bytes: NewBytes(1000, Petabyte), other: NewBytes(1000, Terabyte), want: true},

		{name: "1000 B > 1000 KiB", bytes: NewBytes(1000, Byte), other: NewBytes(1000, Kibibyte), want: false},
		{name: "1000 KiB > 1000 B", bytes: NewBytes(1000, Kibibyte), other: NewBytes(1000, Byte), want: true},

		{name: "1000 KiB > 1000 MiB", bytes: NewBytes(1000, Kibibyte), other: NewBytes(1000, Mebibyte), want: false},
		{name: "1000 MiB > 1000 KiB", bytes: NewBytes(1000, Mebibyte), other: NewBytes(1000, Kibibyte), want: true},

		{name: "1000 MiB > 1000 GiB", bytes: NewBytes(1000, Mebibyte), other: NewBytes(1000, Gibibyte), want: false},
		{name: "1000 GiB > 1000 MiB", bytes: NewBytes(1000, Gibibyte), other: NewBytes(1000, Mebibyte), want: true},

		{name: "1000 GiB > 1000 TiB", bytes: NewBytes(1000, Gibibyte), other: NewBytes(1000, Tebibyte), want: false},
		{name: "1000 TiB > 1000 GiB", bytes: NewBytes(1000, Tebibyte), other: NewBytes(1000, Gibibyte), want: true},

		{name: "1000 TiB > 1000 PiB", bytes: NewBytes(1000, Tebibyte), other: NewBytes(1000, Pebibyte), want: false},
		{name: "1000 PiB > 1000 TiB", bytes: NewBytes(1000, Pebibyte), other: NewBytes(1000, Tebibyte), want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.bytes.GreaterThan(test.other)
			if got != test.want {
				t.Errorf("Bytes.GreaterThan() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestBytesByteCountInUnit(t *testing.T) {
	tests := []struct {
		name  string
		bytes Bytes
		unit  BytesUnit
		want  *big.Float
	}{
		{name: "1000 B -> B", bytes: NewBytes(1000, Byte), unit: Byte, want: big.NewFloat(1000)},
		{name: "1000 B -> KB", bytes: NewBytes(1000, Byte), unit: Kilobyte, want: big.NewFloat(1)},
		{name: "1000 B -> MB", bytes: NewBytes(1000, Byte), unit: Megabyte, want: big.NewFloat(0.001)},
		{name: "1000 B -> GB", bytes: NewBytes(1000, Byte), unit: Gigabyte, want: big.NewFloat(0.000001)},
		{name: "1000 B -> TB", bytes: NewBytes(1000, Byte), unit: Terabyte, want: big.NewFloat(0.000000001)},
		{name: "1000 B -> PB", bytes: NewBytes(1000, Byte), unit: Petabyte, want: big.NewFloat(0.000000000001)},
		{name: "1000 B -> KiB", bytes: NewBytes(1000, Byte), unit: Kibibyte, want: big.NewFloat(0.9765625)},
		{name: "1000 B -> MiB", bytes: NewBytes(1000, Byte), unit: Mebibyte, want: big.NewFloat(0.00095367431640625)},
		{name: "1000 B -> GiB", bytes: NewBytes(1000, Byte), unit: Gibibyte, want: big.NewFloat(0.0000009313225746154785)},
		{name: "1000 B -> TiB", bytes: NewBytes(1000, Byte), unit: Tebibyte, want: big.NewFloat(0.0000000009094947017729282)},
		{name: "1000 B -> PiB", bytes: NewBytes(1000, Byte), unit: Pebibyte, want: big.NewFloat(0.0000000000008881784197001252)},

		{name: "1000 KB -> B", bytes: NewBytes(1000, Kilobyte), unit: Byte, want: big.NewFloat(1000000)},
		{name: "1000 KB -> KB", bytes: NewBytes(1000, Kilobyte), unit: Kilobyte, want: big.NewFloat(1000)},
		{name: "1000 KB -> MB", bytes: NewBytes(1000, Kilobyte), unit: Megabyte, want: big.NewFloat(1)},
		{name: "1000 KB -> GB", bytes: NewBytes(1000, Kilobyte), unit: Gigabyte, want: big.NewFloat(0.001)},
		{name: "1000 KB -> TB", bytes: NewBytes(1000, Kilobyte), unit: Terabyte, want: big.NewFloat(0.000001)},
		{name: "1000 KB -> PB", bytes: NewBytes(1000, Kilobyte), unit: Petabyte, want: big.NewFloat(0.000000001)},
		{name: "1000 KB -> KiB", bytes: NewBytes(1000, Kilobyte), unit: Kibibyte, want: big.NewFloat(976.5625)},
		{name: "1000 KB -> MiB", bytes: NewBytes(1000, Kilobyte), unit: Mebibyte, want: big.NewFloat(0.95367431640625)},
		{name: "1000 KB -> GiB", bytes: NewBytes(1000, Kilobyte), unit: Gibibyte, want: big.NewFloat(0.0009313225746154785)},
		{name: "1000 KB -> TiB", bytes: NewBytes(1000, Kilobyte), unit: Tebibyte, want: big.NewFloat(0.0000009094947017729282)},
		{name: "1000 KB -> PiB", bytes: NewBytes(1000, Kilobyte), unit: Pebibyte, want: big.NewFloat(0.0000000008881784197001252)},

		{name: "1000 MB -> B", bytes: NewBytes(1000, Megabyte), unit: Byte, want: big.NewFloat(1000000000)},
		{name: "1000 MB -> KB", bytes: NewBytes(1000, Megabyte), unit: Kilobyte, want: big.NewFloat(1000000)},
		{name: "1000 MB -> MB", bytes: NewBytes(1000, Megabyte), unit: Megabyte, want: big.NewFloat(1000)},
		{name: "1000 MB -> GB", bytes: NewBytes(1000, Megabyte), unit: Gigabyte, want: big.NewFloat(1)},
		{name: "1000 MB -> TB", bytes: NewBytes(1000, Megabyte), unit: Terabyte, want: big.NewFloat(0.001)},
		{name: "1000 MB -> PB", bytes: NewBytes(1000, Megabyte), unit: Petabyte, want: big.NewFloat(0.000001)},
		{name: "1000 MB -> KiB", bytes: NewBytes(1000, Megabyte), unit: Kibibyte, want: big.NewFloat(976562.5)},
		{name: "1000 MB -> MiB", bytes: NewBytes(1000, Megabyte), unit: Mebibyte, want: big.NewFloat(953.67431640625)},
		{name: "1000 MB -> GiB", bytes: NewBytes(1000, Megabyte), unit: Gibibyte, want: big.NewFloat(0.9313225746154785)},
		{name: "1000 MB -> TiB", bytes: NewBytes(1000, Megabyte), unit: Tebibyte, want: big.NewFloat(0.0009094947017729282)},
		{name: "1000 MB -> PiB", bytes: NewBytes(1000, Megabyte), unit: Pebibyte, want: big.NewFloat(0.0000008881784197001252)},

		{name: "1000 GB -> B", bytes: NewBytes(1000, Gigabyte), unit: Byte, want: big.NewFloat(1000000000000)},
		{name: "1000 GB -> KB", bytes: NewBytes(1000, Gigabyte), unit: Kilobyte, want: big.NewFloat(1000000000)},
		{name: "1000 GB -> MB", bytes: NewBytes(1000, Gigabyte), unit: Megabyte, want: big.NewFloat(1000000)},
		{name: "1000 GB -> GB", bytes: NewBytes(1000, Gigabyte), unit: Gigabyte, want: big.NewFloat(1000)},
		{name: "1000 GB -> TB", bytes: NewBytes(1000, Gigabyte), unit: Terabyte, want: big.NewFloat(1)},
		{name: "1000 GB -> PB", bytes: NewBytes(1000, Gigabyte), unit: Petabyte, want: big.NewFloat(0.001)},
		{name: "1000 GB -> KiB", bytes: NewBytes(1000, Gigabyte), unit: Kibibyte, want: big.NewFloat(976562500)},
		{name: "1000 GB -> MiB", bytes: NewBytes(1000, Gigabyte), unit: Mebibyte, want: big.NewFloat(953674.31640625)},
		{name: "1000 GB -> GiB", bytes: NewBytes(1000, Gigabyte), unit: Gibibyte, want: big.NewFloat(931.3225746154785)},
		{name: "1000 GB -> TiB", bytes: NewBytes(1000, Gigabyte), unit: Tebibyte, want: big.NewFloat(0.9094947017729282)},
		{name: "1000 GB -> PiB", bytes: NewBytes(1000, Gigabyte), unit: Pebibyte, want: big.NewFloat(0.0008881784197001252)},

		{name: "1000 TB -> B", bytes: NewBytes(1000, Terabyte), unit: Byte, want: big.NewFloat(1000000000000000)},
		{name: "1000 TB -> KB", bytes: NewBytes(1000, Terabyte), unit: Kilobyte, want: big.NewFloat(1000000000000)},
		{name: "1000 TB -> MB", bytes: NewBytes(1000, Terabyte), unit: Megabyte, want: big.NewFloat(1000000000)},
		{name: "1000 TB -> GB", bytes: NewBytes(1000, Terabyte), unit: Gigabyte, want: big.NewFloat(1000000)},
		{name: "1000 TB -> TB", bytes: NewBytes(1000, Terabyte), unit: Terabyte, want: big.NewFloat(1000)},
		{name: "1000 TB -> PB", bytes: NewBytes(1000, Terabyte), unit: Petabyte, want: big.NewFloat(1)},
		{name: "1000 TB -> KiB", bytes: NewBytes(1000, Terabyte), unit: Kibibyte, want: big.NewFloat(976562500000)},
		{name: "1000 TB -> MiB", bytes: NewBytes(1000, Terabyte), unit: Mebibyte, want: big.NewFloat(953674316.40625)},
		{name: "1000 TB -> GiB", bytes: NewBytes(1000, Terabyte), unit: Gibibyte, want: big.NewFloat(931322.5746154785)},
		{name: "1000 TB -> TiB", bytes: NewBytes(1000, Terabyte), unit: Tebibyte, want: big.NewFloat(909.4947017729282)},
		{name: "1000 TB -> PiB", bytes: NewBytes(1000, Gigabyte), unit: Pebibyte, want: big.NewFloat(0.0008881784197001252)},

		{name: "1000 PB -> B", bytes: NewBytes(1000, Petabyte), unit: Byte, want: big.NewFloat(1000000000000000000)},
		{name: "1000 PB -> KB", bytes: NewBytes(1000, Petabyte), unit: Kilobyte, want: big.NewFloat(1000000000000000)},
		{name: "1000 PB -> MB", bytes: NewBytes(1000, Petabyte), unit: Megabyte, want: big.NewFloat(1000000000000)},
		{name: "1000 PB -> GB", bytes: NewBytes(1000, Petabyte), unit: Gigabyte, want: big.NewFloat(1000000000)},
		{name: "1000 PB -> TB", bytes: NewBytes(1000, Petabyte), unit: Terabyte, want: big.NewFloat(1000000)},
		{name: "1000 PB -> PB", bytes: NewBytes(1000, Petabyte), unit: Petabyte, want: big.NewFloat(1000)},
		{name: "1000 PB -> KiB", bytes: NewBytes(1000, Petabyte), unit: Kibibyte, want: big.NewFloat(976562500000000)},
		{name: "1000 PB -> MiB", bytes: NewBytes(1000, Petabyte), unit: Mebibyte, want: big.NewFloat(953674316406.25)},
		{name: "1000 PB -> GiB", bytes: NewBytes(1000, Petabyte), unit: Gibibyte, want: big.NewFloat(931322574.6154785)},
		{name: "1000 PB -> TiB", bytes: NewBytes(1000, Petabyte), unit: Tebibyte, want: big.NewFloat(909494.7017729282)},
		{name: "1000 PB -> PiB", bytes: NewBytes(1000, Petabyte), unit: Pebibyte, want: big.NewFloat(888.1784197001252)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.bytes.ByteCountInUnit(test.unit)
			if got.Cmp(test.want) != 0 {
				t.Errorf("Bytes.ByteCountInUnit() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestBytesByteCountInUnitInt64(t *testing.T) {
	tests := []struct {
		name    string
		bytes   Bytes
		unit    BytesUnit
		want    int64
		wantErr error
	}{
		{name: "2048 MiB -> GiB", bytes: NewBytes(2048, Mebibyte), unit: Gibibyte, want: 2, wantErr: nil},
		{name: "2048 EiB -> B", bytes: NewBytes(2048, Exbibyte), unit: Byte, want: 0, wantErr: ErrBytesNotAnInt64},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.bytes.ByteCountInUnitInt64(test.unit)
			if err != nil {
				if test.wantErr == nil {
					t.Fatalf("Bytes.ByteCountInUnitInt64() error = %v, want nil", err)
				}
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("Bytes.ByteCountInUnitInt64() error = %v, want %v", err, test.wantErr)
				}
			} else if got != test.want {
				t.Errorf("Bytes.ByteCountInUnitInt64() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestBytesByteCountInUnitInt32(t *testing.T) {
	tests := []struct {
		name    string
		bytes   Bytes
		unit    BytesUnit
		want    int32
		wantErr error
	}{
		{name: "2048 MiB -> GiB", bytes: NewBytes(2048, Mebibyte), unit: Gibibyte, want: 2, wantErr: nil},
		{name: "2048 EiB -> B", bytes: NewBytes(2048, Exbibyte), unit: Byte, want: 0, wantErr: ErrBytesNotAnInt64},
		{name: "2048 EiB -> KB", bytes: NewBytes(2048, Exbibyte), unit: Kilobyte, want: 0, wantErr: ErrBytesNotAnInt32},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.bytes.ByteCountInUnitInt32(test.unit)
			if err != nil {
				if test.wantErr == nil {
					t.Fatalf("Bytes.ByteCountInUnitInt32() error = %v, want nil", err)
				}
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("Bytes.ByteCountInUnitInt32() error = %v, want %v", err, test.wantErr)
				}
			} else if got != test.want {
				t.Errorf("Bytes.ByteCountInUnitInt32() = %v, want %v", got, test.want)
			}
		})
	}
}
