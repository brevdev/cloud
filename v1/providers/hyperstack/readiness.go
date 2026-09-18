package hyperstack

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	virtualmachine "github.com/NexGenCloud/hyperstack-sdk-go/lib/virtual_machine"
)

const (
	readinessLabel       = "brev-readiness-v1"
	readinessMarker      = "BREV_CLOUD_READY_V1"
	consoleLogLineCount  = 200
	readinessCloudConfig = `#cloud-config
write_files:
  - path: /etc/systemd/system/brev-cloud-ready.service
    permissions: '0644'
    content: |
      [Unit]
      Description=Brev cloud instance readiness signal
      After=cloud-final.service

      [Service]
      Type=oneshot
      ExecStart=/bin/sh -c 'while systemctl list-jobs --no-legend --no-pager | grep -v brev-cloud-ready.service | grep -q .; do sleep 2; done; printf "BREV_CLOUD_READY_V1\n" > /dev/ttyS0'
runcmd:
  - [systemctl, daemon-reload]
  - [systemctl, start, --no-block, brev-cloud-ready.service]
`
)

func requiresConsoleReadiness(labels *[]string) bool {
	return labels != nil && containsString(*labels, readinessLabel)
}

func containsString(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}

func (c *HyperstackClient) consoleReady(ctx context.Context, instanceID int) (bool, error) {
	requestID, ready := c.readinessState(instanceID)
	if ready {
		return true, nil
	}
	if requestID == 0 {
		var err error
		requestID, err = c.requestConsoleLogs(ctx, instanceID)
		if err != nil {
			return false, err
		}
		if requestID == 0 {
			return false, nil
		}
		c.setLogRequest(instanceID, requestID)
	}

	response, err := c.virtualMachines.GetVMLogsWithResponse(ctx, instanceID, &virtualmachine.GetVMLogsParams{
		RequestId: requestID,
	})
	if err != nil {
		if ctx.Err() != nil {
			return false, ctx.Err()
		}
		return false, nil
	}
	switch response.StatusCode() {
	case http.StatusAccepted, http.StatusBadRequest:
		// The request endpoint is asynchronous. The retrieval endpoint documents
		// 400, rather than 202, while a valid request is not ready to read yet.
		return false, nil
	case http.StatusOK:
		c.clearLogRequest(instanceID)
		if response.JSON200 == nil || response.JSON200.Logs == nil {
			return false, nil
		}
		if !strings.Contains(*response.JSON200.Logs, readinessMarker) {
			return false, nil
		}
		c.setReady(instanceID)
		return true, nil
	case http.StatusUnauthorized, http.StatusForbidden:
		return false, responseError("get virtual machine console logs", response.StatusCode(), response.Body, nil)
	default:
		c.clearLogRequest(instanceID)
		return false, nil
	}
}

func (c *HyperstackClient) requestConsoleLogs(ctx context.Context, instanceID int) (int, error) {
	lineCount := consoleLogLineCount
	response, err := c.virtualMachines.RequestVMLogsWithResponse(ctx, instanceID, virtualmachine.RequestInstanceLogsPayload{
		Length: &lineCount,
	})
	if err != nil {
		if ctx.Err() != nil {
			return 0, ctx.Err()
		}
		return 0, nil
	}
	if response.StatusCode() == http.StatusUnauthorized || response.StatusCode() == http.StatusForbidden {
		return 0, responseError("request virtual machine console logs", response.StatusCode(), response.Body, nil)
	}
	if response.StatusCode() != http.StatusOK && response.StatusCode() != http.StatusAccepted {
		return 0, nil
	}

	var payload struct {
		RequestID int `json:"request_id"`
	}
	if err := json.Unmarshal(response.Body, &payload); err != nil {
		return 0, nil
	}
	if payload.RequestID <= 0 {
		return 0, nil
	}
	return payload.RequestID, nil
}

func (c *HyperstackClient) readinessState(instanceID int) (requestID int, ready bool) {
	c.readinessMu.Lock()
	defer c.readinessMu.Unlock()
	_, ready = c.readyInstances[instanceID]
	return c.logRequests[instanceID], ready
}

func (c *HyperstackClient) setLogRequest(instanceID, requestID int) {
	c.readinessMu.Lock()
	defer c.readinessMu.Unlock()
	if c.logRequests == nil {
		c.logRequests = make(map[int]int)
	}
	c.logRequests[instanceID] = requestID
}

func (c *HyperstackClient) clearLogRequest(instanceID int) {
	c.readinessMu.Lock()
	defer c.readinessMu.Unlock()
	delete(c.logRequests, instanceID)
}

func (c *HyperstackClient) setReady(instanceID int) {
	c.readinessMu.Lock()
	defer c.readinessMu.Unlock()
	if c.readyInstances == nil {
		c.readyInstances = make(map[int]struct{})
	}
	delete(c.logRequests, instanceID)
	c.readyInstances[instanceID] = struct{}{}
}
