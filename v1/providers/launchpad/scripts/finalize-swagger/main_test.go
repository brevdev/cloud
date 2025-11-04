package main

import (
	"testing"

	yaml "gopkg.in/yaml.v3"
)

func TestAddMissingEnumValues(t *testing.T) {
	actual := stringToYamlNode(t, `
components:
  schemas:
    CpuManufacturerEnum:
      enum:
      - amd
      - arm
      - intel
`)
	// System under test
	err := AddMissingEnumValues(actual)
	if err != nil {
		t.Fatalf("Error adding missing enum values: %v", err)
	}

	expected := stringToYamlNode(t, `
components:
  schemas:
    CpuManufacturerEnum:
      enum:
      - amd
      - arm
      - intel
      x-enum-varnames:
      - CpuManufacturerAMD
      - CpuManufacturerARM
      - CpuManufacturerIntel
`)

	aMarshaled, _ := yaml.Marshal(actual)
	eMarshaled, _ := yaml.Marshal(expected)
	if string(aMarshaled) != string(eMarshaled) {
		t.Fatalf("Expected\n%s\nGot\n%s", string(eMarshaled), string(aMarshaled))
	}
}

func TestConvertStringToOneOf(t *testing.T) {
	actual := stringToYamlNode(t, `
components:
  schemas:
    Foo:
      properties:
        bar:
          type: string
          format: uuid
        baz:
          type: number
`)
	// System under test
	err := AddOneOfEntries(actual)
	if err != nil {
		t.Fatalf("Error adding one of entries: %v", err)
	}

	expected := stringToYamlNode(t, `
components:
  schemas:
    Foo:
      properties:
        bar:
          oneOf:
            - type:
              - string
              format: uuid
            - $ref: "#/components/schemas/Bar"
          x-go-json-tag: "bar,omitempty"
        baz:
          type: number
`)

	aMarshaled, _ := yaml.Marshal(actual)
	eMarshaled, _ := yaml.Marshal(expected)
	if string(aMarshaled) != string(eMarshaled) {
		t.Fatalf("Expected\n%s\nGot\n%s", string(eMarshaled), string(aMarshaled))
	}
}

func TestConvertArrayToOneOf(t *testing.T) {
	actual := stringToYamlNode(t, `
components:
  schemas:
    Foo:
      properties:
        bars:
          type: array
          items:
            type: string
            format: uuid
        bazes:
          type: array
          items:
            type: number
`)
	// System under test
	err := AddOneOfEntries(actual)
	if err != nil {
		t.Fatalf("Error adding one of entries: %v", err)
	}

	expected := stringToYamlNode(t, `
components:
  schemas:
    Foo:
      properties:
        bars:
          type: array
          items:
            oneOf:
              - type:
                  - string
                format: uuid
              - $ref: "#/components/schemas/Bar"
            x-go-json-tag: "bar,omitempty"
          x-go-json-tag: "bars,omitempty"
        bazes:
          type: array
          items:
            type: number
`)

	aMarshaled, _ := yaml.Marshal(actual)
	eMarshaled, _ := yaml.Marshal(expected)
	if string(aMarshaled) != string(eMarshaled) {
		t.Fatalf("Expected\n%s\nGot\n%s", string(eMarshaled), string(aMarshaled))
	}
}

func TestRemoveBlankEnumValues(t *testing.T) {
	actual := stringToYamlNode(t, `
components:
  schemas:
    Foo:
      properties:
        platform:
          oneOf:
            - $ref: "#/components/schemas/PlatformEnum"
            - $ref: "#/components/schemas/BlankEnum"
            - $ref: "#/components/schemas/NullEnum"
`)
	// System under test
	err := FinalizeNullnessOption(actual)
	if err != nil {
		t.Fatalf("Error finalizing nullness option: %v", err)
	}

	expected := stringToYamlNode(t, `
components:
  schemas:
    Foo:
      properties:
        platform:
          oneOf:
            - $ref: "#/components/schemas/PlatformEnum"
            - type: "null"
`)

	aMarshaled, _ := yaml.Marshal(actual)
	eMarshaled, _ := yaml.Marshal(expected)
	if string(aMarshaled) != string(eMarshaled) {
		t.Fatalf("Expected\n%s\nGot\n%s", string(eMarshaled), string(aMarshaled))
	}
}

func stringToYamlNode(t *testing.T, yamlContent string) *yaml.Node {
	var yamlNode yaml.Node
	if err := yaml.Unmarshal([]byte(yamlContent), &yamlNode); err != nil {
		t.Fatalf("Error unmarshalling yaml: %v", err)
	}
	return &yamlNode
}

func TestUpperCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"cpu_manufacturer", "CpuManufacturer"},
		{"cpu-manufacturer", "CpuManufacturer"},
		{"cpu_manufacturer_arm", "CpuManufacturerArm"},
	}

	for _, test := range tests {
		actual := upperCamelCase(test.input)
		if actual != test.expected {
			t.Fatalf("Expected %s, got %s", test.expected, actual)
		}
	}
}

func TestNewEnumName(t *testing.T) {
	tests := []struct {
		prefix    string
		enumValue string
		expected  string
	}{
		{"CpuManufacturer", "amd", "CpuManufacturerAMD"},
		{"CpuManufacturer", "arm", "CpuManufacturerARM"},
		{"CpuManufacturer", "intel", "CpuManufacturerIntel"},
		{"CpuManufacturer", "amd_foo", "CpuManufacturerAMDFoo"},
		{"CpuManufacturer", "arm_foo", "CpuManufacturerARMFoo"},
		{"CpuManufacturer", "intel_foo", "CpuManufacturerIntelFoo"},
		{"CpuManufacturer", "arm-foo", "CpuManufacturerARMFooLegacy"},
		{"CpuManufacturer", "amd-foo", "CpuManufacturerAMDFooLegacy"},
		{"CpuManufacturer", "intel-foo", "CpuManufacturerIntelFooLegacy"},
		{"CpuManufacturer", "nvme", "CpuManufacturerNVMe"},
		{"CpuManufacturer", "pcie", "CpuManufacturerPCIe"},
	}
	for _, test := range tests {
		actual := newEnumName(test.prefix, test.enumValue)
		if actual != test.expected {
			t.Fatalf("Expected %s, got %s", test.expected, actual)
		}
	}
}
