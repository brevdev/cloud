# Testkube Ubuntu VM Image

This image backs the `test.ok.cpu` testkube instance type:

```text
ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest
```

## Publish to GHCR

Authenticate Docker to GHCR with the GitHub CLI:

```bash
gh auth status
gh auth refresh -h github.com -s write:packages
gh auth token | docker login ghcr.io -u "$(gh api user --jq .login)" --password-stdin
```

Build and push the image from the repository root. For EKS, publish an amd64 image because `test.ok.cpu` advertises `x86_64`:

```bash
docker buildx build \
  --platform linux/amd64 \
  -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest \
  --push \
  ./v1/providers/testkube/images/ubuntu-vm
```

If you need both local Apple Silicon clusters and amd64 EKS nodes to pull the same tag, publish a multi-arch manifest:

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest \
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
  -t ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest \
  ./v1/providers/testkube/images/ubuntu-vm
```
