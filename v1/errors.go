package v1

import "errors"

var (
	ErrInsufficientResources = errors.New("zone has insufficient resources to fulfill the request, InsufficientCapacity")
	ErrOutOfQuota            = errors.New("out of quota in the region fulfill the request, InsufficientQuota")
	ErrImageNotFound         = errors.New("image not found")
	ErrDuplicateFirewallRule = errors.New("duplicate firewall rule")
	ErrInstanceNotFound      = errors.New("instance not found")
	ErrServiceUnavailable    = errors.New("api is temporarily unavailable")
)
