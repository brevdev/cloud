package v1

import (
	openapi "github.com/brevdev/cloud/internal/shadeform/gen/shadeform"
)

type Configuration struct {
	AllowedInstanceTypes map[openapi.Cloud]map[string]bool
}

func (c *Configuration) isAllowed(cloud openapi.Cloud, shadeInstanceType string) bool {

	allowedClouds, found := c.AllowedInstanceTypes[cloud]
	if !found {
		return false
	}

	_, found = allowedClouds[shadeInstanceType]
	if !found {
		return false
	}

	return found
}
