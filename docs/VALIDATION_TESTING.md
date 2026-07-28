# Validation Testing

This document describes the validation testing framework for cloud provider implementations.

## Overview

Validation tests verify that cloud provider implementations correctly implement the SDK interfaces by making real API calls to cloud providers. These tests are separate from unit tests and require actual cloud credentials.

## Running Validation Tests

### Locally

```bash
# Skip validation tests (default)
make test

# Run validation tests
make test-validation

# Run all tests
make test-all
```

### Environment Variables

Each provider's validation test documents the environment variables it requires.

### CI/CD

Validation tests run automatically:
- Daily via scheduled workflows
- On pull requests when labeled with `run-validation`
- Manually via workflow dispatch

## Adding New Providers

1. Create validation test file: `v1/providers/{provider}/validation_test.go`
2. Use the shared validation package with provider-specific configuration:

```go
func TestValidationFunctions(t *testing.T) {
    apiKey := os.Getenv("YOUR_PROVIDER_API_KEY")
    if apiKey == "" {
        t.Skip("YOUR_PROVIDER_API_KEY not set, skipping YourProvider validation tests")
    }

    config := validation.ProviderConfig{
        Credential: NewYourProviderCredential("validation-test", apiKey),
    }
    validation.RunValidationSuite(t, config)
}
```

3. Create CI workflow: `.github/workflows/validation-{provider}.yml`
4. Add environment variables to CI secrets
5. Update this documentation

## Shared Validation Package

The validation tests use a shared package at `internal/validation/` that provides:
- `RunValidationSuite()` - Tests all validation functions from pkg/v1/
- `RunInstanceLifecycleValidation()` - Tests instance lifecycle operations
- `ProviderConfig` - Configuration for provider-specific setup using CloudCredential

The `ProviderConfig` uses the existing `CloudCredential` interface which acts as a factory for `CloudClient` instances. The provider name is automatically obtained from the credential's `GetCloudProviderID()` method. This approach eliminates code duplication and ensures consistent validation testing across all providers while leveraging the existing credential abstraction.

## Test Structure

Validation tests use `testing.Short()` guards:

```go
func TestValidation(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping validation tests in short mode")
    }
    // ... validation logic
}
```

This ensures validation tests only run when explicitly requested.

## Validation Functions Tested

The framework tests all validation functions from the SDK:

### Instance Management
- `ValidateCreateInstance` - Tests instance creation with timing and attribute validation
- `ValidateListCreatedInstance` - Tests instance listing and filtering
- `ValidateTerminateInstance` - Tests instance termination
- `ValidateMergeInstanceForUpdate` - Tests instance update merging logic

### Instance Types
- `ValidateGetInstanceTypes` - Tests instance type retrieval and filtering
- `ValidateRegionalInstanceTypes` - Tests regional instance type filtering
- `ValidateStableInstanceTypeIDs` - Tests instance type ID stability

### Locations
- `ValidateGetLocations` - Tests location retrieval and availability

## Security Considerations

- Validation tests use real cloud credentials stored as GitHub secrets
- Tests create and destroy real cloud resources
- Proper cleanup is implemented to avoid resource leaks
- Tests are designed to be cost-effective and use minimal resources

## Troubleshooting

### Common Issues

1. **Missing credentials**: Ensure environment variables are set
2. **Quota limits**: Tests may skip if quota is exceeded
3. **Resource availability**: Tests adapt to available instance types and locations
4. **Network timeouts**: Tests use appropriate timeouts for cloud operations

### Debugging

```bash
# Run a specific provider's validation test
go test -v -short=false -run TestValidationFunctions ./v1/providers/{provider}/

# Run with verbose output
go test -v -short=false -timeout=20m ./v1/providers/{provider}/
```

## Contributing

When adding new validation functions:

1. Add the validation function to the appropriate `v1/*.go` file
2. Add corresponding test in `v1/providers/{provider}/validation_test.go`
3. Ensure proper cleanup and error handling
4. Update this documentation
