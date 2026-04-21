package v1

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	cloud "github.com/brevdev/cloud/v1"
)

const (
	CloudProviderID                = "naver"
	defaultBaseURL                 = "https://ncloud.apigw.ntruss.com"
	defaultRegionCode              = "KR"
	defaultServerImageProductCode  = "SW.VSVR.OS.LNX64.UBNTU.SVR24.G003"
	defaultSSHUser                 = "root"
	defaultSSHPort                 = 22
	defaultInstanceTypePollMinutes = 5
)

type NaverCredential struct {
	RefID     string
	AccessKey string
	SecretKey string
}

var _ cloud.CloudCredential = &NaverCredential{}

func NewNaverCredential(refID, accessKey, secretKey string) *NaverCredential {
	return &NaverCredential{
		RefID:     refID,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func (c *NaverCredential) GetReferenceID() string {
	return c.RefID
}

func (c *NaverCredential) GetAPIType() cloud.APIType {
	return cloud.APITypeGlobal
}

func (c *NaverCredential) GetCloudProviderID() cloud.CloudProviderID {
	return CloudProviderID
}

func (c *NaverCredential) GetTenantID() (string, error) {
	if c.AccessKey == "" {
		return "", fmt.Errorf("access key is required")
	}
	sum := sha256.Sum256([]byte(c.AccessKey))
	return fmt.Sprintf("%s-%x", CloudProviderID, sum), nil
}

func (c *NaverCredential) GetCapabilities(_ context.Context) (cloud.Capabilities, error) {
	return getNaverCapabilities(), nil
}

func (c *NaverCredential) MakeClient(_ context.Context, location string) (cloud.CloudClient, error) {
	return NewNaverClient(c.RefID, c.AccessKey, c.SecretKey, WithLocation(location))
}

type NaverClient struct {
	cloud.NotImplCloudClient
	refID      string
	accessKey  string
	secretKey  string
	baseURL    string
	httpClient *http.Client
	location   string
	now        func() time.Time
}

var _ cloud.CloudClient = &NaverClient{}

type options struct {
	baseURL    string
	httpClient *http.Client
	location   string
	now        func() time.Time
}

type Option func(*options)

func WithBaseURL(baseURL string) Option {
	return func(opts *options) {
		opts.baseURL = strings.TrimRight(baseURL, "/")
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(opts *options) {
		opts.httpClient = client
	}
}

func WithLocation(location string) Option {
	return func(opts *options) {
		opts.location = location
	}
}

func WithClock(now func() time.Time) Option {
	return func(opts *options) {
		opts.now = now
	}
}

func NewNaverClient(refID, accessKey, secretKey string, opts ...Option) (*NaverClient, error) {
	if refID == "" {
		return nil, fmt.Errorf("refID is required")
	}
	if accessKey == "" || secretKey == "" {
		return nil, fmt.Errorf("accessKey and secretKey are required")
	}

	options := options{
		baseURL:    defaultBaseURL,
		httpClient: http.DefaultClient,
		location:   defaultRegionCode,
		now:        time.Now,
	}
	for _, opt := range opts {
		opt(&options)
	}
	if options.location == "" {
		options.location = defaultRegionCode
	}
	if options.httpClient == nil {
		options.httpClient = http.DefaultClient
	}
	if options.now == nil {
		options.now = time.Now
	}

	return &NaverClient{
		refID:      refID,
		accessKey:  accessKey,
		secretKey:  secretKey,
		baseURL:    options.baseURL,
		httpClient: options.httpClient,
		location:   options.location,
		now:        options.now,
	}, nil
}

func (c *NaverClient) GetAPIType() cloud.APIType {
	return cloud.APITypeGlobal
}

func (c *NaverClient) GetCloudProviderID() cloud.CloudProviderID {
	return CloudProviderID
}

func (c *NaverClient) GetReferenceID() string {
	return c.refID
}

func (c *NaverClient) GetTenantID() (string, error) {
	sum := sha256.Sum256([]byte(c.accessKey))
	return fmt.Sprintf("%s-%x", CloudProviderID, sum), nil
}

func (c *NaverClient) MakeClient(_ context.Context, location string) (cloud.CloudClient, error) {
	if location != "" {
		c.location = location
	}
	return c, nil
}

func (c *NaverClient) do(ctx context.Context, action string, params url.Values, dst any) error {
	if params == nil {
		params = url.Values{}
	}
	params.Set("responseFormatType", "json")

	path := "/vserver/v2/" + action
	query := params.Encode()
	requestURI := path
	if query != "" {
		requestURI += "?" + query
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+requestURI, nil)
	if err != nil {
		return err
	}
	c.sign(req, requestURI)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("naver API %s failed with status %d: %s", action, resp.StatusCode, strings.TrimSpace(string(body)))
	}
	if err := json.Unmarshal(body, dst); err != nil {
		return fmt.Errorf("decode naver API %s response: %w", action, err)
	}
	if err := responseErr(dst); err != nil {
		return fmt.Errorf("naver API %s failed: %w", action, err)
	}
	return nil
}

func (c *NaverClient) sign(req *http.Request, requestURI string) {
	timestamp := fmt.Sprintf("%d", c.now().UnixMilli())
	message := fmt.Sprintf("%s\n%s\n%s\n%s", req.Method, requestURI, timestamp, c.accessKey)
	mac := hmac.New(sha256.New, []byte(c.secretKey))
	_, _ = mac.Write([]byte(message))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	req.Header.Set("x-ncp-apigw-timestamp", timestamp)
	req.Header.Set("x-ncp-iam-access-key", c.accessKey)
	req.Header.Set("x-ncp-apigw-signature-v2", signature)
}

type naverResponse interface {
	apiError() error
}

func responseErr(dst any) error {
	if resp, ok := dst.(naverResponse); ok {
		return resp.apiError()
	}
	return nil
}

type responseMeta struct {
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
}

func (m responseMeta) apiError() error {
	if m.ReturnCode != "" && m.ReturnCode != "0" {
		return fmt.Errorf("%s: %s", m.ReturnCode, m.ReturnMessage)
	}
	return nil
}

type codeName struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
