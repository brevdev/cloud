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
// Actual Nebius quota naming patterns (discovered from API)
"compute.instance.gpu.h100:eu-north1"          // H100 GPUs in eu-north1
"compute.instance.gpu.h200:eu-north1"          // H200 GPUs in eu-north1
"compute.instance.gpu.l40s:eu-north1"          // L40S GPUs in eu-north1
"compute.instance.gpu.b200:us-central1"        // B200 GPUs in us-central1
"compute.instance.non-gpu.vcpu:eu-north1"      // vCPU quota for CPU instances
"compute.instance.non-gpu.memory:eu-north1"    // Memory quota for CPU instances
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

## Practical Testing Commands for Implementation Validation

### Prerequisites

Set up your testing environment with Nebius credentials:

```bash
# Export credentials
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/your/service-account.json'
export NEBIUS_TENANT_ID='tenant-e00xxx'  # Your tenant ID
export NEBIUS_LOCATION='eu-north1'       # Target region
```

### Quick Commands for Testing Instance Types (Quota-Aware)

#### Command 1: Enumerate Instance Types with Quota Information

```bash
# Test GetInstanceTypes with quota filtering
cd /home/jmorgan/VS/brev-cloud-sdk/v1/providers/nebius

# Run integration test that enumerates instance types
go test -v -run TestIntegration_GetInstanceTypes

# Expected output:
# === RUN   TestIntegration_GetInstanceTypes
# === RUN   TestIntegration_GetInstanceTypes/Get_instance_types_with_quota_filtering
#     Found 12 instance types with available quota
#     Instance Type: computeplatform-e00abc-1gpu (L40S) - Location: eu-north1, Available: true
#       Storage: network-ssd, Min: 50 GB, Max: 2560 GB, Elastic: true
#       GPU: NVIDIA L40S (Type: L40S), Count: 1, Manufacturer: NVIDIA
# === RUN   TestIntegration_GetInstanceTypes/Verify_quota_filtering
#     All returned instance types have available quota
# === RUN   TestIntegration_GetInstanceTypes/Verify_preset_enumeration
#     Preset distribution: L40S (4), H100 (4), H200 (2), CPU (3)
```

#### Command 2: Dump Instance Types to JSON (Aggregated with Real Pricing)

This command aggregates instance types across regions with **real pricing from Nebius Billing API**, matching the LaunchPad API format:

```bash
# Set tenant-level credentials (no project ID needed!)
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='tenant-e00xxx'

# Run WITH real pricing (takes ~60 seconds, queries Nebius Billing Calculator API)
cd /home/jmorgan/VS/brev-cloud-sdk/v1/providers/nebius
FETCH_PRICING=true go run ./cmd/dump_instance_types/main.go > complete_catalog.json

# Or run WITHOUT pricing (instant, pricing = 0)
go run ./cmd/dump_instance_types/main.go > instance_types.json

# View GPU types with pricing
cat complete_catalog.json | jq '.[] | select(.gpu != null) | {preset: .preset, regions, gpu: {count: .gpu.count, family: .gpu.family}, price}'

# Show L40S pricing comparison
cat complete_catalog.json | jq -r '.[] | select(.gpu.family == "l40s") | "\(.preset): $\(.price.on_demand_per_hour)/hr ($\(.price.estimated_monthly | floor)/mo)"'

# Expected output:
# 1gpu-16vcpu-96gb: $1.8172/hr ($1326/mo)
# 2gpu-64vcpu-384gb: $4.5688/hr ($3335/mo)
# 4gpu-128vcpu-768gb: $9.1376/hr ($6670/mo)

# Show H200 with cross-region capacity and pricing
cat complete_catalog.json | jq '.[] | select(.gpu.family == "h200")'

# Expected: H200 available in 3 regions with real pricing ($3.50-$28/hr)
```

**Example Output** (Aggregated Format with Semantic IDs):
```json
{
  "id": "gpu-l40s-d-4gpu-128vcpu-768gb",
  "nebius_platform_id": "computeplatform-e00q7xea367y069e81",
  "cloud": "nebius",
  "platform": "gpu-l40s-d",
  "preset": "4gpu-128vcpu-768gb",
  "capacity": {
    "eu-north1": 1
  },
  "regions": ["eu-north1"],
  "cpu": 128,
  "memory_gb": 768,
  "gpu": {
    "count": 4,
    "family": "l40s",
    "model": "NVIDIA L40S",
    "manufacturer": "NVIDIA"
  },
  "storage": [
    {
      "type": "network-ssd",
      "size_min_gb": 50,
      "size_max_gb": 2560,
      "is_elastic": true
    }
  ],
  "system_arch": "amd64",
  "price": {
    "currency": "USD",
    "on_demand_per_hour": 9.1376,      ← Real Nebius pricing!
    "estimated_monthly": 6670.448       ← With FETCH_PRICING=true
  }
}
```

**Key Features**:
- ✅ One entry per preset configuration (not per region)
- ✅ `capacity` map shows availability across all regions
- ✅ `regions` list shows where quota exists
- ✅ **Real pricing from Nebius Billing Calculator API** (with FETCH_PRICING=true)
- ✅ Decimal precision for accurate cost estimates
- ✅ Matches LaunchPad API format for easy comparison

**Note**: The SDK's `GetInstanceTypes()` returns one entry **per region** (this is intentional and matches LaunchPad SDK behavior). This dump utility **aggregates them** for easier visualization.
```

#### Command 3: View Regional Capacity Distribution

```bash
# Show which regions have which GPU types available
cat instance_types_aggregated.json | jq -r '.[] | select(.gpu != null) | "\(.gpu.family) (\(.gpu.count)x): \(.regions | join(", "))"' | sort | uniq

# Expected output:
# h100 (1x): eu-north1
# h100 (8x): eu-north1
# h200 (1x): eu-north1, eu-west1, us-central1
# h200 (8x): eu-north1, eu-west1, us-central1
# l40s (1x): eu-north1
# l40s (2x): eu-north1
# l40s (4x): eu-north1

# Count total instance types by GPU family
cat instance_types_aggregated.json | jq -r '.[] | select(.gpu != null) | .gpu.family' | sort | uniq -c

# Show capacity breakdown
cat instance_types_aggregated.json | jq '.[] | select(.gpu != null) | {family: .gpu.family, count: .gpu.count, capacity, regions}'
```

### Testing Commands for GetImages

#### Command 4: Enumerate Available Images

```bash
# Test GetImages with architecture filtering
cd /home/jmorgan/VS/brev-cloud-sdk/v1/providers/nebius

# Create test script for images
cat > test_images.go << 'EOF'
package main

import (
    "context"
    "fmt"
    "os"
    nebius "github.com/brevdev/cloud/v1/providers/nebius"
    v1 "github.com/brevdev/cloud/v1"
)

func main() {
    ctx := context.Background()
    
    saJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
    tenantID := os.Getenv("NEBIUS_TENANT_ID")
    location := os.Getenv("NEBIUS_LOCATION")
    
    if saJSON == "" || tenantID == "" || location == "" {
        fmt.Fprintln(os.Stderr, "Error: Set required environment variables")
        os.Exit(1)
    }
    
    saKey, _ := os.ReadFile(saJSON)
    credential := nebius.NewNebiusCredentialWithOrg("test", string(saKey), tenantID, "")
    client, err := credential.MakeClient(ctx, location)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
        os.Exit(1)
    }
    
    // Get x86_64 images (default for GPU instances)
    fmt.Println("=== x86_64 Images ===")
    x86Images, err := client.GetImages(ctx, v1.GetImageArgs{
        Architectures: []string{"x86_64"},
    })
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting x86 images: %v\n", err)
    } else {
        for _, img := range x86Images {
            fmt.Printf("  - %s (%s) - Arch: %s\n", img.Name, img.ID, img.Architecture)
        }
    }
    
    // Get all images
    fmt.Println("\n=== All Available Images ===")
    allImages, err := client.GetImages(ctx, v1.GetImageArgs{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting all images: %v\n", err)
    } else {
        fmt.Printf("Total images available: %d\n", len(allImages))
    }
}
EOF

go run test_images.go
```

### Testing Commands for GetLocations

#### Command 5: Enumerate Available Locations

```bash
# Test GetLocations
cat > test_locations.go << 'EOF'
package main

import (
    "context"
    "fmt"
    "os"
    nebius "github.com/brevdev/cloud/v1/providers/nebius"
    v1 "github.com/brevdev/cloud/v1"
)

func main() {
    ctx := context.Background()
    
    saJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
    tenantID := os.Getenv("NEBIUS_TENANT_ID")
    location := os.Getenv("NEBIUS_LOCATION")
    
    if saJSON == "" || tenantID == "" {
        fmt.Fprintln(os.Stderr, "Error: Set required environment variables")
        os.Exit(1)
    }
    
    saKey, _ := os.ReadFile(saJSON)
    credential := nebius.NewNebiusCredentialWithOrg("test", string(saKey), tenantID, "")
    client, err := credential.MakeClient(ctx, location)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
        os.Exit(1)
    }
    
    locations, err := client.GetLocations(ctx, v1.GetLocationsArgs{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting locations: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("=== Available Nebius Locations ===")
    for _, loc := range locations {
        fmt.Printf("  - %s: %s (Available: %t, Country: %s)\n",
            loc.Name, loc.Description, loc.Available, loc.Country)
    }
}
EOF

go run test_locations.go
```

### Testing Commands for GetCapabilities

#### Command 6: Check Provider Capabilities

```bash
# Test GetCapabilities
cat > test_capabilities.go << 'EOF'
package main

import (
    "context"
    "fmt"
    "os"
    nebius "github.com/brevdev/cloud/v1/providers/nebius"
)

func main() {
    ctx := context.Background()
    
    saJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
    tenantID := os.Getenv("NEBIUS_TENANT_ID")
    location := os.Getenv("NEBIUS_LOCATION")
    
    if saJSON == "" || tenantID == "" {
        fmt.Fprintln(os.Stderr, "Error: Set required environment variables")
        os.Exit(1)
    }
    
    saKey, _ := os.ReadFile(saJSON)
    credential := nebius.NewNebiusCredentialWithOrg("test", string(saKey), tenantID, "")
    client, err := credential.MakeClient(ctx, location)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
        os.Exit(1)
    }
    
    capabilities, err := client.GetCapabilities(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting capabilities: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("=== Nebius Provider Capabilities ===")
    for _, cap := range capabilities {
        fmt.Printf("  ✓ %s\n", cap)
    }
}
EOF

go run test_capabilities.go
```

### Testing Commands for Full Instance Lifecycle

#### Command 7: End-to-End Instance Creation Test

```bash
# Run smoke test to create/verify/terminate an instance
export RUN_SMOKE_TESTS=true
export CLEANUP_RESOURCES=true

cd /home/jmorgan/VS/brev-cloud-sdk/v1/providers/nebius

# Run the smoke test (creates actual cloud resources)
go test -v -run TestSmoke_InstanceLifecycle -timeout=20m

# Expected flow:
# 1. ✅ Authentication and project setup
# 2. ✅ Network infrastructure creation (VPC, subnet)
# 3. ✅ Boot disk creation
# 4. ✅ Instance creation with L40S GPU
# 5. ✅ Instance verification (GetInstance)
# 6. ✅ Instance termination
# 7. ✅ Resource cleanup
```

### Ad-Hoc Testing Commands

#### Command 8: Test Specific Instance Type Creation

```bash
# Test creating an instance with a specific instance type
cat > test_create_instance.go << 'EOF'
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    nebius "github.com/brevdev/cloud/v1/providers/nebius"
    v1 "github.com/brevdev/cloud/v1"
)

func main() {
    ctx := context.Background()
    
    saJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
    tenantID := os.Getenv("NEBIUS_TENANT_ID")
    location := os.Getenv("NEBIUS_LOCATION")
    
    if saJSON == "" || tenantID == "" || location == "" {
        fmt.Fprintln(os.Stderr, "Error: Set required environment variables")
        os.Exit(1)
    }
    
    saKey, _ := os.ReadFile(saJSON)
    credential := nebius.NewNebiusCredentialWithOrg("test-adhoc", string(saKey), tenantID, "")
    client, err := credential.MakeClient(ctx, location)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
        os.Exit(1)
    }
    
    // First, get available instance types
    instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting instance types: %v\n", err)
        os.Exit(1)
    }
    
    if len(instanceTypes) == 0 {
        fmt.Println("No instance types available")
        return
    }
    
    // Use first available instance type
    selectedType := instanceTypes[0]
    fmt.Printf("Selected instance type: %s\n", selectedType.ID)
    
    // Create instance
    testID := fmt.Sprintf("adhoc-test-%d", time.Now().Unix())
    attrs := v1.CreateInstanceAttrs{
        RefID:        testID,
        Name:         testID,
        InstanceType: string(selectedType.ID),
        ImageID:      "ubuntu22.04-cuda12",  // Default image
        DiskSize:     50 * 1024 * 1024 * 1024,  // 50 GB
        Location:     location,
    }
    
    fmt.Printf("Creating instance '%s'...\n", testID)
    instance, err := client.CreateInstance(ctx, attrs)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating instance: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Printf("✅ Instance created successfully!\n")
    fmt.Printf("   ID: %s\n", instance.CloudID)
    fmt.Printf("   Name: %s\n", instance.Name)
    fmt.Printf("   Status: %s\n", instance.Status.LifecycleStatus)
    fmt.Printf("\n⚠️  Remember to terminate this instance manually:\n")
    fmt.Printf("   Instance ID: %s\n", instance.CloudID)
}
EOF

# Run with caution - creates real resources
go run test_create_instance.go
```

#### Command 9: Test Quota Limits Discovery

```bash
# Use the Nebius CLI to check quotas directly
# Install Nebius CLI first if not already installed
curl -sSfL https://storage.googleapis.com/nebius-cli/install.sh | bash

# Authenticate
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='tenant-e00xxx'
nebius init

# List all quota allowances
nebius quotas quota-allowance list \
  --parent-id $NEBIUS_TENANT_ID \
  --format json | jq '.items[] | {name: .metadata.name, region: .spec.region, limit: .spec.limit, usage: .status.usage, state: .status.state}'

# Check specific GPU quota (note the correct format)
nebius quotas quota-allowance list \
  --parent-id $NEBIUS_TENANT_ID \
  --format json | jq '.items[] | select(.metadata.name | contains("instance.gpu"))'

# Expected output:
# {
#   "name": "compute.instance.gpu.l40s",
#   "region": "eu-north1",
#   "limit": 8,
#   "usage": 0,
#   "state": "STATE_ACTIVE"
# }

# Show quota summary by GPU type
nebius quotas quota-allowance list \
  --parent-id $NEBIUS_TENANT_ID \
  --format json | jq -r '.items[] | select(.metadata.name | contains("instance.gpu")) | "\(.metadata.name) in \(.spec.region): \(.spec.limit) total, \(.status.usage) used, \(.spec.limit - .status.usage) available"'
```

#### Command 10: Compare Instance Type Counts Across Providers

```bash
# Quick comparison using the dump utility
echo "=== Provider Instance Type Comparison ==="
echo

echo "Nebius (aggregated by preset):"
cat instance_types_aggregated.json | jq '. | length'
echo "  Unique presets found (see instance_types_aggregated.json for details)"

echo
echo "Nebius (per-region expansion):"
go test -run TestIntegration_GetInstanceTypes -v 2>&1 | grep "Found.*instance types" | head -1

echo
echo "Note: Nebius uses quota-based filtering across multiple regions"
echo "      - Aggregated view: One entry per preset configuration"
echo "      - SDK view: One entry per preset per region (matches LaunchPad pattern)"
```

#### Command 11: Estimate Pricing (Nebius Billing Calculator API)

**Now using REAL Nebius Billing API!** ✅

See: https://github.com/nebius/api/blob/main/nebius/billing/v1alpha1/calculator_service.proto

```bash
# Run the pricing estimator (queries actual Nebius Billing Calculator API)
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='tenant-xxx'
export NEBIUS_PROJECT_ID='project-xxx'  # Your project ID

cd /home/jmorgan/VS/brev-cloud-sdk/v1/providers/nebius
go run ./cmd/estimate_pricing/main.go > pricing_estimates.json

# View L40S GPU pricing
cat pricing_estimates.json | jq -r '.[] | select(.platform_name | contains("l40s")) | "\(.preset_name): $\(.hourly_rate)/hr ($\(.monthly_rate | floor)/mo)"'

# Expected output (actual rates from Nebius):
# 1gpu-16vcpu-96gb: $1.82/hr ($1326/mo)
# 2gpu-64vcpu-384gb: $4.57/hr ($3335/mo)
# 4gpu-128vcpu-768gb: $9.14/hr ($6670/mo)

# View H100/H200 pricing
cat pricing_estimates.json | jq -r '.[] | select(.platform_name | contains("h100") or contains("h200")) | "\(.platform_name) \(.preset_name): $\(.hourly_rate)/hr ($\(.monthly_rate | floor)/mo)"'

# Expected output:
# gpu-h100-sxm 1gpu-16vcpu-200gb: $2.95/hr ($2153/mo)
# gpu-h100-sxm 8gpu-128vcpu-1600gb: $23.6/hr ($17228/mo)
# gpu-h200-sxm 1gpu-16vcpu-200gb: $3.5/hr ($2555/mo)
# gpu-h200-sxm 8gpu-128vcpu-1600gb: $28/hr ($20440/mo)

# Join pricing with instance types
jq -s '
  [.[0][] as $it | .[1][] as $price | 
   if ($it.id | startswith($price.platform_id)) and ($it.preset == $price.preset_name) 
   then $it + {price: {currency: $price.currency, on_demand_per_hour: $price.hourly_rate, estimated_monthly: $price.monthly_rate}} 
   else empty end]
' instance_types_aggregated.json pricing_estimates.json | jq '.[0:3]'
```

**How It Works**:
1. Uses `sdk.Services().Billing().V1Alpha1().Calculator().Estimate()` API
2. Queries pricing for each platform/preset combination
3. Returns hourly, daily, monthly, and annual rates
4. Real pricing data from Nebius billing system

**Note**: Pricing may vary by region and contract type. This shows standard on-demand pricing.
```

### Comprehensive Testing Checklist

Use this checklist to validate the Nebius implementation:

```bash
# Quick way: Use the provided test runner script
./RUN_TESTS.sh

# Or manually:
# 1. Authentication
go test -v -run TestIntegration_ClientCreation

# 2. Instance Types (Quota-Aware)
go test -v -run TestIntegration_GetInstanceTypes

# 3. Images Discovery
go test -v -run TestIntegration_GetImages

# 4. Locations
go test -v -run TestIntegration_GetLocations

# 5. Capabilities
go test -v -run TestIntegration_GetCapabilities

# 6. Full Lifecycle (Creates Real Resources!)
export RUN_SMOKE_TESTS=true
export CLEANUP_RESOURCES=true
go test -v -run TestSmoke_InstanceLifecycle -timeout=20m

# 7. Cleanup Verification
# After smoke tests, verify no orphaned resources remain
nebius compute instance list --parent-id $NEBIUS_PROJECT_ID | grep "smoke-test-"
nebius compute disk list --parent-id $NEBIUS_PROJECT_ID | grep "smoke-test-"
```

### Common Test Issues and Troubleshooting

#### Issue 1: "No GPU quota allocated - only CPU instances available"

**Symptom**: The test passes but shows only CPU instance types, with a warning about no GPU quota.

**Example Output**:
```
Instance type distribution:
  CPU-only: 6
⚠️  No GPU quota allocated - only CPU instances available
   To test GPU instances, request GPU quota from Nebius support
```

**Cause**: Your Nebius tenant doesn't have GPU quota allocated. The quota-aware filtering is **working correctly** - it only returns instance types where you have available quota.

**What's Happening**:
- ✅ The implementation is working as designed
- ✅ Quota-aware filtering is functioning correctly
- ✅ You have CPU quota (cpu-d3, cpu-e2) which is being returned
- ⚠️ You don't have GPU quota (L40S, H100, H200, etc.)

**Solution**:

1. **Request GPU Quota** (for real GPU testing):
```bash
# Check current quotas
nebius quotas quota-allowance list \
  --parent-id $NEBIUS_TENANT_ID \
  --format json | jq '.items[] | select(.metadata.name | contains("gpu"))'

# If empty, contact Nebius support to request:
# - L40S GPU quota (good for testing)
# - H100/H200 GPU quota (production workloads)
```

2. **Or continue with CPU-only testing**:
   The implementation is still fully functional and can be tested with CPU instances.

#### Issue 2: Test Skipped Due to Missing Environment Variables

**Symptom**:
```
Skipping integration test: NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set
```

**Solution**:
```bash
# Set required environment variables
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/your-service-account.json'
export NEBIUS_TENANT_ID='tenant-e00xxx'
export NEBIUS_LOCATION='eu-north1'  # Optional, defaults to eu-north1

# Then run the test
go test -v -run TestIntegration_GetInstanceTypes
```

Or use the provided test runner:
```bash
./RUN_TESTS.sh
```

#### Issue 3: Authentication Failures

**Symptom**: `failed to initialize Nebius SDK` or `invalid service account`

**Solutions**:
```bash
# Verify JSON format
cat $NEBIUS_SERVICE_ACCOUNT_JSON | jq .

# Check required fields exist
jq -r '.subject_credentials.subject, .subject_credentials.private_key' $NEBIUS_SERVICE_ACCOUNT_JSON

# Ensure file permissions are correct
chmod 600 $NEBIUS_SERVICE_ACCOUNT_JSON
```

## Provider Comparison: Nebius vs Lambdalabs vs Shadeform

### Feature Parity Matrix

| Feature | Nebius | Lambdalabs | Shadeform | Notes |
|---------|--------|------------|-----------|-------|
| **Core Instance Operations** |
| CreateInstance | ✅ | ✅ | ✅ | All support basic instance creation |
| GetInstance | ✅ | ✅ | ✅ | All support instance retrieval |
| TerminateInstance | ✅ | ✅ | ✅ | All support termination |
| ListInstances | ⚠️ | ✅ | ✅ | Nebius: pending implementation |
| RebootInstance | ⚠️ | ✅ | ✅ | Nebius: pending implementation |
| StopInstance | ⚠️ | ❌ | ❌ | Nebius: pending, others don't support |
| StartInstance | ⚠️ | ❌ | ❌ | Nebius: pending, others don't support |
| **Resource Discovery** |
| GetInstanceTypes | ✅ | ✅ | ✅ | All support with different strategies |
| GetInstanceTypes (Quota) | ✅ | ❌ | ❌ | Only Nebius has quota-aware filtering |
| GetImages | ✅ | ❌ | ✅ | Lambdalabs has no image API |
| GetLocations | ✅ | ✅ | ✅ | All support location discovery |
| GetCapabilities | ✅ | ✅ | ✅ | All support capability reporting |
| **Advanced Features** |
| Tags/Labels | ✅ | ❌ | ✅ | Nebius and Shadeform support tagging |
| Elastic Volumes | ✅ | ❌ | ❌ | Nebius supports volume resizing |
| Firewall Rules | ⚠️ | ⚠️ | ⚠️ | Limited support across all providers |
| SSH Key Management | ✅ | ✅ | ✅ | All support SSH key injection |
| **Network Management** |
| VPC/Network Creation | ✅ | ❌ | ❌ | Only Nebius manages networks |
| Subnet Management | ✅ | ❌ | ❌ | Only Nebius manages subnets |
| **Authentication** |
| API Key | N/A | ✅ | ✅ | Lambdalabs and Shadeform use API keys |
| Service Account | ✅ | N/A | N/A | Nebius uses service account JSON |
| OAuth | ❌ | ❌ | ❌ | None support OAuth |

### Implementation Comparison

#### Instance Type Discovery

**Nebius** (Quota-Aware + Pricing API):
```go
// Queries actual quota from Nebius Quotas API
// Filters platforms by active quota state
// Only returns instance types with available capacity
// Supports elastic disk configuration (50GB-2560GB)
// Real pricing via Billing Calculator API
instanceTypes, _ := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
// Returns: L40S, H100, H200, etc. (only with quota)
// Pricing: go run ./cmd/estimate_pricing/main.go (real Nebius rates)
```

**Lambdalabs** (Capacity-Based):
```go
// Queries instance types from Lambda API
// Checks RegionsWithCapacityAvailable per type
// Returns all types with per-region availability
instanceTypes, _ := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
// Returns: A10, A100, H100, etc. (all types, marked available/unavailable)
```

**Shadeform** (Configuration-Filtered):
```go
// Queries all shade instance types
// Applies configuration-based allow/deny list
// Can filter by cloud provider and instance type
client.WithConfiguration(Configuration{
    AllowedInstanceTypes: map[openapi.Cloud]map[string]bool{
        openapi.HYPERSTACK: {"A4000": true},
    },
})
instanceTypes, _ := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
// Returns: Only configured types (e.g., hyperstack_A4000)
```

#### Authentication Patterns

**Nebius**:
```go
// Service account JSON with RSA key pairs
credential := NewNebiusCredential(refID, serviceAccountJSON, tenantID)
client, _ := credential.MakeClient(ctx, "eu-north1")
// Creates per-user projects automatically
```

**Lambdalabs**:
```go
// Simple API key authentication
credential := NewLambdaLabsCredential(refID, apiKey)
client, _ := credential.MakeClient(ctx, "us-west-1")
// Global API, no project management
```

**Shadeform**:
```go
// API key with tag-based resource tracking
credential := NewShadeformCredential(refID, apiKey)
client, _ := credential.MakeClient(ctx, "")
// Uses tags to identify resources
```

### Key Differences

1. **Resource Management Model**:
   - **Nebius**: Hierarchical (Tenant → Project → Resources)
   - **Lambdalabs**: Flat (Account → Instances)
   - **Shadeform**: Tag-based (Account → Tagged Instances)

2. **Quota Management**:
   - **Nebius**: Explicit quota API with state tracking
   - **Lambdalabs**: Implicit capacity via RegionsWithCapacityAvailable
   - **Shadeform**: Configuration-based filtering

3. **Network Infrastructure**:
   - **Nebius**: Full VPC/Subnet management required
   - **Lambdalabs**: Automatic network assignment
   - **Shadeform**: Provider-managed networking

4. **Instance Type Filtering**:
   - **Nebius**: Quota-based (only show what you can use)
   - **Lambdalabs**: Availability-based (show all, mark availability)
   - **Shadeform**: Configuration-based (pre-filter allowed types)

### Feature Gaps Analysis

**Nebius Missing Features (vs others)**:
- ⚠️ ListInstances: Not yet implemented (but easy to add)
- ⚠️ RebootInstance: Not yet implemented (API supports it)

**Lambdalabs Missing Features (vs others)**:
- ❌ GetImages: No API available
- ❌ Stop/Start: No API endpoints
- ❌ Tags: No tagging support
- ❌ GetInstanceTypeQuotas: No quota API

**Shadeform Missing Features (vs others)**:
- ❌ Stop/Start: Not supported by underlying API
- ❌ Elastic Volumes: Fixed disk sizes

### Recommendation for Feature Parity

To achieve full feature parity, Nebius should implement:

1. **High Priority** (Simple to add):
   - ✅ ListInstances - Straightforward SDK call
   - ✅ RebootInstance - SDK supports instance restart

2. **Medium Priority** (Requires testing):
   - ✅ StopInstance/StartInstance - SDK supports, needs validation
   - ✅ UpdateInstanceTags - SDK supports resource labels

3. **Low Priority** (Nice to have):
   - ResizeInstanceVolume - Already structured, needs implementation
   - Firewall Rules - Requires security group integration

All critical features for parity with Lambdalabs and Shadeform are either:
- ✅ Already implemented
- ⚠️ Partially implemented (needs completion)
- 📋 Structured and ready for implementation

## Summary

This comprehensive testing guide provides:

✅ **Updated Authentication**: Proper Nebius service account credentials (replacing GCP-specific format)

✅ **Complete Test Suite**: Unit tests, integration tests, and smoke tests

✅ **Test Implementation**:
- `client_test.go` - Unit tests for client and credential functionality
- `instance_test.go` - Unit tests for instance operations
- `integration_test.go` - Real API integration testing including instance type enumeration
- `smoke_test.go` - End-to-end instance lifecycle validation

✅ **Practical Testing Commands**: Ad-hoc commands for enumerating instance types, images, locations, and testing full lifecycle

✅ **Provider Comparison**: Comprehensive analysis of Nebius vs Lambdalabs vs Shadeform

✅ **Feature Parity Assessment**: Clear roadmap for achieving full feature parity

✅ **Testing Guidelines**: Comprehensive execution strategies for development, CI/CD, and production

✅ **Production Readiness**: Detailed checklists and validation procedures

✅ **Instance Type Enumeration**: Quota-aware discovery with elastic disk support

The test suite accommodates the current implementation and provides comprehensive validation of quota-based filtering, preset enumeration, and elastic disk support.