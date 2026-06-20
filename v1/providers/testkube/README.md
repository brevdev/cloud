# Test Kubernetes Provider

`test-kubernetes` is a developer-only provider that backs cloud instance lifecycle calls with Kubernetes resources.

## Credentials

The credential supports either a base64-encoded kubeconfig or in-cluster Kubernetes authentication:

```go
type TestKubeCredential struct {
	RefID            string
	AuthMode         TestKubeAuthMode // "kubeconfig" or "in-cluster"
	KubeconfigBase64 string
	Namespace        string
}
```

When `AuthMode` is empty, it defaults to `"kubeconfig"` for compatibility. When `AuthMode` is `"in-cluster"`, the provider uses `rest.InClusterConfig()` and requires `KubeconfigBase64` to be empty. This is intended for dev-plane running inside the same Kubernetes cluster it will use as the testkube target. The pod's Kubernetes service account must have RBAC permissions to manage the target namespace's testkube resources.

## Validation

The validation tests are opt-in and skipped unless `TESTKUBE_KUBECONFIG_BASE64` is set. The inventory and failure tests only need Kubernetes API access. The real lifecycle validation needs a runnable image tagged as `ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest` and a working Kubernetes `LoadBalancer` implementation.

`test.ok.cpu` uses a `LoadBalancer` Service for SSH. The instance remains `pending` until the pod is ready and the Service has a load balancer ingress address. This more closely emulates real providers because arbitrary machines can use the returned `PublicIP`/`PublicDNS` and `SSHPort` without sharing the provider process.

### Local: minikube

Local validation should use minikube with `minikube tunnel`. The tunnel updates normal Kubernetes `LoadBalancer` Service status and makes the reported external IP reachable from the host, so the provider does not need local-cluster-specific endpoint translation.

```bash
brew install minikube kubectl

minikube start --driver=docker --profile testkube
kubectl config use-context testkube

docker build -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest ./v1/providers/testkube/images/ubuntu-vm
minikube --profile testkube image load ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest
kubectl create namespace testkube

# In another terminal, keep this running while validation runs.
sudo minikube --profile testkube tunnel

# Polulate .env with the contents of:
kubectl config view --raw --minify | base64 | tr -d '\n'

# .env
TESTKUBE_KUBECONFIG_BASE64=<above base64>
TESTKUBE_NAMESPACE=testkube
```

For in-cluster validation, set `TESTKUBE_AUTH_MODE=in-cluster` and omit `TESTKUBE_KUBECONFIG_BASE64`.

Clean up:

```bash
minikube --profile testkube delete
```

### CI: kind

kind remains the preferred CI path because it runs a disposable Kubernetes cluster on the GitHub Linux runner's Docker runtime. The CI path still uses `cloud-provider-kind` to provide `LoadBalancer` Services, but the provider itself only reads standard Kubernetes Service status.

Minimal GitHub Actions sketch:

```yaml
jobs:
  testkube-validation:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
      - name: Install cloud-provider-kind
        run: go install sigs.k8s.io/cloud-provider-kind@latest
      - name: Create kind cluster
        run: |
          kind create cluster --name testkube --wait 5m
          kubectl label node testkube-control-plane node.kubernetes.io/exclude-from-external-load-balancers- || true
          docker build -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest ./v1/providers/testkube/images/ubuntu-vm
          kind load docker-image ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest --name testkube
          kubectl create namespace testkube
          echo "TESTKUBE_KUBECONFIG_BASE64=$(kind get kubeconfig --name testkube | base64 | tr -d '\n')" >> "$GITHUB_ENV"
          echo "TESTKUBE_NAMESPACE=testkube" >> "$GITHUB_ENV"
          echo "VALIDATION_TEST=1" >> "$GITHUB_ENV"
      - name: Start cloud-provider-kind
        run: |
          sudo "$(go env GOPATH)/bin/cloud-provider-kind" > /tmp/cloud-provider-kind.log 2>&1 &
      - name: Run provider validation
        run: go test -v -run 'TestValidationFunctions|TestInstanceLifecycleValidation|TestFailureInstanceTypesValidation' ./v1/providers/testkube
```

These tests validate inventory, Kubernetes resource creation/listing, SSH access, stop/start/delete, and configured failure types. They do not validate dev-plane software setup.

### Image-Backed SSH Validation

Build the local image:

```bash
docker build -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest ./v1/providers/testkube/images/ubuntu-vm
```

For local minikube validation, load the image into the minikube profile:

```bash
minikube --profile testkube image load ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest
```

For CI kind validation, load the image into the kind cluster:

```bash
kind load docker-image ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest --name testkube
```

Then run the focused SSH validation:

```bash
go test -v -run TestImageBackedInstanceValidation ./v1/providers/testkube
```

`TestImageBackedInstanceValidation` creates a `test.ok.cpu` instance with the baked image tag, waits for the pod and load balancer to become ready, uses the provider-returned SSH endpoint, and verifies key-based SSH plus `sudo`, `apt-get`, and systemd basics.

CI can add the same build/load step before running the image-backed test:

```yaml
      - name: Build local testkube image
        run: docker build -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest ./v1/providers/testkube/images/ubuntu-vm
      - name: Load local testkube image into kind
        run: kind load docker-image ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest --name testkube
      - name: Run image-backed validation
        run: go test -v -run TestImageBackedInstanceValidation ./v1/providers/testkube
```

## Image Contract

`test.ok.cpu` points at `ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest`. Before that image is published to GHCR, build the local image with that same tag. Once the image is published, local and CI validation can either pull it or keep building and loading the same tag for hermetic tests.

The image at `images/ubuntu-vm` is expected to behave like a minimal Ubuntu VM for dev-plane: SSH access, `sudo`, `apt-get`, and systemd compatibility.

The cloud provider does not lay down dev-plane software; dev-plane owns that. The image only needs to provide a base OS environment that dev-plane setup can use.

Future exposure modes, such as fixed `NodePort` for one-off local debugging, should be separate instance type specs rather than credential fields.
