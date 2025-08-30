# Lambda Labs Provider

This directory contains the Lambda Labs provider implementation for the compute package.

## Overview

The Lambda Labs provider implements the CloudClient interface defined in `pkg/v1` to provide access to Lambda Labs cloud infrastructure. This implementation is based on the official Lambda Labs API documentation at https://cloud.lambda.ai/api/v1/openapi.json.

## Supported Features

Based on the Lambda Labs API documentation, the following features are **SUPPORTED**:

### Instance Management
- ✅ **Create Instance**: `POST /api/v1/instance-operations/launch`
- ✅ **Get Instance**: `GET /api/v1/instances/{id}`
- ✅ **List Instances**: `GET /api/v1/instances`
- ✅ **Terminate Instance**: `POST /api/v1/instance-operations/terminate`
- ✅ **Reboot Instance**: `POST /api/v1/instance-operations/restart`

### Instance Types
- ✅ **Get Instance Types**: `GET /api/v1/instance-types`

### Firewall Management
- ✅ **Create Firewall Ruleset**: `POST /api/v1/firewall-rulesets`
- ✅ **Get Firewall Rulesets**: `GET /api/v1/firewall-rulesets`
- ✅ **Update Firewall Ruleset**: `PATCH /api/v1/firewall-rulesets/{id}`
- ✅ **Delete Firewall Ruleset**: `DELETE /api/v1/firewall-rulesets/{id}`
- ✅ **Global Firewall Rules**: `GET /api/v1/firewall-rulesets/global`

## Unsupported Features

The following features are **NOT SUPPORTED** (no API endpoints found):

### Instance Operations
- ❌ **Stop/Start Instance**: No stop or start endpoints found
- ❌ **Change Instance Type**: No instance type modification endpoint
- ❌ **Instance Tags**: No tagging endpoints found

### Volume Management
- ❌ **Resize Instance Volume**: No volume resizing endpoints found

### Machine Images
- ❌ **Get Images**: No image listing endpoints found

### Quota Management
- ❌ **Get Instance Type Quotas**: No quota endpoints found

### Location Management
- ❌ **Get Locations**: No location listing endpoints found

## Implementation Approach

This implementation uses the `NotImplCloudClient` pattern for unsupported features:
- Supported features have TODO implementations with API endpoint references
- Unsupported features return `ErrNotImplemented` (handled by embedded NotImplCloudClient)
- Full CloudClient interface compliance is maintained

## Lambda Labs API

The provider integrates with the Lambda Labs Cloud API:
- Base URL: `https://cloud.lambda.ai/api/v1`
- Authentication: API key based (Bearer token)
- Documentation: https://cloud.lambda.ai/api/v1/openapi.json

## Key Features

Lambda Labs is known for:
- GPU instances (A10, A100, H100, etc.)
- Global API (not location-specific)
- Simple pricing model
- Fast instance provisioning
- Firewall rulesets for network security

## TODO

- [ ] Implement actual API integration for supported features
- [ ] Add comprehensive error handling
- [ ] Add logging and monitoring
- [ ] Add comprehensive testing 