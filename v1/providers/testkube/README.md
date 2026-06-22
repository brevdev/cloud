# Test Kubernetes Provider

`test-kubernetes` is a test/non-prod-only provider that backs cloud instance lifecycle calls with Kubernetes resources.

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

### Mode: In-Cluster

The `testkube` provider can be used as any other, with the caveat being that "VM" resources are actually represented by pods within the context k8s cluster. This allows "environments" to be spun up and down quickly and cheaply, even though they don't necessarily perfectly emulate cloud-provided VMs.

### Mode: Kubeconfig

Alternatively to the in-cluster mode, resources can also be hosted by an arbitrary kubernetes cluster. This cluster can be hosted (e.g.: another EKS cluster) but can also be running locally. For example, local validation can use minikube with `minikube tunnel`. The tunnel updates normal Kubernetes `LoadBalancer` Service status and makes the reported external IP reachable from the host, so the provider does not need local-cluster-specific endpoint translation.

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