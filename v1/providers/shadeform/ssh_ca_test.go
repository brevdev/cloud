package v1

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/shadeform/gen/shadeform"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
)

type bootstrapTransport func(*http.Request) (*http.Response, error)

func (f bootstrapTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func TestCreateInstanceSSHCertificateBootstrap(t *testing.T) { //nolint:funlen // test ok
	public, _, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	key, err := ssh.NewPublicKey(public)
	require.NoError(t, err)
	caKey := strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key)))
	for _, test := range []struct {
		name        string
		certificate bool
		existingKey string
		wantError   bool
	}{
		{name: "legacy"},
		{
			name:        "certificate",
			certificate: true,
		},
		{
			name:        "retry with matching CA",
			certificate: true,
			existingKey: caKey + " provider-comment",
		},
		{
			name:        "name collision cannot select a user key",
			certificate: true,
			existingKey: "legacy-user-key",
			wantError:   true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			args := v1.CreateInstanceAttrs{
				RefID:        "instance-ref",
				Name:         "workspace",
				InstanceType: "cloud_H100",
				PublicKey:    "legacy-user-key",
			}
			if test.certificate {
				args.SSHCertificateAuthority = &v1.SSHCertificateAuthority{
					PublicKey:       caKey,
					PrincipalPrefix: "brev:v1:vm:env-shadeform:login:",
				}
			}
			var registered openapi.AddSshKeyRequest
			var posted *openapi.CreateRequest
			config := openapi.NewConfiguration()
			config.HTTPClient = &http.Client{
				Transport: bootstrapTransport(func(request *http.Request) (*http.Response, error) {
					status := http.StatusOK
					var body any
					switch request.Method + " " + request.URL.Path {
					case "POST /v1/sshkeys/add":
						require.NoError(t, json.NewDecoder(request.Body).Decode(&registered))
						body = openapi.AddSshKeyResponse{Id: "registered-key"}
						if test.existingKey != "" {
							status = http.StatusConflict
							body = map[string]string{"error": "name must be unique"}
						}
					case "GET /v1/sshkeys":
						body = openapi.SshKeysResponse{SshKeys: []openapi.SshKey{
							{
								Id:        "default-user-key",
								Name:      "other-key",
								PublicKey: "legacy-user-key",
							},
							{
								Id:        "existing-ca-key",
								Name:      registered.Name,
								PublicKey: test.existingKey,
							},
						}}
					case "POST /v1/instances/create":
						require.NoError(t, json.NewDecoder(request.Body).Decode(&posted))
						body = openapi.CreateResponse{Id: "vm-id"}
					case "GET /v1/instances/vm-id/info":
						body = openapi.InstanceInfoResponse{
							Id:                "vm-id",
							Name:              posted.Name,
							Tags:              posted.Tags,
							Cloud:             "cloud",
							ShadeInstanceType: "H100",
							SshUser:           "shadeform",
							Status:            openapi.ACTIVE,
						}
					default:
						t.Fatalf("unexpected request: %s %s", request.Method, request.URL.Path)
					}
					encoded, err := json.Marshal(body)
					require.NoError(t, err)
					return &http.Response{
						StatusCode: status,
						Header:     http.Header{"Content-Type": {"application/json"}},
						Body:       io.NopCloser(strings.NewReader(string(encoded))),
						Request:    request,
					}, nil
				}),
			}
			client := NewShadeformClient("credential", "test-key")
			client.client = openapi.NewAPIClient(config)
			instance, err := client.CreateInstance(t.Context(), args)
			if test.wantError {
				require.Error(t, err)
				require.Nil(t, posted, "a conflicting key must never select the account default")
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.certificate, instance.SSHCertificateBootstrap)
			script, err := base64.StdEncoding.DecodeString(posted.LaunchConfiguration.ScriptConfiguration.Base64Script)
			require.NoError(t, err)
			require.Contains(t, string(script), "ufw allow 22/tcp")
			if !test.certificate {
				require.Equal(t, "legacy-user-key", registered.PublicKey)
				require.Equal(t, args.RefID, registered.Name)
				require.NotContains(t, string(script), "cert-authority")
				return
			}
			require.Equal(t, caKey, registered.PublicKey)
			require.Equal(t, "brev-sshca-"+args.RefID, registered.Name)
			id := "registered-key"
			if test.existingKey != "" {
				id = "existing-ca-key"
			}
			require.Equal(t, id, *posted.SshKeyId)
			require.Contains(t, string(script), "brev:v1:vm:env-shadeform:login:")
			require.NotContains(t, string(script), "legacy-user-key")
		})
	}
}
