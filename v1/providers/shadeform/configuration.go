package v1

type Configuration struct {
	AllowedInstanceTypes map[string]map[string]bool
}

func (c *Configuration) isAllowed(cloud string, shadeInstanceType string) bool {
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
