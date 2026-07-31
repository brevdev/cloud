# Testkube Ubuntu VM Image

This image backs the architecture-specific TestKube instance types:

```text
test.ok.cpu       -> ghcr.io/brevdev/cloud/testkube-ubuntu-vm:multiarch-v2
test.ok.cpu.arm64 -> ghcr.io/brevdev/cloud/testkube-ubuntu-vm:multiarch-v2
```

## Publish to GHCR

Authenticate Docker to GHCR with the GitHub CLI:

```bash
gh auth status
gh auth refresh -h github.com -s write:packages
gh auth token | docker login ghcr.io -u "$(gh api user --jq .login)" --password-stdin
```

Build and push the image from the repository root. For EKS, publish the versioned multi-arch manifest used by both `test.ok.cpu` and `test.ok.cpu.arm64`:

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:multiarch-v2 \
  --push \
  ./v1/providers/testkube/images/ubuntu-vm
```

You can also use an explicit token instead of `gh auth token`:

```bash
echo "$GITHUB_TOKEN" | docker login ghcr.io -u "$GITHUB_USER" --password-stdin
```

The token needs `write:packages` to publish and `read:packages` for clusters pulling a private GHCR package.

## Local Build

For local minikube or kind validation where the image is loaded directly into the cluster, a normal local build is enough:

```bash
docker build \
  -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:multiarch-v2 \
  ./v1/providers/testkube/images/ubuntu-vm
```
