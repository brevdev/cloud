# Nebius Cloud SDK Integration - Testing & Development Guide

## Overview

This guide provides comprehensive instructions for testing and developing the Nebius cloud provider integration within the Brev Cloud SDK. The implementation has been revised based on analysis of the official Nebius Go SDK and existing provider patterns.

## Current Implementation Status

### ✅ Completed
- **Authentication Framework**: ✅ **WORKING** - Uses proper Nebius service account JSON format with real SDK authentication
- **Project-Per-User Model**: ✅ **WORKING** - Groups each Brev user's instances into dedicated Nebius projects
- **Client Structure**: ✅ **WORKING** - Follows Cloud SDK patterns with tenant → project → resources hierarchy
- **Interface Compliance**: ✅ **WORKING** - All required CloudClient methods implemented
- **Error Handling**: ✅ **WORKING** - Proper error wrapping and context handling
- **Build System**: ✅ **WORKING** - Compiles and tests pass with Go 1.24+

### 🚧 In Progress (Mock Implementation)
- **Instance Management**: Methods return **mock data** instead of creating real Nebius VMs
  - `CreateInstance()`: Returns mock instance (no real VM created)
  - `GetInstance()`: Returns mock instance data
  - `TerminateInstance()`: Returns "not yet implemented" error
  - `Stop/Start/Reboot`: Return "not yet implemented" errors
- **Real API Integration**: Framework ready for actual Nebius compute API calls

## Prerequisites

### 1. Development Environment
```bash
# Minimum Go version
go version # Should be >= 1.22

# Nebius SDK dependency
go list -m github.com/nebius/gosdk
# Should show: github.com/nebius/gosdk v0.0.0-20250826102719-940ad1dfb5de

# Required testing dependencies
go list -m github.com/stretchr/testify
# Should show: github.com/stretchr/testify v1.11.0
```

### 2. Nebius Account Setup
- Nebius AI Cloud account with billing enabled
- Service account with appropriate compute permissions
- Service account key pair (JSON format preferred)
- Folder ID (Nebius equivalent to project in other clouds)
- Access to target regions (e.g., eu-north1)

### 3. Nebius Authentication Setup

#### Recommended: Service Account Credentials

Nebius AI Cloud supports multiple authentication methods. For production use, service account credentials are strongly recommended.

##### Option A: Service Account JSON File (Preferred)
Create a service account in the Nebius AI Console and download the JSON credentials file:

```json
{
  "id": "service-account-id",
  "service_account_id": "your-service-account-id",
  "created_at": "2024-01-01T00:00:00Z",
  "key_algorithm": "RSA_2048",
  "public_key": "-----BEGIN PUBLIC KEY-----\\n...\\n-----END PUBLIC KEY-----\\n",
  "private_key": "-----BEGIN PRIVATE KEY-----\\n...\\n-----END PRIVATE KEY-----\\n"
}
```

##### Option B: Separate Private Key File
Alternatively, store the private key in a separate PEM file:

**service_account.json:**
```json
{
  "service_account_id": "your-service-account-id",
  "key_id": "your-key-id"
}
```

**private_key.pem:**
```
-----BEGIN PRIVATE KEY-----
YOUR_PRIVATE_KEY_CONTENT_HERE
-----END PRIVATE KEY-----
```

##### Option C: IAM Token (Development Only)
For quick testing or development environments, you can use an IAM token directly:

```bash
export NEBIUS_IAM_TOKEN="your-iam-token"
```

**⚠️ Note:** IAM tokens require manual refresh and are not recommended for production use.

#### Obtaining Credentials

1. **Access Nebius AI Console**: Log into https://console.nebius.ai
2. **Create Service Account**:
   - Navigate to IAM & Admin > Service Accounts
   - Click "Create Service Account"
   - Assign necessary permissions (Compute Admin, etc.)
3. **Generate Key Pair**:
   - Select your service account
   - Go to "Keys" tab
   - Click "Add Key" > "Create new key"
   - Choose JSON format and download

```bash
export SA_ID=$(nebius iam service-account get-by-name \
  --name jmorgan-sa \
  --format json \
  | jq -r ".metadata.id")

nebius iam auth-public-key generate \
  --service-account-id $SA_ID \
  --output ~/.nebius/$SA_ID-credentials.json
```

4. **Set Environment Variables**:
   ```bash
   export NEBIUS_SERVICE_ACCOUNT_JSON="/path/to/service-account.json"
   export NEBIUS_TENANT_ID="your-tenant-id"
   export NEBIUS_LOCATION="eu-north1"  # Optional, defaults to eu-north1
   ```

#### Required Permissions
Your service account needs these IAM roles:
- `compute.admin` - For instance management
- `vpc.admin` - For networking (if using VPC features)
- `iam.serviceAccountUser` - For service account operations

## Build and Testing

### 1. Build the Provider
```bash
# Build all Nebius provider components
go build ./v1/providers/nebius/...

# Build entire SDK to ensure integration
go build ./...

# Run static analysis
go vet ./v1/providers/nebius/...
golangci-lint run ./v1/providers/nebius/...
```

### 2. Unit Testing
```bash
# Run all unit tests
go test ./v1/providers/nebius/... -v

# Run tests with coverage
go test ./v1/providers/nebius/... -cover -coverprofile=nebius.out
go tool cover -html=nebius.out

# Run specific test files
go test ./v1/providers/nebius/ -run TestNebiusCredential -v
go test ./v1/providers/nebius/ -run TestNebiusClient -v

# Run benchmarks
go test ./v1/providers/nebius/... -bench=. -benchmem
```

### 3. Integration Testing Framework

#### Test Structure Overview
The Nebius provider includes comprehensive test suites:

1. **Unit Tests** (`*_test.go`): Test individual functions and methods
2. **Integration Tests** (`integration_test.go`): Test against real Nebius API
3. **Smoke Tests** (`smoke_test.go`): End-to-end instance lifecycle testing

#### Running Unit Tests
```bash
# All unit tests
go test ./v1/providers/nebius/ -v

# Specific test suites
go test ./v1/providers/nebius/ -run TestNebiusCredential -v
go test ./v1/providers/nebius/ -run TestNebiusClient_CreateInstance -v
go test ./v1/providers/nebius/ -run TestNebiusClient_NotImplementedMethods -v
```

#### Running Integration Tests
```bash
# Set up credentials
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='your-tenant-id'

# Run integration tests (requires real credentials)
go test ./v1/providers/nebius/ -run TestIntegration -v

# Skip integration tests in CI/short mode
go test ./v1/providers/nebius/ -short -v
```

#### Running Smoke Tests (End-to-End)

**✅ Current Implementation Status**: The smoke test creates **actual Nebius cloud instances** for true end-to-end validation:
- ✅ **CreateInstance**: Creates real L40S GPU instances in Nebius cloud
- ✅ **GetInstance**: Retrieves and validates actual instance data
- ✅ **TerminateInstance**: Properly cleans up cloud resources
- ✅ **Platform Targeting**: Supports L40S GPU and custom configurations
- ✅ **Architecture Compatibility**: Uses working x86_64 image families
- ✅ **Resource Cleanup**: Automated cleanup with manual fallback options

```bash
# Enable smoke tests with proper credentials
export RUN_SMOKE_TESTS=true
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='your-tenant-id'
export NEBIUS_LOCATION='eu-north1'  # Optional, defaults to eu-north1

# Run comprehensive instance lifecycle test (creates real cloud resources)
go test ./v1/providers/nebius/ -run TestSmoke_InstanceLifecycle -v -timeout=15m

# Run with cleanup (recommended)
CLEANUP_RESOURCES=true RUN_SMOKE_TESTS=true go test ./v1/providers/nebius/ -run TestSmoke_InstanceLifecycle -v -timeout=15m

# Target specific platforms and configurations
NEBIUS_TARGET_PLATFORM=l40s NEBIUS_DISK_SIZE_GB=50 CLEANUP_RESOURCES=true RUN_SMOKE_TESTS=true go test ./v1/providers/nebius/ -run TestSmoke_InstanceLifecycle -v -timeout=15m
```

### Manual Cleanup Guide for Smoke Test Resources

If smoke tests fail or cleanup doesn't complete properly, use these commands to manually clean up resources with `smoke-test-*` names:

#### Prerequisites
```bash
# Install Nebius CLI if not already installed
curl -sSfL https://storage.googleapis.com/nebius-cli/install.sh | bash

# Set up authentication (use same credentials as for tests)
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='your-tenant-id'
nebius init
```

#### 1. Cleanup Instances

```bash
# List smoke test instances
nebius compute instance list --parent-id PROJECT_ID | grep "smoke-test-"

# Delete specific instance
nebius compute instance delete INSTANCE_ID

# Bulk delete smoke test instances (requires jq)
for instance_id in $(nebius compute instance list --parent-id PROJECT_ID --format json | jq -r '.items[] | select(.metadata.name | startswith("smoke-test-")) | .metadata.id'); do
  echo "Deleting instance: $instance_id"
  nebius compute instance delete $instance_id
done
```

#### 2. Cleanup Disks

```bash
# List smoke test disks
nebius compute disk list --parent-id PROJECT_ID | grep "smoke-test-"

# Delete specific disk (after instances are terminated)
nebius compute disk delete DISK_ID

# Bulk delete smoke test disks
for disk_id in $(nebius compute disk list --parent-id PROJECT_ID --format json | jq -r '.items[] | select(.metadata.name | startswith("smoke-test-")) | .metadata.id'); do
  echo "Deleting disk: $disk_id"
  nebius compute disk delete $disk_id
done
```

#### 3. Cleanup Networks and Subnets

```bash
# List smoke test subnets
nebius vpc subnet list --parent-id PROJECT_ID | grep "smoke-test-"

# Delete specific subnet
nebius vpc subnet delete SUBNET_ID

# Bulk delete smoke test subnets
for subnet_id in $(nebius vpc subnet list --parent-id PROJECT_ID --format json | jq -r '.items[] | select(.metadata.name | startswith("smoke-test-")) | .metadata.id'); do
  echo "Deleting subnet: $subnet_id"
  nebius vpc subnet delete $subnet_id
done

# List smoke test networks
nebius vpc network list --parent-id PROJECT_ID | grep "smoke-test-"

# Delete specific network (after subnets are deleted)
nebius vpc network delete NETWORK_ID

# Bulk delete smoke test networks
for network_id in $(nebius vpc network list --parent-id PROJECT_ID --format json | jq -r '.items[] | select(.metadata.name | startswith("smoke-test-")) | .metadata.id'); do
  echo "Deleting network: $network_id"
  nebius vpc network delete $network_id
done
```

#### 4. Cleanup Project (if created for testing)

```bash
# List projects with brev-user prefix
nebius iam project list --parent-id TENANT_ID | grep "brev-user-"

# Delete test project (this will delete all resources within)
nebius iam project delete PROJECT_ID

# ⚠️ WARNING: This deletes the entire project and all resources within it
# Only use if the project was created specifically for testing
```

#### Complete Cleanup Script

Create a script for comprehensive cleanup:

```bash
#!/bin/bash
# complete-cleanup.sh - Clean up all smoke-test resources

set -e  # Exit on error

PROJECT_ID="${NEBIUS_PROJECT_ID:-$(echo 'Set NEBIUS_PROJECT_ID environment variable')}"
TENANT_ID="${NEBIUS_TENANT_ID:-$(echo 'Set NEBIUS_TENANT_ID environment variable')}"

if [[ -z "$PROJECT_ID" || -z "$TENANT_ID" ]]; then
    echo "❌ Required environment variables not set"
    echo "   export NEBIUS_PROJECT_ID='your-project-id'"
    echo "   export NEBIUS_TENANT_ID='your-tenant-id'"
    exit 1
fi

echo "🧹 Starting complete cleanup of smoke-test resources..."
echo "   Project: $PROJECT_ID"
echo "   Tenant: $TENANT_ID"

# Function to safely delete resources
delete_resources() {
    local resource_type=$1
    local list_cmd=$2
    local delete_cmd=$3

    echo "🗑️  Cleaning up ${resource_type}s..."

    ids=$(eval "$list_cmd" 2>/dev/null | jq -r '.items[]? | select(.metadata.name | startswith("smoke-test-")) | .metadata.id' || echo "")

    if [[ -z "$ids" ]]; then
        echo "   No smoke-test ${resource_type}s found"
        return
    fi

    for id in $ids; do
        echo "   Deleting $resource_type: $id"
        eval "$delete_cmd $id" || echo "   Failed to delete $id (may already be deleted)"
    done
}

# 1. Delete instances first
delete_resources "instance" \
    "nebius compute instance list --parent-id $PROJECT_ID --format json" \
    "nebius compute instance delete"

# Wait for instances to terminate
echo "⏳ Waiting for instances to terminate..."
sleep 30

# 2. Delete disks (should be detached after instance deletion)
delete_resources "disk" \
    "nebius compute disk list --parent-id $PROJECT_ID --format json" \
    "nebius compute disk delete"

# 3. Delete subnets
delete_resources "subnet" \
    "nebius vpc subnet list --parent-id $PROJECT_ID --format json" \
    "nebius vpc subnet delete"

# 4. Delete networks
delete_resources "network" \
    "nebius vpc network list --parent-id $PROJECT_ID --format json" \
    "nebius vpc network delete"

# 5. Optionally delete test project
read -p "🗑️  Delete test project $PROJECT_ID? This will remove ALL resources in the project. (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "🗑️  Deleting project: $PROJECT_ID"
    nebius iam project delete $PROJECT_ID || echo "Failed to delete project (may not exist)"
else
    echo "   Project preserved"
fi

echo "✅ Cleanup completed!"

# Verify cleanup
echo "🔍 Verification - remaining smoke-test resources:"
echo "   Instances: $(nebius compute instance list --parent-id $PROJECT_ID --format json 2>/dev/null | jq -r '.items[]? | select(.metadata.name | startswith("smoke-test-")) | .metadata.name' | wc -l || echo '0')"
echo "   Disks: $(nebius compute disk list --parent-id $PROJECT_ID --format json 2>/dev/null | jq -r '.items[]? | select(.metadata.name | startswith("smoke-test-")) | .metadata.name' | wc -l || echo '0')"
echo "   Subnets: $(nebius vpc subnet list --parent-id $PROJECT_ID --format json 2>/dev/null | jq -r '.items[]? | select(.metadata.name | startswith("smoke-test-")) | .metadata.name' | wc -l || echo '0')"
echo "   Networks: $(nebius vpc network list --parent-id $PROJECT_ID --format json 2>/dev/null | jq -r '.items[]? | select(.metadata.name | startswith("smoke-test-")) | .metadata.name' | wc -l || echo '0')"
```

Save as `cleanup-smoke-test.sh`, make executable with `chmod +x cleanup-smoke-test.sh`, and run:

```bash
export NEBIUS_PROJECT_ID="your-project-id"
export NEBIUS_TENANT_ID="your-tenant-id"
./cleanup-smoke-test.sh
```

### What the Smoke Test Actually Does

The smoke test (`TestSmoke_InstanceLifecycle`) is a **comprehensive end-to-end test framework** that exercises the full instance lifecycle. Here's what happens when you run it:

#### ✅ **Current Behavior** (Mock Implementation):
1. **Authentication Test**: ✅ Connects to real Nebius API using your service account
2. **Project Creation**: ✅ Generates project ID for your user (`brev-{hash}`)
3. **Mock Instance Creation**: ✅ Returns mock instance data (no real VM)
4. **Mock Instance Get**: ✅ Returns mock instance data
5. **Lifecycle Operations**: ❌ Fail with "not yet implemented" (expected)

#### 🚀 **Future Behavior** (When SDK Integration Complete):
1. **Real Instance Creation**: Creates actual Nebius VM in your project
2. **Instance Verification**: Checks VM exists and is accessible
3. **Power Management**: Tests stop/start/reboot operations
4. **Resource Management**: Updates tags, resizes volumes
5. **Cleanup**: Terminates VM and verifies deletion

### Expected Test Output

When you run the smoke test currently, you'll see:
```
🚀 Starting Nebius smoke test with ID: smoke-test-1727123456
✅ Authentication successful! (connects to real Nebius API)
✅ Project ID generated: brev-f85ac825d102
✅ Step 1: Mock instance created
✅ Step 2: Mock instance verified
❌ Step 3: Stop instance failed - "not yet implemented" (expected)
```

The test **validates your authentication and project setup** but doesn't create real VMs yet.

### Quick Authentication Test

To verify your credentials are working without running the full smoke test:

```bash
# Test authentication only
export NEBIUS_SERVICE_ACCOUNT_JSON='/home/jmorgan/.nebius/serviceaccount-e00r1azfy8hw51q1fq-credentials.json'
export NEBIUS_TENANT_ID='tenant-e00eb38h7v3ph9b343'

go test ./v1/providers/nebius/ -run TestIntegration_ClientCreation -v
```

Expected output:
```
✅ Authentication successful!
✅ Client created with project-per-user model: brev-f85ac825d102
```

This confirms:
- ✅ Service account JSON format is correct
- ✅ Nebius SDK authentication works
- ✅ Project-per-user mapping is functional
- ✅ Ready for real instance operations

## API Integration Testing Guidelines

### 1. Test Environment Setup

#### Local Development
```bash
# Set up credentials for testing
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='your-tenant-id'
export NEBIUS_LOCATION='eu-north1'  # Optional

# Enable debug logging
export NEBIUS_DEBUG=true
export NEBIUS_LOG_LEVEL=debug
```

#### CI/CD Environment
```yaml
# Example GitHub Actions setup
env:
  NEBIUS_SERVICE_ACCOUNT_JSON: ${{ secrets.NEBIUS_SERVICE_ACCOUNT_JSON }}
  NEBIUS_TENANT_ID: ${{ secrets.NEBIUS_TENANT_ID }}
  RUN_SMOKE_TESTS: 'false'  # Disable destructive tests in CI
```

### 2. Test Categories and Execution

#### Unit Tests (No External Dependencies)
```bash
# Fast tests for development
go test ./v1/providers/nebius/ -short -v

# With coverage
go test ./v1/providers/nebius/ -short -cover -coverprofile=unit.out
go tool cover -html=unit.out
```

#### Integration Tests (Requires API Access)
```bash
# Test authentication and basic API calls
go test ./v1/providers/nebius/ -run TestIntegration -v

# Test specific integration scenarios
go test ./v1/providers/nebius/ -run TestIntegration_GetCapabilities -v
go test ./v1/providers/nebius/ -run TestIntegration_GetLocations -v
go test ./v1/providers/nebius/ -run TestIntegration_ErrorHandling -v
```

#### Smoke Tests (Full Instance Lifecycle)
```bash
# Complete end-to-end testing
export RUN_SMOKE_TESTS=true
go test ./v1/providers/nebius/ -run TestSmoke -v -timeout=15m

# Individual smoke test operations
go test ./v1/providers/nebius/ -run TestSmoke_InstanceLifecycle -v -timeout=15m
```

### 3. Performance and Load Testing

#### Benchmarking
```bash
# Benchmark instance creation
go test -bench=BenchmarkCreateInstance ./v1/providers/nebius/ -benchtime=10s

# Memory profiling
go test -bench=BenchmarkCreateInstance -memprofile=mem.prof ./v1/providers/nebius/
go tool pprof mem.prof

# CPU profiling
go test -bench=. -cpuprofile=cpu.prof ./v1/providers/nebius/
go tool pprof cpu.prof
```

#### Rate Limit Testing
```bash
# Test API rate limits
go test ./v1/providers/nebius/ -run TestIntegration -count=10 -parallel=5
```

### 4. Test Data Management

#### Instance Naming Convention
```go
// Format: {test-type}-{timestamp}-{random}
testInstanceName := fmt.Sprintf("test-instance-%d-%s",
    time.Now().Unix(),
    generateRandomString(8))
```

#### Cleanup Strategy
```bash
# Tag all test resources for automated cleanup
Tags: map[string]string{
    "test-type":    "automated",
    "created-by":   "nebius-integration-test",
    "auto-delete":  "true",
    "ttl-hours":    "2",  // Auto-cleanup after 2 hours
}
```

#### Manual Cleanup
```bash
# List test instances for manual cleanup
# (requires implementation of ListInstances)
go run tools/cleanup-test-instances.go -tenant-id="$NEBIUS_TENANT_ID" -dry-run
```

### 5. Test Execution Strategies

#### Development Workflow
```bash
# Quick development cycle
go test ./v1/providers/nebius/ -short -v  # Unit tests only

# Before committing
go test ./v1/providers/nebius/ -run TestIntegration_ClientCreation -v
go test ./v1/providers/nebius/ -cover
```

#### Pre-deployment Testing
```bash
# Comprehensive validation
go test ./v1/providers/nebius/ -v  # All tests
export RUN_SMOKE_TESTS=true
go test ./v1/providers/nebius/ -run TestSmoke -v -timeout=20m
```

#### Continuous Integration
```bash
# CI-safe test run (no destructive operations)
go test ./v1/providers/nebius/ -short -v
go test ./v1/providers/nebius/ -run TestIntegration_GetCapabilities -v
# Smoke tests disabled in CI unless explicitly enabled
```

### 6. Error Scenarios and Edge Cases

#### Authentication Error Testing
```bash
# Test with invalid credentials
NEBIUS_SERVICE_ACCOUNT_JSON='{"invalid": "json"}' \
go test ./v1/providers/nebius/ -run TestIntegration_ErrorHandling -v
```

#### Network and Timeout Testing
```bash
# Test with network issues (using network simulation)
go test ./v1/providers/nebius/ -run TestIntegration -timeout=30s
```

#### Resource Limit Testing
```bash
# Test quota and limit scenarios
go test ./v1/providers/nebius/ -run TestIntegration_ResourceLimits -v
```

## Development Workflow and Implementation Guide

### 1. Test-Driven Development Approach

#### Implementation Order (with corresponding tests):

1. **Authentication & Client Setup**
   ```bash
   # Implement and test credential handling
   go test ./v1/providers/nebius/ -run TestNebiusCredential -v
   go test ./v1/providers/nebius/ -run TestNebiusClient_Creation -v
   ```

2. **Core Instance Operations**
   ```bash
   # Implement CreateInstance -> GetInstance -> TerminateInstance
   go test ./v1/providers/nebius/ -run TestNebiusClient_CreateInstance -v
   go test ./v1/providers/nebius/ -run TestIntegration_InstanceLifecycle -v
   ```

3. **Instance Management**
   ```bash
   # Implement Stop/Start/Reboot operations
   go test ./v1/providers/nebius/ -run TestSmoke_InstanceLifecycle -v
   ```

4. **Resource Discovery**
   ```bash
   # Implement GetInstanceTypes and GetImages
   go test ./v1/providers/nebius/ -run TestIntegration_GetInstanceTypes -v
   go test ./v1/providers/nebius/ -run TestIntegration_GetImages -v
   ```

### 2. Implementation Testing Strategy

#### For Each New Method Implementation:
1. **Write failing unit test first**
2. **Implement minimal functionality**
3. **Run integration test with real API**
4. **Add to smoke test suite**
5. **Update documentation**

#### Example Implementation Cycle:
```bash
# 1. Write test
go test ./v1/providers/nebius/ -run TestGetInstanceTypes -v  # Should fail

# 2. Implement method in instancetype.go
# 3. Test implementation
go test ./v1/providers/nebius/ -run TestGetInstanceTypes -v  # Should pass

# 4. Integration test
go test ./v1/providers/nebius/ -run TestIntegration_GetInstanceTypes -v

# 5. Add to smoke test
export RUN_SMOKE_TESTS=true
go test ./v1/providers/nebius/ -run TestSmoke -v
```

### 3. Testing New Implementations

#### Method-Specific Testing
```bash
# Test individual method implementations
go test ./v1/providers/nebius/ -run TestNebiusClient_GetInstanceTypes -v
go test ./v1/providers/nebius/ -run TestNebiusClient_CreateInstance -v
go test ./v1/providers/nebius/ -run TestNebiusClient_TerminateInstance -v
```

#### Cross-Method Integration
```bash
# Test method interactions (create -> get -> terminate)
go test ./v1/providers/nebius/ -run TestIntegration_InstanceLifecycle -v
```

### 4. Integration with Brev Backend

#### Local Development Server
```bash
# Set up environment for backend integration
export BREV_CLOUD_SDK_PATH="$(pwd)"
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='your-tenant-id'

# Start local backend with Nebius provider
go run ../brev-backend/cmd/server/main.go --cloud-provider nebius --debug
```

#### Backend Integration Testing
```bash
# Test SDK integration with Brev backend
curl -X POST http://localhost:8080/api/instances \
  -H "Content-Type: application/json" \
  -d '{
    "provider": "nebius",
    "instance_type": "standard-2",
    "image_id": "ubuntu-20.04",
    "name": "integration-test"
  }'
```

## Testing Troubleshooting and Common Issues

### 1. Test Environment Issues

#### Authentication Test Failures
**Problem**: `"failed to initialize Nebius SDK"` or `"invalid service account"`
**Solutions**:
```bash
# Verify JSON format
cat $NEBIUS_SERVICE_ACCOUNT_JSON | jq .  # Should parse without errors

# Check required fields
jq -r '.service_account_id, .private_key' $NEBIUS_SERVICE_ACCOUNT_JSON

# Test with minimal credentials
echo '{
  "service_account_id": "test",
  "private_key": "test"
}' | go test ./v1/providers/nebius/ -run TestNebiusCredential_ValidJSON -v
```

#### Integration Test Skipping
**Problem**: Integration tests are being skipped
**Solutions**:
```bash
# Ensure environment variables are set
echo "Service Account: $NEBIUS_SERVICE_ACCOUNT_JSON"
echo "Folder ID: $NEBIUS_TENANT_ID"

# Run with explicit credential check
go test ./v1/providers/nebius/ -run TestIntegration_ClientCreation -v
```

### 2. Test Execution Issues

#### Smoke Test Failures
**Problem**: Smoke tests fail or timeout
**Solutions**:
```bash
# Increase timeout for slower operations
go test ./v1/providers/nebius/ -run TestSmoke -timeout=20m -v

# Run individual smoke test steps
go test ./v1/providers/nebius/ -run TestSmoke_InstanceLifecycle -v

# Check test resource cleanup
export RUN_SMOKE_TESTS=true
go test ./v1/providers/nebius/ -run TestSmoke -v -cleanup=true
```

#### Rate Limiting Issues
**Problem**: API rate limit exceeded during tests
**Solutions**:
```bash
# Run tests with delays
go test ./v1/providers/nebius/ -parallel=1 -v

# Use test-specific credentials with higher limits
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/testing-service-account.json'
```

### 3. Implementation Testing Issues

#### "Not Implemented" Method Testing
**Problem**: Tests fail because methods aren't fully implemented
**Expected Behavior**:
```bash
# These should pass even with placeholder implementation
go test ./v1/providers/nebius/ -run TestNebiusClient_NotImplementedMethods -v

# Integration tests should handle not-implemented gracefully
go test ./v1/providers/nebius/ -run TestIntegration_InstanceLifecycle -v
```

#### Build and Import Issues
**Problem**: Import path or dependency issues
**Solutions**:
```bash
# Clean and rebuild
go clean -modcache
go mod download
go mod tidy

# Verify imports
go list -m github.com/nebius/gosdk
go list -m github.com/brevdev/cloud
```

### 4. Test Resource Management

#### Orphaned Test Resources
**Problem**: Test instances not cleaned up properly
**Prevention**:
```bash
# Always use consistent tagging
Tags: map[string]string{
    "created-by":   "nebius-integration-test",
    "test-run-id":  testRunID,
    "auto-delete":  "true",
}

# Manual cleanup (when ListInstances is implemented)
go run tools/cleanup-test-resources.go -tenant-id=$NEBIUS_TENANT_ID
```

#### Test Data Conflicts
**Problem**: Tests interfere with each other
**Solutions**:
```bash
# Use unique test identifiers
testID := fmt.Sprintf("test-%d-%s", time.Now().Unix(), randomString(8))

# Run tests sequentially if needed
go test ./v1/providers/nebius/ -parallel=1 -v
```

### 5. Debug and Monitoring

#### Test Debugging
```bash
# Enable verbose SDK logging
export NEBIUS_DEBUG=true
export NEBIUS_LOG_LEVEL=debug

# Run single test with maximum verbosity
go test ./v1/providers/nebius/ -run TestSmoke_InstanceLifecycle -v -count=1

# Use test timeout to prevent hanging
go test ./v1/providers/nebius/ -timeout=5m -v
```

#### Performance Issues
```bash
# Profile test execution
go test -bench=. -memprofile=mem.prof -cpuprofile=cpu.prof ./v1/providers/nebius/
go tool pprof mem.prof
go tool pprof cpu.prof

# Memory leak detection
go test -run TestIntegration -memprofile=mem.prof ./v1/providers/nebius/
go tool pprof -alloc_space mem.prof
```

## Production Readiness and Testing Checklist

### Testing Completeness Checklist

#### Unit Testing Requirements
- [x] Client creation and configuration tests
- [x] Credential validation tests
- [x] Method signature and return value tests
- [x] Error handling and edge case tests
- [x] Benchmark tests for performance
- [ ] Mock SDK integration tests (when SDK interface is stable)
- [ ] Concurrent operation tests
- [ ] Memory leak detection tests

#### Integration Testing Requirements
- [x] Authentication with real Nebius API
- [x] Basic capability and location queries
- [x] Error handling with invalid credentials
- [ ] Instance creation with real API
- [ ] Instance lifecycle operations (stop/start/reboot)
- [ ] Resource discovery (instance types, images)
- [ ] Instance management (tags, volume resize)
- [ ] Network and timeout handling
- [ ] Rate limiting and retry logic

#### Smoke Testing Requirements
- [x] End-to-end instance lifecycle test
- [x] Proper test resource cleanup
- [x] Multi-operation workflow testing
- [ ] Performance under load
- [ ] Long-running operation handling
- [ ] Failure recovery testing

### Implementation Readiness Checklist

#### Core Functionality
- [x] Client authentication and initialization
- [x] Basic instance operations (create/get placeholder)
- [ ] **GetInstanceTypes** - List available VM configurations
- [ ] **GetImages** - List available base images
- [ ] **CreateInstance** - Full VM creation with Nebius API
- [ ] **ListInstances** - Bulk instance listing
- [ ] **TerminateInstance** - Instance deletion
- [ ] **StopInstance/StartInstance** - Power management
- [ ] **RebootInstance** - Restart functionality
- [ ] **UpdateInstanceTags** - Tag management
- [ ] **ResizeInstanceVolume** - Storage management

#### Error Handling and Resilience
- [ ] Comprehensive error wrapping and context
- [ ] Proper logging integration
- [ ] Rate limiting and retry logic with exponential backoff
- [ ] Circuit breaker for API failures
- [ ] Timeout handling for long operations
- [ ] Graceful degradation for partial failures

#### Security Implementation
- [ ] Service account key secure parsing and handling
- [ ] No credentials in logs or error messages
- [ ] Proper IAM permission scope validation
- [ ] TLS verification for API connections
- [ ] Input validation and sanitization
- [ ] Audit logging for sensitive operations

#### Performance and Scalability
- [ ] Connection pooling and reuse
- [ ] Request batching where applicable
- [ ] Caching of frequently accessed data
- [ ] Performance benchmarks established and met
- [ ] Memory usage optimization
- [ ] Concurrent operation support

### Test Execution Checklist

#### Pre-commit Testing
```bash
# Run before every commit
go test ./v1/providers/nebius/ -short -v                    # Unit tests
go test ./v1/providers/nebius/ -cover -coverprofile=cov.out # Coverage check
go vet ./v1/providers/nebius/...                            # Static analysis
golangci-lint run ./v1/providers/nebius/...                 # Linting
```

#### Pre-deployment Testing
```bash
# Comprehensive validation before deployment
go test ./v1/providers/nebius/ -v                         # All tests
go test ./v1/providers/nebius/ -run TestIntegration -v    # Integration tests
export RUN_SMOKE_TESTS=true
go test ./v1/providers/nebius/ -run TestSmoke -timeout=20m # End-to-end tests
go test -bench=. ./v1/providers/nebius/                   # Performance tests
```

#### Production Deployment Validation
```bash
# Post-deployment smoke test in production environment
export NEBIUS_SERVICE_ACCOUNT_JSON="$PROD_SERVICE_ACCOUNT"
export NEBIUS_TENANT_ID="$PROD_FOLDER_ID"
export RUN_SMOKE_TESTS=true
go test ./v1/providers/nebius/ -run TestSmoke_InstanceLifecycle -v -timeout=15m
```

## Monitoring and Observability

### 1. Metrics to Track
- Client creation latency
- API call success/failure rates
- Instance operation durations
- Error distribution by type

### 2. Logging Best Practices
```go
// Use structured logging
logger := log.FromContext(ctx).WithValues(
    "provider", "nebius",
    "operation", "CreateInstance",
    "folderID", c.folderID,
)

logger.Info("Creating instance", "name", attrs.Name)
```

### 3. Error Reporting
- Implement proper error categorization
- Add retry logic for transient failures
- Report metrics to monitoring system

## Support and Troubleshooting

### Debug Environment Variables
```bash
export NEBIUS_DEBUG=true           # Enable debug logging
export NEBIUS_API_TIMEOUT=30s      # API timeout
export NEBIUS_RETRY_ATTEMPTS=3     # Retry logic
```

### Common Debug Commands
```bash
# Check SDK connectivity
go run tools/nebius-debug.go connectivity

# Validate credentials
go run tools/nebius-debug.go auth-test

# List available resources
go run tools/nebius-debug.go list-resources
```

### Testing Resources and References

#### Documentation
1. **Nebius AI Cloud API Documentation**: https://docs.nebius.ai/
2. **Nebius Go SDK**: https://github.com/nebius/gosdk
3. **Brev Cloud SDK Patterns**: Review other provider implementations
   - `v1/providers/lambdalabs/` - Similar cloud provider pattern
   - `v1/providers/fluidstack/` - Instance lifecycle examples

#### Test Execution Examples

**Development Testing:**
```bash
# Quick development loop
go test ./v1/providers/nebius/ -short -v

# With real API testing
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/creds.json'
export NEBIUS_TENANT_ID='your-folder'
go test ./v1/providers/nebius/ -run TestIntegration -v
```

**Production Validation:**
```bash
# Full end-to-end validation
export RUN_SMOKE_TESTS=true
export NEBIUS_SERVICE_ACCOUNT_JSON="$PROD_CREDS"
export NEBIUS_TENANT_ID="$PROD_FOLDER"
go test ./v1/providers/nebius/ -run TestSmoke -timeout=20m -v
```

**Continuous Integration:**
```bash
# CI-safe testing (no destructive operations)
go test ./v1/providers/nebius/ -short -cover -v
if [[ "$CI_BRANCH" == "main" ]]; then
  go test ./v1/providers/nebius/ -run TestIntegration_GetCapabilities -v
fi
```

### Getting Help
1. **Testing Issues**: Check the troubleshooting section above
2. **API Integration**: Review Nebius AI Cloud documentation
3. **SDK Usage**: Examine Nebius Go SDK examples and documentation
4. **Provider Patterns**: Study existing provider implementations in the codebase
5. **Nebius Support**: Contact support for API-specific questions
6. **Brev Integration**: Review Brev Cloud SDK integration patterns

---

## Instance Type Enumeration

### Overview

The Nebius provider implements **quota-aware instance type discovery** that dynamically returns available instance types based on:
1. **Active quota allocations** across all regions
2. **Any GPU platform** with available quota (L40S, H100, H200, A100, V100, etc.)
3. **Supported CPU platforms**: cpu-d3, cpu-e2 (limited to 3 presets each)
4. **Available presets** per platform (e.g., 1, 2, 4, 8 GPUs)

### How Instance Types Are Discovered

#### 1. Quota-Based Filtering

The provider queries the Nebius Quotas API to determine which resources are available:

```go
// Example quota lookups
"compute.gpu.h100:eu-north1"  // H100 GPUs in eu-north1
"compute.gpu.h200:eu-north1"  // H200 GPUs in eu-north1
"compute.gpu.l40s:eu-north1"  // L40S GPUs in eu-north1
"compute.cpu:eu-north1"       // vCPU quota for CPU instances
"compute.memory:eu-north1"    // Memory quota for CPU instances
```

**Key Behavior**:
- Only instance types with **active quota** (State: ACTIVE) are returned
- Instance types are filtered by **available capacity** (Limit - Usage > 0)
- If **no quota exists** for a GPU type in a region, those instance types are excluded
- For GPU instances, quota is checked per GPU count (e.g., 4x L40S requires 4 GPUs available)

#### 2. Platform Filtering

**GPU Platforms:**
- ✅ **Dynamically discovered** - Any GPU platform with available quota is included
- ✅ No hardcoded restrictions (L40S, H100, H200, A100, V100, A10, T4, L4, etc.)
- ✅ Filtered only by quota availability

**CPU Platforms:**
- ✅ **Explicitly filtered** to cpu-d3 and cpu-e2 only
- ✅ **Limited to 3 presets per platform** to avoid list pollution
- ✅ Other CPU platforms are excluded even if they have quota

```go
// Example: If you have quota for these GPUs, they will ALL appear:
- "H100"  // NVIDIA H100 (80GB HBM3)
- "H200"  // NVIDIA H200 (141GB HBM3e)
- "L40S"  // NVIDIA L40S (48GB GDDR6)
- "A100"  // NVIDIA A100 (40GB/80GB)
- "V100"  // NVIDIA V100 (16GB/32GB)

// CPU Platforms (only these two, max 3 presets each):
- "cpu-d3" // Intel Ice Lake (first 3 presets only)
- "cpu-e2" // AMD EPYC (first 3 presets only)
```

#### 3. Preset Enumeration

Each platform exposes **multiple presets** based on GPU count and resource configuration:

```
Platform: L40S
├── Preset: 1gpu-24vcpu-200gb    (1x L40S, 24 vCPU, 200GB RAM)
├── Preset: 2gpu-48vcpu-400gb    (2x L40S, 48 vCPU, 400GB RAM)
├── Preset: 4gpu-96vcpu-800gb    (4x L40S, 96 vCPU, 800GB RAM)
└── Preset: 8gpu-192vcpu-1600gb  (8x L40S, 192 vCPU, 1600GB RAM)
```

**Instance Type ID Format**: `{platform-id}-{preset-name}`
Example: `computeplatform-e00abc123-8gpu-192vcpu-1600gb`

### Elastic Disk Support

All Nebius instance types support **dynamically allocatable network SSD disks**:

```go
Storage Configuration:
├── Type: "network-ssd"
├── Min Size: 50 GB
├── Max Size: 2560 GB
├── Elastic: true
└── Price: ~$0.00014 per GB-hour
```

This is exposed via:
- `InstanceType.ElasticRootVolume = true`
- `InstanceType.SupportedStorage[0].IsElastic = true`
- `InstanceType.SupportedStorage[0].MinSize = 50GB`
- `InstanceType.SupportedStorage[0].MaxSize = 2560GB`

### Testing Instance Type Enumeration

#### Manual Enumeration Test

```bash
# Set up credentials
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='tenant-e00xxx'

# Run the instance types integration test
go test ./v1/providers/nebius/ -run TestIntegration_GetInstanceTypes -v

# Expected output:
# === RUN   TestIntegration_GetInstanceTypes
# === RUN   TestIntegration_GetInstanceTypes/Get_instance_types_with_quota_filtering
#     Found 12 instance types with available quota
#     Instance Type: computeplatform-e00abc-1gpu (...) - Location: eu-north1, Available: true
#       Storage: network-ssd, Min: 50 GB, Max: 2560 GB, Elastic: true
#       GPU: NVIDIA L40S (Type: L40S), Count: 1, Manufacturer: NVIDIA
# === RUN   TestIntegration_GetInstanceTypes/Filter_by_supported_platforms
#     Instance type distribution:
#       L40S: 4
#       H100: 4
#       H200: 4
#       CPU-only: 0
# === RUN   TestIntegration_GetInstanceTypes/Verify_preset_enumeration
#     Preset enumeration by platform:
#       L40S: 4 presets
#         - computeplatform-e00abc-1gpu
#         - computeplatform-e00abc-2gpu
#         - computeplatform-e00abc-4gpu
#         - computeplatform-e00abc-8gpu
```

#### Programmatic Enumeration

```go
import (
    v1 "github.com/brevdev/cloud/v1"
    nebius "github.com/brevdev/cloud/v1/providers/nebius"
)

// Get all instance types with available quota
instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})

// Filter by specific location
instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{
    Locations: v1.LocationsFilter{"eu-north1"},
})

// Filter by GPU manufacturer
instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{
    GPUManufactererFilter: &v1.GPUManufacturerFilter{
        IncludeGPUManufacturers: []v1.Manufacturer{v1.ManufacturerNVIDIA},
    },
})
```

### Expected Output Structure

Each returned instance type includes:

```go
InstanceType{
    ID:                "computeplatform-e00abc123-4gpu-96vcpu-800gb",
    Location:          "eu-north1",
    Type:              "L40S Platform (4gpu-96vcpu-800gb)",
    VCPU:              96,
    Memory:            858993459200, // 800 GiB in bytes
    IsAvailable:       true,
    ElasticRootVolume: true,
    SupportedGPUs: []GPU{
        {
            Count:        4,
            Type:         "L40S",
            Name:         "NVIDIA L40S",
            Manufacturer: ManufacturerNVIDIA,
        },
    },
    SupportedStorage: []Storage{
        {
            Type:         "network-ssd",
            Count:        1,
            MinSize:      53687091200,    // 50 GiB
            MaxSize:      2748779069440,  // 2560 GiB
            IsElastic:    true,
            PricePerGBHr: &currency.Amount{Number: "0.00014", Currency: "USD"},
        },
    },
}
```

### Quota Management

#### Checking Current Quotas

```bash
# List all quota allowances for your tenant
nebius quotas quota-allowance list --parent-id TENANT_ID

# Check specific GPU quota
nebius quotas quota-allowance get-by-name \
  --parent-id TENANT_ID \
  --name "compute.gpu.l40s" \
  --region "eu-north1"
```

#### Understanding Quota States

```go
QuotaAllowanceStatus_State:
├── STATE_ACTIVE        // Quota is allocated and usable
├── STATE_PROVISIONING  // Quota is being allocated (not yet usable)
├── STATE_FROZEN        // Quota exists but cannot be used
└── STATE_DELETED       // Quota has been removed
```

**Only quotas in STATE_ACTIVE are considered available.**

### Troubleshooting Instance Type Enumeration

#### Problem: No Instance Types Returned

**Possible Causes**:
1. **No active quotas**: Check `nebius quotas quota-allowance list`
2. **Quotas fully consumed**: Check Usage vs Limit in quota status
3. **Wrong tenant ID**: Verify NEBIUS_TENANT_ID matches your organization
4. **Region mismatch**: Quotas are region-specific

**Solution**:
```bash
# Check quotas
export NEBIUS_TENANT_ID="tenant-e00xxx"
nebius quotas quota-allowance list --parent-id $NEBIUS_TENANT_ID --format json | \
  jq '.items[] | {name: .metadata.name, region: .spec.region, limit: .spec.limit, usage: .status.usage, state: .status.state}'

# Example output:
# {
#   "name": "compute.gpu.l40s",
#   "region": "eu-north1",
#   "limit": 8,
#   "usage": 0,
#   "state": "STATE_ACTIVE"
# }
```

#### Problem: Expected Platform Not Showing

**Check**:
1. Is the platform in the supported list? (L40S, H100, H200, cpu-d3, cpu-e2)
2. Does quota exist for that platform?
3. Is there available capacity (Limit - Usage > 0)?

```bash
# Check for specific GPU quota
nebius quotas quota-allowance list --parent-id $NEBIUS_TENANT_ID --format json | \
  jq '.items[] | select(.metadata.name | contains("gpu"))'
```

#### Problem: Wrong Number of Presets

**Explanation**: The number of presets depends on what Nebius has configured for each platform. Common configurations:
- **GPU platforms**: 1, 2, 4, 8 GPU presets
- **CPU platforms**: Various vCPU/memory combinations

If you see fewer presets than expected, check:
```bash
# List available platforms and their presets
nebius compute platform list --parent-id PROJECT_ID --format json | \
  jq '.items[] | {name: .metadata.name, presets: [.spec.presets[].name]}'
```

### Best Practices

1. **Cache Instance Types**: Results are relatively stable (poll every 5 minutes)
2. **Handle Empty Results**: Always check for zero instance types and provide fallback
3. **Log Quota Issues**: Help users understand why certain types aren't available
4. **Regional Awareness**: Quotas are per-region; multi-region queries may have different results
5. **Preset Validation**: Verify the selected preset has sufficient quota before creating instances

## Summary

This comprehensive testing guide provides:

✅ **Updated Authentication**: Proper Nebius service account credentials (replacing GCP-specific format)

✅ **Complete Test Suite**: Unit tests, integration tests, and smoke tests

✅ **Test Implementation**:
- `client_test.go` - Unit tests for client and credential functionality
- `instance_test.go` - Unit tests for instance operations
- `integration_test.go` - Real API integration testing including instance type enumeration
- `smoke_test.go` - End-to-end instance lifecycle validation

✅ **Testing Guidelines**: Comprehensive execution strategies for development, CI/CD, and production

✅ **Production Readiness**: Detailed checklists and validation procedures

✅ **Instance Type Enumeration**: Quota-aware discovery with elastic disk support

The test suite accommodates the current implementation and provides comprehensive validation of quota-based filtering, preset enumeration, and elastic disk support.