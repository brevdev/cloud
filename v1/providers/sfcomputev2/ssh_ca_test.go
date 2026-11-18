package v2

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

func testPublicKey(t *testing.T) string {
	t.Helper()
	public, _, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	key, err := ssh.NewPublicKey(public)
	require.NoError(t, err)
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key)))
}

func TestCreateInstanceSSHCertificateBootstrap(t *testing.T) {
	userKey, caKey := testPublicKey(t), testPublicKey(t)
	for _, certificate := range []bool{false, true} {
		args := v1.CreateInstanceAttrs{
			RefID:     "instance-ref",
			PublicKey: userKey + " legacy-comment",
		}
		if certificate {
			args.SSHCertificateAuthority = &v1.SSHCertificateAuthority{
				PublicKey:       caKey,
				PrincipalPrefix: "brev:v1:vm:env-sfc:login:",
			}
		}
		var posted *createInstanceRequest
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			switch request.Method + " " + request.URL.Path {
			case "GET /integrations/brev/v1/pools/sfc:pool:account:workspace:default":
				writeJSON(t, writer, poolResponse{AllocationSchedule: allocationSchedule{
					ByInstanceSKU: map[string][]scheduleEntry{"sku": {{StartAt: 0, NodeCount: 1}}},
				}})
			case "GET /integrations/brev/v1/instances":
				writeJSON(t, writer, listInstancesResponse{})
			case "POST /integrations/brev/v1/instances":
				require.NoError(t, json.NewDecoder(request.Body).Decode(&posted))
				writeJSON(t, writer, instanceResponse{
					ID:     "vm-id",
					Status: instanceStatusAwaitingAllocation,
					Tags:   posted.Tags,
				})
			default:
				http.NotFound(writer, request)
			}
		}))
		t.Cleanup(server.Close)
		instance, err := newTestSFCClient(server.URL).CreateInstance(t.Context(), args)
		require.NoError(t, err)
		require.Equal(t, certificate, instance.SSHCertificateBootstrap)
		require.Equal(t, args.RefID, instance.RefID)
		userData, err := base64.StdEncoding.DecodeString(*posted.CloudInitUserData)
		require.NoError(t, err)
		var config struct {
			Keys []string `yaml:"ssh_authorized_keys"`
		}
		require.NoError(t, yaml.Unmarshal(userData, &config))
		require.Len(t, config.Keys, 1)
		if !certificate {
			require.Equal(t, []string{userKey}, config.Keys, "retain legacy key canonicalization")
			continue
		}
		key, _, options, _, err := ssh.ParseAuthorizedKey([]byte(config.Keys[0]))
		require.NoError(t, err)
		require.Equal(t, caKey, strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))))
		require.ElementsMatch(t, []string{"cert-authority", `principals="brev:v1:vm:env-sfc:login:ubuntu"`}, options)
		require.NotContains(t, string(userData), userKey)
	}
}
