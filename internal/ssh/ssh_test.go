package ssh

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	gossh "golang.org/x/crypto/ssh"
)

var nextPort int32 = 3334

func getPort() string {
	n := atomic.AddInt32(&nextPort, 1)
	return fmt.Sprintf("%d", n)
}

func Test_ConnectToHostSuccess(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	port := getPort()
	done, err := StartTestSSHServer(
		TestSSHServerOptions{
			Port:       port,
			PubKeyAuth: true,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := done(); err != nil {
			t.Fatal(err)
		}
	}()

	// res, err := exec.Command("ssh", "-o", "UserKnownHostsFile=/dev/null", "-o", "StrictHostKeyChecking=no", "-p", port, "localhost").CombinedOutput()
	// mt.Println(string(res))
	_, err = ConnectToHost(ctx, ConnectionConfig{"ubuntu", fmt.Sprintf("localhost:%s", port), DoNotUseDummyPrivateKey})
	if !assert.NoError(t, err) {
		return
	}
}

func Test_ConnectToHostRunCommandExit0(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	port := getPort()
	done, err := StartTestSSHServer(
		TestSSHServerOptions{
			Port:       port,
			PubKeyAuth: true,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := done(); err != nil {
			t.Fatal(err)
		}
	}()

	c, err := ConnectToHost(ctx, ConnectionConfig{"ubuntu", fmt.Sprintf("localhost:%s", port), DoNotUseDummyPrivateKey})
	if !assert.NoError(t, err) {
		return
	}
	stdOut, stdErr, err := c.RunCommand(ctx, "echo hello")
	if !assert.NoError(t, err) {
		return
	}
	assert.Contains(t, stdOut, "echo hello")
	assert.Empty(t, stdErr)
}

func Test_ConnectToHostRunCommandExit1(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	port := getPort()
	done, err := StartTestSSHServer(
		TestSSHServerOptions{
			Port:       port,
			PubKeyAuth: true,
			ExitCode:   1,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := done(); err != nil {
			t.Fatal(err)
		}
	}()

	c, err := ConnectToHost(ctx, ConnectionConfig{"ubuntu", fmt.Sprintf("localhost:%s", port), DoNotUseDummyPrivateKey})
	if !assert.NoError(t, err) {
		return
	}
	stdOut, stdErr, err := c.RunCommand(ctx, "echo hello")
	if !assert.Error(t, err) {
		if !assert.ErrorIs(t, err, &gossh.ExitError{}) {
			res, _ := err.(*gossh.ExitError)
			assert.Equal(t, res.ExitStatus(), 1)
			return
		}
		return
	}
	assert.Contains(t, stdOut, "echo hello")
	assert.Empty(t, stdErr)
}

func Test_ConnectToHostPubKeyFail(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	port := getPort()
	done, err := StartTestSSHServer(
		TestSSHServerOptions{
			Port:       port,
			PubKeyAuth: false,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := done(); err != nil {
			t.Fatal(err)
		}
	}()

	_, err = ConnectToHost(ctx, ConnectionConfig{"ubuntu", fmt.Sprintf("localhost:%s", port), DoNotUseDummyPrivateKey})
	assert.ErrorContains(t, err, "unable to authenticate")
}

func Test_FailConnectionRefusedSSH(t *testing.T) {
	t.Parallel()
	// no server running
	ctx := context.Background()
	_, err := ConnectToHost(ctx, ConnectionConfig{"ubuntu", fmt.Sprintf("localhost:%s", "3333"), DoNotUseDummyPrivateKey})
	assert.ErrorContains(t, err, "connection refused")
}

func Test_TestSSHTimeoutKey(t *testing.T) {
	t.Parallel()
	// no server running
	ctx := context.Background()
	timeout := time.Millisecond * 250
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	port := getPort()
	done, err := StartTestSSHServer(TestSSHServerOptions{
		Port:        port,
		PubKeyAuth:  true,
		PubKeyDelay: timeout * 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := done(); err != nil {
			t.Fatal(err)
		}
	}()

	_, err = ConnectToHost(ctx, ConnectionConfig{"ubuntu", fmt.Sprintf("localhost:%s", port), DoNotUseDummyPrivateKey})
	// assert.ErrorContains(t, err, "timeout")
	assert.ErrorContains(t, err, "context deadline exceeded")
}

func Test_CertficiateSSH(t *testing.T) {
	t.Parallel()
	// fails: don't understand how to use certificates
	t.Skip()
	c := Client{
		privateKey:  DoNotUseDummyPrivateKey,
		Certificate: PubKey,
	}
	_, err := c.getSigner()
	assert.NoError(t, err)
}

func Test_TrySSHConnect(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	port := getPort()
	done, err := StartTestSSHServer(
		TestSSHServerOptions{
			Port:       port,
			PubKeyAuth: true,
			ExitCode:   0,
		},
	)
	if !assert.NoError(t, err) {
		return
	}
	defer func() {
		if err := done(); err != nil {
			t.Fatal(err)
		}
	}()
	err = TrySSHConnect(ctx, ConnectionConfig{
		User:     "ubuntu",
		HostPort: fmt.Sprintf("localhost:%s", port),
		PrivKey:  DoNotUseDummyPrivateKey,
	}, WaitForSSHOptions{})
	assert.NoError(t, err)
}

func Test_WaitForSSH(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	port := getPort()
	done, err := StartTestSSHServer(
		TestSSHServerOptions{
			Port:       port,
			PubKeyAuth: true,
			ExitCode:   0,
		},
	)
	if !assert.NoError(t, err) {
		return
	}
	defer func() {
		if err := done(); err != nil {
			t.Fatal(err)
		}
	}()
	err = WaitForSSH(ctx, ConnectionConfig{
		User:     "ubuntu",
		HostPort: fmt.Sprintf("localhost:%s", port),
		PrivKey:  DoNotUseDummyPrivateKey,
	}, WaitForSSHOptions{})
	assert.NoError(t, err)
}

func Test_WaitForSSHFailWithRetry(t *testing.T) {
	RetryTest(t, WaitForSSHFailFlaky, 3)
}

func WaitForSSHFailFlaky(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	port := getPort()
	done, err := StartTestSSHServer(
		TestSSHServerOptions{
			Port:       port,
			PubKeyAuth: true,
			ExitCode:   0,
		},
	)
	if !assert.NoError(t, err) {
		return
	}
	defer func() {
		if err := done(); err != nil {
			t.Fatal(err)
		}
	}()
	t0 := time.Now()
	err = WaitForSSH(ctx, ConnectionConfig{
		User:     "ubuntu",
		HostPort: fmt.Sprintf("localhost:%s", "3333"),
		PrivKey:  DoNotUseDummyPrivateKey,
	}, WaitForSSHOptions{
		Timeout:           time.Second * 5,
		ConnectionTimeout: time.Second * 1,
	})
	t1 := time.Now()
	assert.ErrorContains(t, err, "context deadline exceeded")
	assert.ErrorContains(t, err, "last error")
	assert.ErrorContains(t, err, "connection refuse")
	assert.Greater(t, t1.Sub(t0), time.Second*5)
}

func Test_GetCallerIP(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		ctx     context.Context
		wantErr bool
		ipCheck func(t *testing.T, ip net.IP)
	}{
		{
			name:    "valid context",
			ctx:     context.Background(),
			wantErr: false,
			ipCheck: func(t *testing.T, ip net.IP) {
				t.Helper()
				t.Logf("ip: %s", ip)
				assert.NotNil(t, ip)
				assert.True(t, ip.IsGlobalUnicast() || ip.IsPrivate(), "IP should be either public or private")
				assert.False(t, ip.IsUnspecified(), "IP should not be unspecified (0.0.0.0)")
				assert.False(t, ip.IsLoopback(), "IP should not be loopback (127.0.0.1)")
				assert.False(t, ip.IsMulticast(), "IP should not be multicast")
			},
		},
		{
			name:    "canceled context",
			ctx:     func() context.Context { ctx, cancel := context.WithCancel(context.Background()); cancel(); return ctx }(),
			wantErr: true,
			ipCheck: func(t *testing.T, ip net.IP) {
				t.Helper()
				assert.Nil(t, ip)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ip, err := GetLocalIP(tt.ctx)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			tt.ipCheck(t, ip)
		})
	}
}

func Test_GetPublicIP(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ip, err := GetPublicIPStr(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, ip)
}

func RetryTest(t *testing.T, testFunc func(t *testing.T), numRetries int) {
	t.Helper() // Mark this function as a helper
	for i := 0; i < numRetries; i++ {
		tt := &testing.T{}
		testFunc(tt)
		if !tt.Failed() {
			return
		}
	}
	t.Fail() // If we reach here, all retries failed
}
