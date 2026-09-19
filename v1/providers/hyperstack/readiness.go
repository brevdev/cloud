package hyperstack

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	virtualmachine "github.com/NexGenCloud/hyperstack-sdk-go/lib/virtual_machine"
)

const (
	readinessMarker      = "BREV_CLOUD_READY_V1"
	consoleLogLineCount  = 200
	consoleLogPolls      = 5
	consoleLogPollPeriod = 500 * time.Millisecond
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

      [Install]
      WantedBy=multi-user.target
runcmd:
  - [systemctl, daemon-reload]
  - [systemctl, enable, brev-cloud-ready.service]
  - [systemctl, start, --no-block, brev-cloud-ready.service]
`
)

func (c *HyperstackClient) consoleReady(ctx context.Context, instanceID int) (bool, error) {
	requestID, err := c.requestConsoleLogs(ctx, instanceID)
	if err != nil || requestID == 0 {
		return false, err
	}

	for poll := 0; poll < consoleLogPolls; poll++ {
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
		case http.StatusOK:
			if response.JSON200 != nil && response.JSON200.Logs != nil {
				return strings.Contains(*response.JSON200.Logs, readinessMarker), nil
			}
		case http.StatusAccepted, http.StatusBadRequest:
			// Hyperstack returns either status while the asynchronous request is processing.
		case http.StatusUnauthorized, http.StatusForbidden:
			return false, responseError("get virtual machine console logs", response.StatusCode(), response.Body, nil)
		default:
			return false, nil
		}

		if poll < consoleLogPolls-1 {
			select {
			case <-ctx.Done():
				return false, ctx.Err()
			case <-time.After(consoleLogPollPeriod):
			}
		}
	}
	return false, nil
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
