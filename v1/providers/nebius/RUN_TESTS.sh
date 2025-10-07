#!/bin/bash

# Nebius Integration Test Runner
# This script helps you run the integration tests with proper environment setup

set -e

echo "===================================="
echo "Nebius Integration Test Runner"
echo "===================================="
echo

# Check for required environment variables
if [ -z "$NEBIUS_SERVICE_ACCOUNT_JSON" ]; then
    echo "❌ NEBIUS_SERVICE_ACCOUNT_JSON not set"
    echo
    echo "Please set it to the path of your service account JSON file:"
    echo "  export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'"
    exit 1
fi

if [ -z "$NEBIUS_TENANT_ID" ]; then
    echo "❌ NEBIUS_TENANT_ID not set"
    echo
    echo "Please set it to your Nebius tenant ID:"
    echo "  export NEBIUS_TENANT_ID='tenant-e00xxx'"
    exit 1
fi

# Optional: set location (defaults to eu-north1 in tests)
if [ -z "$NEBIUS_LOCATION" ]; then
    export NEBIUS_LOCATION="eu-north1"
    echo "ℹ️  Using default location: $NEBIUS_LOCATION"
fi

echo "✅ Environment configured:"
echo "   Service Account: $NEBIUS_SERVICE_ACCOUNT_JSON"
echo "   Tenant ID: $NEBIUS_TENANT_ID"
echo "   Location: $NEBIUS_LOCATION"
echo

# Check if service account file exists
if [ ! -f "$NEBIUS_SERVICE_ACCOUNT_JSON" ]; then
    echo "❌ Service account file not found: $NEBIUS_SERVICE_ACCOUNT_JSON"
    exit 1
fi

echo "===================================="
echo "Running Integration Tests"
echo "===================================="
echo

# Run the test
go test -v -run TestIntegration_GetInstanceTypes

echo
echo "===================================="
echo "Test run completed!"
echo "===================================="
echo
echo "Note: If you see 'No GPU quota allocated', that's normal."
echo "      Your account only has CPU quota. The quota-aware"
echo "      filtering is working correctly - it only shows"
echo "      instance types where you have available quota."
echo
echo "To get GPU quota, contact Nebius support and request:"
echo "  - L40S GPU quota (for testing)"
echo "  - H100/H200 GPU quota (for production workloads)"



