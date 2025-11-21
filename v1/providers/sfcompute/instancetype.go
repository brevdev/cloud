package v1

import (
	"fmt"
)

func (c *SFCClient) getInstanceTypeID(region string) string {
	return fmt.Sprintf("h100v_%v", region)
}
