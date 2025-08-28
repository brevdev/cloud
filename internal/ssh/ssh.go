package ssh

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
)

func init() {
	setPublicIP(context.Background())
	setLocalIP(context.Background())
}

func ConnectToHost(ctx context.Context, config ConnectionConfig) (*Client, error) {
	d := net.Dialer{}
	sshClient := &Client{
		addr:       config.HostPort,
		user:       config.User,
		dial:       d.DialContext,
		privateKey: config.PrivKey,
	}
	fmt.Printf("local_ip: %s, public_ip: %s\n", localIP, publicIP)
	err := sshClient.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	return sshClient, nil
}

// modified from https://github.com/superfly/flyctl/blob/master/ssh/client.go
// https://github.com/golang/go/issues/20288#issuecomment-832033017
type Client struct {
	addr string
	user string

	dial func(ctx context.Context, network, addr string) (net.Conn, error)

	privateKey, Certificate string

	hostKeyAlgorithms []string

	client *gossh.Client
	conn   gossh.Conn
}

func sshBackoff() backoff.BackOff {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 6 * time.Second
	b.InitialInterval = 250 * time.Millisecond
	b.MaxInterval = 3 * time.Second
	return b
}

func (c *Client) RunCommand(ctx context.Context, cmd string) (string, string, error) {
	var stdout, stderr string
	var err error

	err = backoff.Retry(func() error {
		stdout, stderr, err = c.runCommand(ctx, cmd)
		if err != nil && (err == io.EOF || strings.Contains(err.Error(), "unexpected packet in response to channel open: <nil>")) {
			cerr := c.Connect(ctx)
			if cerr != nil {
				return fmt.Errorf("connection error: %w, original error: %w", cerr, err)
			}
			return err
		} else if err != nil {
			return backoff.Permanent(err)
		}
		return nil
	}, backoff.WithContext(sshBackoff(), ctx))

	return stdout, stderr, err
}

// returns stdout, stderr and error that may be an ssh ExitError
func (c *Client) runCommand(ctx context.Context, cmd string) (string, string, error) {
	if c.client == nil {
		if err := c.Connect(ctx); err != nil {
			return "", "", fmt.Errorf("failed to connect: %w", err)
		}
	}

	sess, err := c.client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session, try a new connection: %w", err)
	}
	defer func() {
		if err := sess.Close(); err != nil {
			fmt.Printf("failed to close session: %v\n", err)
		}
	}()

	var stdOutBuffer bytes.Buffer
	sess.Stdout = &stdOutBuffer
	var stdErrBuffer bytes.Buffer
	sess.Stderr = &stdErrBuffer

	err = sess.Run(cmd)
	return stdOutBuffer.String(), stdErrBuffer.String(), err
}

func (c *Client) Close() error {
	if c == nil {
		return nil
	}
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return fmt.Errorf("failed to close connection: %w", err)
		}
	}

	c.conn = nil
	return nil
}

func (c Client) getSigner() (gossh.Signer, error) {
	signer, err := gossh.ParsePrivateKey([]byte(c.privateKey))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	if c.Certificate != "" {
		pubKey, _, _, _, err := gossh.ParseAuthorizedKey([]byte(c.Certificate))
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate: %w", err)
		}

		cert, ok := pubKey.(*gossh.Certificate)
		if !ok {
			return nil, fmt.Errorf("SSH public key must be a certificate")
		}
		signer, err = gossh.NewCertSigner(cert, signer)
		if err != nil {
			return nil, fmt.Errorf("failed to create cert signer: %w", err)
		}
	}

	return signer, nil
}

type connResp struct {
	err    error
	conn   gossh.Conn
	client *gossh.Client
}

func (c *Client) Connect(ctx context.Context) error {
	signer, err := c.getSigner()
	if err != nil {
		return fmt.Errorf("failed to get signer: %w", err)
	}

	tcpConn, err := c.dial(ctx, "tcp", c.addr)
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}

	conf := &gossh.ClientConfig{
		User: c.user,
		Auth: []gossh.AuthMethod{
			gossh.PublicKeys(signer),
		},
		HostKeyCallback:   gossh.InsecureIgnoreHostKey(), //nolint:gosec // audited
		HostKeyAlgorithms: c.hostKeyAlgorithms,
	}

	respCh := make(chan connResp)

	// ssh.NewClientConn doesn't take a context, so we need to handle cancelation on our end
	go func() {
		conn, chans, reqs, errr := gossh.NewClientConn(tcpConn, tcpConn.RemoteAddr().String(), conf)
		if errr != nil {
			respCh <- connResp{err: errr}
			return
		}

		client := gossh.NewClient(conn, chans, reqs)

		respCh <- connResp{nil, conn, client}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case resp := <-respCh:
			if resp.err != nil {
				return resp.err
			}
			c.conn = resp.conn
			c.client = resp.client
			return nil
		}
	}
}

//nolint:gosec // WARNING: do not use these keys for anything other than testing
const DoNotUseDummyPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCEvcaEC2HvVDV277n6n23KXPwHoWX5mEkuoezqurwSJgq5grQz
Ka3pwdTmRd1CPM9UAXV7aK7UmpMyjSmukmna6CyLXCv61BDrodFb488p4MaPUnwG
FhilkjgcQLBWdHRKcUJZoszdY0kWVWbeUXzSrmTLzuGMmaN32dAXop31CQIDAQAB
AoGACAK33zIcp+fKDjJrY8+JPaQc5Yz87XIeQH0vIf9A6Et5bDaSD2BdiXTUF01y
C9RFoskvwNHRcy0c4vkX4dweHSvHboFAU0ygKU5Dfou1JlmJeK6J+2xrEVGLIyKP
aMWVpyqmDCAKUzO0jEzzDJCZ95KDw9OWS7SBxC9bsRS2soECQQDwyiO6dBWYvRE0
8GDte+c9MbIdnzYuEeCXyGsK1prAaGLNHoOylp9yXj7M8UKyU+1LOZexjAXvv3tP
imEjHteRAkEAjSBdGf6LAPpGGwK3TuSi2GlJsLWW2trWBuY6+LY9nPnMDCTzYkxt
lD4lkCOxhcB6bNstbL9nBjoo3vHciC85+QJBAOBlUOSLGDFOSUHPnlTTOk1yCa7H
WAOZD3gEA5WHJ5KV9TV48Xy2GAPKRrZRRDnSMvr+whppBoNGLFGVAS9sp7ECQGvj
AumdWzyrF68Me4A3b3qLuwb5O1MiGp55oTmDcESx/liGYv2Rue+rNuIjN1It3Cmd
wPMyu5raGWaedV4y5FkCQQChhO3jMmLXQwCwLVCCfd9duiC1swwpvXm94Byk2h81
l2FHbn+D8BPAoE/vO/eLAOQVDAgLu0evktWWdtBckUoZ
-----END RSA PRIVATE KEY-----`

const PubKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCEvcaEC2HvVDV277n6n23KXPwH
oWX5mEkuoezqurwSJgq5grQzKa3pwdTmRd1CPM9UAXV7aK7UmpMyjSmukmna6CyL
XCv61BDrodFb488p4MaPUnwGFhilkjgcQLBWdHRKcUJZoszdY0kWVWbeUXzSrmTL
zuGMmaN32dAXop31CQIDAQAB
-----END PUBLIC KEY-----`

type TestSSHServerOptions struct {
	Port        string
	PubKeyAuth  bool
	PubKeyDelay time.Duration
	ExitCode    int
}

func StartTestSSHServer(options TestSSHServerOptions) (func() error, error) {
	handler := func(s ssh.Session) {
		authorizedKey := gossh.MarshalAuthorizedKey(s.PublicKey())
		_, err := io.WriteString(s, fmt.Sprintf("public key used by %s:\n", s.User())) // writes to client output
		if err != nil {
			fmt.Println(err)
		}
		_, err = s.Write(authorizedKey)
		if err != nil {
			fmt.Println(err)
		}
		_, err = s.Write([]byte(s.RawCommand()))
		if err != nil {
			fmt.Println(err)
		}
		err = s.Exit(options.ExitCode)
		if err != nil {
			fmt.Println(err)
		}
	}

	publicKeyOption := ssh.PublicKeyAuth(func(_ ssh.Context, _ ssh.PublicKey) bool {
		time.Sleep(options.PubKeyDelay)
		return options.PubKeyAuth // allow all keys, or use ssh.KeysEqual() to compare against known keys
	})

	server := ssh.Server{
		Addr: fmt.Sprintf(":%s", options.Port),
	}
	server.Handler = handler
	err := server.SetOption(publicKeyOption)
	if err != nil {
		return nil, fmt.Errorf("failed to set public key option: %w", err)
	}

	go func() {
		err1 := server.ListenAndServe()
		if err1 != nil {
			fmt.Println(err1)
		}
	}()
	time.Sleep(100 * time.Millisecond)
	return server.Close, nil
}

type ConnectionConfig struct {
	User, HostPort, PrivKey string
}

type WaitForSSHOptions struct {
	Timeout           time.Duration
	ConnectionTimeout time.Duration
	CheckCmd          string
	WaitTime          time.Duration
}

func (o *WaitForSSHOptions) SetDefault() {
	if o.Timeout == 0 {
		o.Timeout = 60 * time.Second
	}
	if o.ConnectionTimeout == 0 {
		o.ConnectionTimeout = 30 * time.Second
	}
	if o.CheckCmd == "" {
		o.CheckCmd = "echo 'connected'" // WARNING: assumes echo command exists
	}
	if o.WaitTime == 0 {
		o.WaitTime = 1 * time.Second
	}
}

func WaitForSSH(ctx context.Context, c ConnectionConfig, options WaitForSSHOptions) error {
	options.SetDefault()
	errChan := make(chan error, 1)
	errChan <- nil
	t0 := time.Now()

	// Create a timeout context
	timeoutCtx, cancel := context.WithTimeout(ctx, options.Timeout)
	defer cancel()

	err := doWithTimeout(timeoutCtx, func(ctx context.Context) error {
		waitForSSH(ctx, errChan, c, options)
		return nil
	})

	lastSSHErr := <-errChan
	t1 := time.Now()
	if err != nil {
		err = fmt.Errorf("took %s: %w", t1.Sub(t0).String(), err)
		if lastSSHErr != nil {
			lastSSHErr = fmt.Errorf("last error: %w", lastSSHErr)
			return fmt.Errorf("%w, %w", err, lastSSHErr)
		}
		return err
	}
	return nil
}

func doWithTimeout(ctx context.Context, fn func(context.Context) error) error {
	done := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				done <- fmt.Errorf("panic recovered: %v", r)
			}
		}()
		done <- fn(ctx)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func waitForSSH(ctx context.Context, errChan chan error, c ConnectionConfig, options WaitForSSHOptions) {
	for ctx.Err() == nil {
		_ = <-errChan
		tryCtx, cancel := context.WithTimeout(ctx, options.ConnectionTimeout)
		sshErr := TrySSHConnect(tryCtx, c, options)
		cancel()
		errChan <- sshErr
		if sshErr == nil {
			return
		}
		time.Sleep(options.WaitTime)
	}
}

func TrySSHConnect(ctx context.Context, c ConnectionConfig, options WaitForSSHOptions) error {
	con, err := ConnectToHost(ctx, c)
	if err != nil {
		return fmt.Errorf("failed to connect to host: %w", err)
	}
	defer func() {
		if closeErr := con.Close(); closeErr != nil {
			// Log close error but don't return it as it's not the primary error
			fmt.Printf("warning: failed to close SSH connection: %v\n", closeErr)
		}
	}()
	_, _, err = con.RunCommand(ctx, options.CheckCmd)
	if err != nil {
		return fmt.Errorf("failed to run check command: %w", err)
	}
	return nil
}

var (
	localIP     string
	localIPOnce sync.Once
)

func setLocalIP(ctx context.Context) {
	localIPOnce.Do(func() {
		ip, err := GetLocalIP(ctx)
		if err != nil {
			fmt.Printf("failed to get local IP: %v\n", err)
			return
		}
		localIP = ip.String()
	})
}

func GetLocalIP(ctx context.Context) (net.IP, error) {
	dialer := net.Dialer{}
	conn, err := dialer.DialContext(ctx, "udp", "8.8.8.8:80")
	if err != nil {
		return nil, fmt.Errorf("failed to dial for local IP: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("failed to close connection: %v\n", err)
		}
	}()

	localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return nil, fmt.Errorf("error getting local IP")
	}
	return localAddr.IP, nil
}

var (
	publicIP     string
	publicIPOnce sync.Once
)

// setPublicIP retrieves and sets the public IP address once.
// It uses sync.Once to ensure the IP is only fetched one time.
func setPublicIP(ctx context.Context) {
	publicIPOnce.Do(func() {
		ip, err := GetPublicIPStr(ctx)
		if err != nil {
			fmt.Printf("failed to get public IP: %v\n", err)
			return
		}
		publicIP = ip
	})
}

func GetPublicIPStr(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.ipify.org", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get public IP: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(ip), nil
}

func GetTestPrivateKey() string {
	privateKey, err := base64.StdEncoding.DecodeString(os.Getenv("TEST_PRIVATE_KEY_BASE64"))
	if err != nil {
		panic(err)
	}
	return string(privateKey)
}

func GetTestPublicKey() string {
	pubKey, err := base64.StdEncoding.DecodeString(os.Getenv("TEST_PUBLIC_KEY_BASE64"))
	if err != nil {
		panic(err)
	}
	return string(pubKey)
}
