package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultAPIURL = "https://api.sfcompute.com"
	brevAPIPath   = "/integrations/brev/v1"
)

type apiClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type createInstanceRequest struct {
	Name                    *string           `json:"name,omitempty"`
	Pool                    string            `json:"pool"`
	Image                   string            `json:"image"`
	InstanceSKU             string            `json:"instance_sku"`
	CloudInitUserData       *string           `json:"cloud_init_user_data,omitempty"`
	Tags                    map[string]string `json:"tags,omitempty"`
	PreviewEnableInfiniband bool              `json:"_preview_enable_infiniband"`
	EnablePublicIPv4        bool              `json:"enable_public_ipv4,omitempty"`
	Firewall                string            `json:"firewall,omitempty"`
}

type instanceStatus string

const (
	instanceStatusAwaitingAllocation instanceStatus = "awaiting_allocation"
	instanceStatusRunning            instanceStatus = "running"
	instanceStatusTerminated         instanceStatus = "terminated"
	instanceStatusFailed             instanceStatus = "failed"
)

type instanceSKUSummary struct {
	ID string `json:"id"`
}

type instanceResponse struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	Status           instanceStatus      `json:"status"`
	InstanceSKU      *instanceSKUSummary `json:"instance_sku"`
	CreatedAt        int64               `json:"created_at"`
	Tags             map[string]string   `json:"tags"`
	EnablePublicIPv4 bool                `json:"enable_public_ipv4"`
	Firewall         string              `json:"firewall"`
	PublicIP         string              `json:"public_ip"`
}

type listInstancesResponse struct {
	Cursor  *string            `json:"cursor,omitempty"`
	HasMore bool               `json:"has_more"`
	Data    []instanceResponse `json:"data"`
}

type instanceSSHInfo struct {
	Hostname string `json:"hostname"`
	Port     int64  `json:"port"`
}

func (i *instanceSSHInfo) GetHostname() string {
	if i == nil {
		return ""
	}
	return i.Hostname
}

func (i *instanceSSHInfo) GetPort() int64 {
	if i == nil {
		return 0
	}
	return i.Port
}

type scheduleEntry struct {
	StartAt   int64  `json:"start_at"`
	EndAt     *int64 `json:"end_at"`
	NodeCount int    `json:"node_count"`
}

type allocationSchedule struct {
	ByInstanceSKU map[string][]scheduleEntry `json:"by_instance_sku"`
}

type poolResponse struct {
	AllocationSchedule allocationSchedule `json:"allocation_schedule"`
	PublicIPv4SKUs     *[]string          `json:"public_ipv4_skus,omitempty"`
}

type apiError struct {
	statusCode int
	body       string
}

func (e *apiError) Error() string {
	return fmt.Sprintf("SFCompute API request failed: status %d: %s", e.statusCode, e.body)
}

func newAPIClient(apiKey string) *apiClient {
	return &apiClient{
		apiKey:     apiKey,
		baseURL:    defaultAPIURL,
		httpClient: &http.Client{Timeout: 60 * time.Second},
	}
}

func (c *apiClient) createInstance(ctx context.Context, request createInstanceRequest) (*instanceResponse, error) {
	var response instanceResponse
	if err := c.do(ctx, http.MethodPost, "/instances", nil, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *apiClient) getInstance(ctx context.Context, id string) (*instanceResponse, error) {
	var response instanceResponse
	if err := c.do(ctx, http.MethodGet, "/instances/"+url.PathEscape(id), nil, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *apiClient) listInstances(ctx context.Context, workspace, pool string) (*listInstancesResponse, error) {
	query := url.Values{"workspace": {workspace}, "pool": {pool}, "limit": {"200"}}
	var response listInstancesResponse
	for {
		var page listInstancesResponse
		if err := c.do(ctx, http.MethodGet, "/instances", query, nil, &page); err != nil {
			return nil, err
		}
		response.Data = append(response.Data, page.Data...)
		response.Cursor = page.Cursor
		response.HasMore = page.HasMore

		if !page.HasMore {
			return &response, nil
		}
		if page.Cursor == nil || *page.Cursor == "" {
			return nil, fmt.Errorf("list instances response has_more without a cursor")
		}
		if query.Get("starting_after") == *page.Cursor {
			return nil, fmt.Errorf("list instances response repeated cursor %q", *page.Cursor)
		}
		query.Set("starting_after", *page.Cursor)
	}
}

func (c *apiClient) terminateInstance(ctx context.Context, id string) error {
	_, err := c.terminateInstanceWithResponse(ctx, id)
	return err
}

func (c *apiClient) terminateInstanceWithResponse(ctx context.Context, id string) (*instanceResponse, error) {
	var response instanceResponse
	if err := c.do(ctx, http.MethodPost, "/instances/"+url.PathEscape(id)+"/terminate", nil, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *apiClient) getSSHInfo(ctx context.Context, id string) (*instanceSSHInfo, error) {
	var response instanceSSHInfo
	if err := c.do(ctx, http.MethodGet, "/instances/"+url.PathEscape(id)+"/ssh", nil, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *apiClient) getPool(ctx context.Context, id string) (*poolResponse, error) {
	var response poolResponse
	if err := c.do(ctx, http.MethodGet, "/pools/"+url.PathEscape(id), nil, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *apiClient) do(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	requestBody any,
	responseBody any,
) error {
	_, err := c.doRequest(ctx, method, brevAPIPath+path, query, requestBody, responseBody, nil)
	return err
}

func (c *apiClient) doRequest(
	ctx context.Context, method, path string, query url.Values,
	requestBody, responseBody any, headers http.Header,
) (http.Header, error) {
	var body io.Reader
	if requestBody != nil {
		encoded, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(encoded)
	}

	request, err := http.NewRequestWithContext(
		ctx,
		method,
		strings.TrimRight(c.baseURL, "/")+path,
		body,
	)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = query.Encode()
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.apiKey)
	for key, values := range headers {
		request.Header[key] = values
	}
	if requestBody != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return response.Header, &apiError{statusCode: response.StatusCode, body: string(responseBytes)}
	}
	if responseBody == nil || len(responseBytes) == 0 {
		return response.Header, nil
	}
	return response.Header, json.Unmarshal(responseBytes, responseBody)
}
