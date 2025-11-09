package v1

import (
	"encoding/json"
	"errors"
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
		{name: "1000 EB", value: 1000, unit: Exabyte, want: NewBytes(1000, Exabyte), wantErr: nil},
		{name: "1000 KiB", value: 1000, unit: Kibibyte, want: NewBytes(1000, Kibibyte), wantErr: nil},
		{name: "1000 MiB", value: 1000, unit: Mebibyte, want: NewBytes(1000, Mebibyte), wantErr: nil},
		{name: "1000 GiB", value: 1000, unit: Gibibyte, want: NewBytes(1000, Gibibyte), wantErr: nil},
		{name: "1000 TiB", value: 1000, unit: Tebibyte, want: NewBytes(1000, Tebibyte), wantErr: nil},
		{name: "1000 PiB", value: 1000, unit: Pebibyte, want: NewBytes(1000, Pebibyte), wantErr: nil},
		{name: "1000 EiB", value: 1000, unit: Exbibyte, want: NewBytes(1000, Exbibyte), wantErr: nil},
		{name: "Negative value", value: -1000, unit: Byte, want: zeroBytes, wantErr: ErrNegativeValue},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bytes := NewBytes(test.value, test.unit)
			if bytes != test.want {
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
		{name: "1000 EB", bytes: NewBytes(1000, Exabyte), want: `{"value":1000,"unit":"EB"}`, wantErr: nil},
		{name: "1000 KiB", bytes: NewBytes(1000, Kibibyte), want: `{"value":1000,"unit":"KiB"}`, wantErr: nil},
		{name: "1000 MiB", bytes: NewBytes(1000, Mebibyte), want: `{"value":1000,"unit":"MiB"}`, wantErr: nil},
		{name: "1000 GiB", bytes: NewBytes(1000, Gibibyte), want: `{"value":1000,"unit":"GiB"}`, wantErr: nil},
		{name: "1000 TiB", bytes: NewBytes(1000, Tebibyte), want: `{"value":1000,"unit":"TiB"}`, wantErr: nil},
		{name: "1000 PiB", bytes: NewBytes(1000, Pebibyte), want: `{"value":1000,"unit":"PiB"}`, wantErr: nil},
		{name: "1000 EiB", bytes: NewBytes(1000, Exbibyte), want: `{"value":1000,"unit":"EiB"}`, wantErr: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			json, err := json.Marshal(test.bytes)
			if test.wantErr != nil {
				if err == nil {
					t.Fatalf("json.Marshal() error = nil, want %v", test.wantErr)
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
		{name: "1000 B", json: `{"value":1000,"unit":"B"}`, want: NewBytes(1000, Byte), wantErr: nil},
		{name: "1000 KB", json: `{"value":1000,"unit":"KB"}`, want: NewBytes(1000, Kilobyte), wantErr: nil},
		{name: "1000 MB", json: `{"value":1000,"unit":"MB"}`, want: NewBytes(1000, Megabyte), wantErr: nil},
		{name: "1000 GB", json: `{"value":1000,"unit":"GB"}`, want: NewBytes(1000, Gigabyte), wantErr: nil},
		{name: "1000 TB", json: `{"value":1000,"unit":"TB"}`, want: NewBytes(1000, Terabyte), wantErr: nil},
		{name: "1000 PB", json: `{"value":1000,"unit":"PB"}`, want: NewBytes(1000, Petabyte), wantErr: nil},
		{name: "1000 EB", json: `{"value":1000,"unit":"EB"}`, want: NewBytes(1000, Exabyte), wantErr: nil},
		{name: "1000 KiB", json: `{"value":1000,"unit":"KiB"}`, want: NewBytes(1000, Kibibyte), wantErr: nil},
		{name: "1000 MiB", json: `{"value":1000,"unit":"MiB"}`, want: NewBytes(1000, Mebibyte), wantErr: nil},
		{name: "1000 GiB", json: `{"value":1000,"unit":"GiB"}`, want: NewBytes(1000, Gibibyte), wantErr: nil},
		{name: "1000 TiB", json: `{"value":1000,"unit":"TiB"}`, want: NewBytes(1000, Tebibyte), wantErr: nil},
		{name: "1000 PiB", json: `{"value":1000,"unit":"PiB"}`, want: NewBytes(1000, Pebibyte), wantErr: nil},
		{name: "1000 EiB", json: `{"value":1000,"unit":"EiB"}`, want: NewBytes(1000, Exbibyte), wantErr: nil},

		{name: "Negative value", json: `{"value":-1000,"unit":"B"}`, want: zeroBytes, wantErr: ErrNegativeValue},
		{name: "Empty unit", json: `{"value":1000,"unit":""}`, want: zeroBytes, wantErr: ErrEmptyUnit},
		{name: "Invalid unit", json: `{"value":1000,"unit":"invalid"}`, want: zeroBytes, wantErr: ErrInvalidUnit},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var bytes Bytes
			err := json.Unmarshal([]byte(test.json), &bytes)
			if test.wantErr != nil {
				if err == nil {
					t.Fatalf("json.Unmarshal() error = nil, want %v", test.wantErr)
				}
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("json.Unmarshal() error = %v, want %v", err, test.wantErr)
				}
			} else if bytes != test.want {
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
		{name: "Exabyte", bytes: NewBytes(1000, Exabyte), want: "1000 EB"},
		{name: "Kibibyte", bytes: NewBytes(1000, Kibibyte), want: "1000 KiB"},
		{name: "Mebibyte", bytes: NewBytes(1000, Mebibyte), want: "1000 MiB"},
		{name: "Gibibyte", bytes: NewBytes(1000, Gibibyte), want: "1000 GiB"},
		{name: "Tebibyte", bytes: NewBytes(1000, Tebibyte), want: "1000 TiB"},
		{name: "Pebibyte", bytes: NewBytes(1000, Pebibyte), want: "1000 PiB"},
		{name: "Exbibyte", bytes: NewBytes(1000, Exbibyte), want: "1000 EiB"},
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
