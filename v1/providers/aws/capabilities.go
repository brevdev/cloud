package v1

import (
	v1 "github.com/brevdev/cloud/v1"
)

func GetAWSCapabilities() v1.Capabilities {
	return v1.Capabilities{
		v1.CapabilityVPC,
	}
}
