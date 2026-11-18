package v1

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/brevdev/cloud/internal/collections"
	"github.com/brevdev/cloud/internal/sshbootstrap"
	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
)

type bootstrapTransport func(*http.Request) (*http.Response, error)

func (f bootstrapTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func TestCreateInstanceSSHCertificateBootstrap(t *testing.T) {
	public, _, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	key, err := ssh.NewPublicKey(public)
	require.NoError(t, err)
	caKey := strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key)))
	for _, test := range []struct {
		name        string
		publicKey   string
		certificate bool
		invalidCA   bool
	}{
		{name: "legacy", publicKey: "legacy-user-key"},
		{name: "certificate replaces user key", publicKey: "legacy-user-key", certificate: true},
		{name: "certificate needs no user key", certificate: true},
		{name: "invalid CA never reaches provider", publicKey: "legacy-user-key", certificate: true, invalidCA: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			args := v1.CreateInstanceAttrs{
				Name: "workspace", RefID: "env-launchpad", PublicKey: test.publicKey,
				InstanceType: "dmz.h100x1.pcie", Location: "pdx",
				FirewallRules: v1.FirewallRules{IngressRules: []v1.FirewallRule{{IPRanges: []string{"192.0.2.1/32"}}}},
			}
			if test.certificate {
				args.SSHCertificateAuthority = &v1.SSHCertificateAuthority{
					PublicKey: caKey + " provider-comment\n", PrincipalPrefix: "brev:v1:vm:env-launchpad:login:",
				}
				if test.invalidCA {
					args.SSHCertificateAuthority.PublicKey = "invalid-key"
				}
			}
			var posted map[string]json.RawMessage
			config := openapi.NewConfiguration()
			config.HTTPClient = &http.Client{Transport: bootstrapTransport(func(request *http.Request) (*http.Response, error) {
				if test.invalidCA {
					t.Fatal("invalid CA must be rejected before any provider request")
				}
				status := http.StatusOK
				switch request.Method + " " + request.URL.Path {
				case "POST /v1/catalog/deployments/":
					require.NoError(t, json.NewDecoder(request.Body).Decode(&posted))
					status = http.StatusCreated
				case "GET /v1/catalog/deployments/vm-id/":
				default:
					t.Fatalf("unexpected request: %s %s", request.Method, request.URL.Path)
				}
				response := openapi.Deployment{
					Id: "vm-id", SshUser: "launchpad-login", SshPort: 22, State: openapi.DeploymentStateWaiting,
					Cluster: &openapi.ClusterPipelineCluster{Cluster: &openapi.Cluster{
						Deployment: openapi.ClusterDeployment{String: collections.Ptr("vm-id")},
					}},
					Tags: map[string]string{"RefID": args.RefID},
				}
				encoded, err := json.Marshal(response)
				require.NoError(t, err)
				return &http.Response{
					StatusCode: status, Header: http.Header{"Content-Type": {"application/json"}},
					Body: io.NopCloser(strings.NewReader(string(encoded))), Request: request,
				}, nil
			})}
			client, err := NewLaunchpadClient(*NewLaunchpadCredential("credential", "test-key", "https://launchpad.invalid"))
			require.NoError(t, err)
			client.client = openapi.NewAPIClient(config)
			instance, err := client.CreateInstance(t.Context(), args)
			if test.invalidCA {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.certificate, instance.SSHCertificateBootstrap)
			require.JSONEq(t, `"env-launchpad"`, string(posted["request_id"]))
			require.JSONEq(t, `["192.0.2.1"]`, string(posted["ip_allowlist"]))
			var installedKey string
			require.NoError(t, json.Unmarshal(posted["public_key"], &installedKey))
			if test.certificate {
				require.Equal(t, caKey, installedKey)
				var script string
				require.NoError(t, json.Unmarshal(posted["startup_script"], &script))
				want, err := sshbootstrap.InstallForKeyOwners(*args.SSHCertificateAuthority)
				require.NoError(t, err)
				require.Equal(t, want, script, "send the installer as plain text without changing its quoting")
			} else {
				require.Equal(t, args.PublicKey, installedKey)
				require.NotContains(t, posted, "startup_script")
			}
			fetched, err := client.GetInstance(t.Context(), instance.CloudID)
			require.NoError(t, err)
			require.False(t, fetched.SSHCertificateBootstrap, "only CreateInstance reports bootstrap selection")
		})
	}
}
