# Contributing to Nebius Brev Compute SDK

Nebius has a [golang SDK](https://github.com/nebius/gosdk) that is used to interact with the Nebius API.

Get started by reading the [Nebius API documentation](https://github.com/nebius/api).

## Local Development

### Prerequisites

1. **Nebius Account**: Create an account at [Nebius AI Cloud](https://nebius.com)
2. **Service Account**: Create a service account in Nebius IAM
3. **Service Account Key**: Generate and download a JSON service account key

### Setup

1. **Install Dependencies**:
   ```bash
   go mod download
   ```

2. **Configure Credentials**:
   Place your service account JSON key file in your home directory:
   ```bash
   cp /path/to/your/service-account-key.json ~/.nebius-credentials.json
   ```

3. **Set Environment Variables**:
   ```bash
   export NEBIUS_SERVICE_ACCOUNT_KEY_FILE=~/.nebius-credentials.json
   export NEBIUS_PROJECT_ID=your-project-id
   ```

### Running Tests

```bash
# Run all tests
make test

# Run Nebius-specific tests
go test ./internal/nebius/v1/...

# Run with verbose output
go test -v ./internal/nebius/v1/...
```

### Development Workflow

1. **Code Changes**: Make changes to the Nebius provider implementation
2. **Lint**: Run `make lint` to ensure code quality
3. **Test**: Run `make test` to verify functionality
4. **Commit**: Follow conventional commit messages

### Implementation Status

The current implementation provides boilerplate stubs for all CloudClient interface methods:

**Implemented (Stubs)**:
- Instance management (Create, Get, List, Terminate, Stop, Start, Reboot)
- Instance types and quotas
- Image management
- Location management
- Firewall/Security Group management
- Volume management
- Tag management

**Next Steps**:
- Replace stub implementations with actual Nebius API calls
- Add comprehensive error handling
- Implement proper resource mapping between Brev and Nebius models
- Add integration tests with real Nebius resources

### API Reference

- **Nebius Go SDK**: https://github.com/nebius/gosdk
- **Nebius API Documentation**: https://github.com/nebius/api
- **Compute Service**: Focus on `services/nebius/compute/v1/` for instance management
