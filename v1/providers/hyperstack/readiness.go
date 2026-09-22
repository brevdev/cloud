package hyperstack

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	virtualmachine "github.com/NexGenCloud/hyperstack-sdk-go/lib/virtual_machine"
	v1 "github.com/brevdev/cloud/v1"
)

const (
	readinessMarker      = "BREV_CLOUD_READY_V1"
	consoleLogLineCount  = 200
	consoleLogPolls      = 5
	consoleLogPollPeriod = 500 * time.Millisecond
	readinessLabel       = "brev-cloud-ready"
	readinessCloudConfig = `#cloud-config
write_files:
  - path: /etc/systemd/system/brev-cloud-ready.service
    permissions: '0644'
    content: |
      [Unit]
      Description=Brev cloud instance readiness signal
      Wants=network-online.target
      After=network-online.target

      [Service]
      Type=oneshot
      ExecStart=/bin/sh -c 'printf "BREV_CLOUD_READY_V1\n" > /dev/ttyS0'

      [Install]
      WantedBy=multi-user.target
runcmd:
  - [systemctl, daemon-reload]
  - [systemctl, enable, brev-cloud-ready.service]
  - [systemctl, start, --no-block, brev-cloud-ready.service]
`
)

// Hyperstack is somewhat silly with its statuses. The VM instance itself can achieve "readiness," with the API indicating as such,
// but the operating system itself may not have fully booted yet. In order to get around this, at VM provision time we create a oneshot
// systemd service that write a "BREV_CLOUD_READY_V1" marker to the serial console. We then poll the console for this marker, and if it's found,
// we consider the VM ready. The 'vmOperatingSystemReportsReady' therefore returns true if the marker is found in the console logs. In order
// to speed up this check, this function has a side-effect of adding the readiness label to the VM if the marker is found in the console logs.
// See: https://docs.hyperstack.cloud/docs/virtual-machines/virtual-machine-features/#managing-virtual-machines
func (c *HyperstackClient) vmOperatingSystemReportsReady(ctx context.Context, providerInstance virtualmachine.InstanceFields) (bool, error) {
	// If the VM has the readiness label, we consider it ready.
	if readinessLabelExists(providerInstance) {
		return true, nil
	}

	instanceID := intValue(providerInstance.Id)
	requestID, err := c.requestConsoleLogs(ctx, instanceID)
	if err != nil || requestID == 0 {
		return false, err
	}

	for poll := 0; poll < consoleLogPolls; poll++ {
		// Fetch the logs from the console.
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
				readinessMarkerFound := strings.Contains(*response.JSON200.Logs, readinessMarker)
				c.updateReadinessLabel(ctx, providerInstance, readinessMarkerFound)
				return readinessMarkerFound, nil
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

func readinessLabelExists(providerInstance virtualmachine.InstanceFields) bool {
	if providerInstance.Labels == nil {
		return false
	}
	return slices.Contains(*providerInstance.Labels, readinessLabel)
}

func (c *HyperstackClient) updateReadinessLabel(ctx context.Context, providerInstance virtualmachine.InstanceFields, readinessMarkerFound bool) {
	if !readinessMarkerFound {
		return
	}

	response, err := c.virtualMachines.AddVMLabelWithResponse(ctx, *providerInstance.Id, virtualmachine.AddVMLabelJSONRequestBody{
		Labels: &[]string{readinessLabel},
	})
	if err != nil {
		c.logger.Warn(ctx, fmt.Sprintf("error adding virtual machine label: %v", err), v1.LogField("instance_id", *providerInstance.Id))
		return
	}

	if response.StatusCode() != http.StatusOK {
		c.logger.Warn(ctx, fmt.Sprintf("error adding virtual machine label: %v", response.StatusCode()), v1.LogField("instance_id", *providerInstance.Id))
	}
}
