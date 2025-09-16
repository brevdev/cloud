package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestLaunchpadClient_CreateInstance(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewLaunchpadCredential("brev-cloud-sdk-test", apiKey, "https://stage.launchpad.api.nvidia.com"),
	}

	client, err := config.Credential.MakeClient(context.Background(), config.Location)
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	instance, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		Name:         "cloud-sdk-test",
		RefID:        "cloud-sdk-test",
		PublicKey:    "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCW3Xmg9/Xu/hY7CMDR1twMD9G0l/amVbOm6mSTTYjxxv7nGlIIiaL8G/XOuHg6KUBf+7NQxZYRLiGqKAlXeC3YlzvQFsroyYd8SvBLHp60aHsvHh/KEFRTFtWlSaE3Aa9Qhr14YXL51tN3+sfTBVU6oqn1t6vX9SUS3FQPCMuKDCVyBQsvz27bI4AZisYyn7mVs3Xr22h/4Pmfa/+KsNdG6nEhaZGqkM+nCQd6lfjIxe234SDBsA/4xCw/F0BBkK3mc2TsKz8sQHsrVnBDotsRFMdHmkj0msFsslh3/teSVPvVpod90A8MZEfOhMbMtPSNySdwsoLAkjOTnhUROvx/kKu4eIW5cz6oHODQ+GRNxAwzS9RXb7ES5Mex09tbgtP7yC1sgQ9GxFRHKD5ZDfrfKRj5gc098x0eEX16wAnGu8j4vsidzn4jPOJwf6UtNYqvecEVuDpWq5QwzZRpcanAjVSD6UZEJjEZDwtol8SbtwkgaaWj59EeMg5G6M6J/5hry04Me7wO4QwP/V1l25uhAg0Vt/0l+oz9Z5qKodllVRCEc55ile2qWAJHLv1B3YHFQFts4DesZcxoCyBZX+KkRYzJoZ4muOSCNYXTAplaqDfYxRsz/YlK30U5lEwDOifXGNKqLdzrZwhNTRWUJR36DHBI+ok8zs3yPHicuZQ5Ww== dmalin@nvidia.com",
		InstanceType: "dmz.h100x1.pcie",
		Location:     "pdx",
	})
	require.NoError(t, err)

	t.Logf("instance: %v", instance)
}
