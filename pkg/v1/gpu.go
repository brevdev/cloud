package v1

const (
	GPUA10  = "A10"
	GPUA10G = "A10G"
	GPUA100 = "A100"

	GPUB200 = "B200"

	GPUH100 = "H100"
	GPUH200 = "H200"

	GPUL4   = "L4"
	GPUL40  = "L40"
	GPUL40S = "L40S"

	GPUM60  = "M60"
	GPUP4   = "P4"
	GPUT4   = "T4"
	GPUV100 = "V100"
)

func AllGPUNames() []string {
	return []string{
		GPUA10, GPUA10G, GPUA100,
		GPUB200,
		GPUH100, GPUH200,
		GPUL4, GPUL40, GPUL40S,
		GPUM60, GPUP4, GPUT4, GPUV100,
	}
}

func NormalizeGPUName(name string) string {
	normalized := name
	prefixes := []string{"NVIDIA ", "Tesla ", "GeForce "}
	for _, prefix := range prefixes {
		if len(normalized) > len(prefix) && normalized[:len(prefix)] == prefix {
			normalized = normalized[len(prefix):]
			break
		}
	}

	for _, gpuName := range AllGPUNames() {
		if normalized == gpuName {
			return gpuName
		}
	}

	return name
}
