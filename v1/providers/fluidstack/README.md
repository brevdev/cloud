# FluidStack Provider

FluidStack is an AI cloud platform designed for high-stakes AI workloads, offering bare-metal and virtualized instances with comprehensive security and compliance features.

## Provider Information

- **Provider Name**: `fluidstack`
- **API Documentation**: https://docs.fluidstack.io/api/infrastructure/
- **Base URL**: `https://api.fluidstack.io/v1alpha1`
- **Authentication**: Bearer token (API key)

## Supported Features

### ✅ Instance Management
- **Create Instance**: `POST /instances` - Create new instances with project scoping
- **Get Instance**: `GET /instances/{id}` - Retrieve instance details
- **List Instances**: `GET /instances` - List all instances in project
- **Terminate Instance**: `DELETE /instances/{id}` - Delete instances
- **Start Instance**: `POST /instances/{id}/start` - Start stopped instances
- **Stop Instance**: `POST /instances/{id}/stop` - Stop running instances

### ✅ Instance Types
- **List Instance Types**: `GET /instance-types` - Get available instance configurations
- **GPU Support**: NVIDIA GPU instances for AI/ML workloads
- **Bare Metal & Virtualized**: Both deployment options available

### ✅ Project Management
- **Create Project**: `POST /projects` - Create isolated project environments
- **List Projects**: `GET /projects` - List all projects
- **Get Project**: `GET /projects/{id}` - Get project details
- **Delete Project**: `DELETE /projects/{id}` - Remove projects
- **Project Scoping**: All resources are scoped to projects via `X-PROJECT-ID` header

### ✅ Security Features
- **Disk Encryption**: Hardware-level self-encrypting drives (SEDs) for data at rest
- **Transit Encryption**: SSL/TLS encryption for all network traffic
- **Network Isolation**: Project-level isolation using VXLAN and eBPF
- **Compliance**: HIPAA, GDPR, ISO27001, SOC 2 TYPE I certified

### ✅ Additional Features
- **Tagging**: Support for resource tagging and organization
- **Filesystem Management**: Block and file storage management
- **Kubernetes Clusters**: Managed Kubernetes cluster support
- **Slurm Clusters**: Managed Slurm batch orchestration

## Unsupported Features

### ❌ Firewall Rules
- **Individual Instance Firewalls**: No API endpoints for per-instance firewall rules
- **Security Groups**: Uses project-level isolation instead of security groups
- **Network ACLs**: No granular network access control lists

### ❌ Storage Operations
- **Volume Resizing**: No API support for resizing instance volumes
- **Snapshot Management**: Volume snapshot operations not available
- **Volume Attachment**: Dynamic volume attach/detach not supported

### ❌ Advanced Networking
- **VPC Management**: No virtual private cloud configuration
- **Load Balancers**: No managed load balancer services
- **Custom Networks**: Limited to project-level network isolation

## Implementation Notes

### Authentication
```go
client := NewFluidStackClient("your-api-key")
```

### Project Scoping
All instance operations require a project context:
```go
// All API calls include X-PROJECT-ID header
// Projects provide isolation boundary for resources
```

### Security Model
- **Project-Level Isolation**: Resources are isolated at the project level
- **Default Network Behavior**: Instances can egress to internet, communicate within project
- **Encryption**: Automatic encryption at rest and in transit

## API Capabilities

FluidStack provides a comprehensive Infrastructure API with the following endpoint categories:

- **Projects**: Full CRUD operations for project management
- **Instances**: Complete instance lifecycle management
- **Instance Types**: Hardware configuration discovery
- **Filesystems**: Storage management operations
- **Kubernetes**: Managed Kubernetes cluster operations
- **Slurm**: Managed Slurm cluster operations
- **Capacity**: Real-time capacity checking

## Compliance & Security

FluidStack meets enterprise security requirements with:
- Hardware-level disk encryption using self-encrypting drives
- SSL/TLS encryption for all network traffic
- Tier-3 data centers with biometric access controls
- Multiple compliance certifications (HIPAA, GDPR, ISO27001, SOC 2)
- Project-level network isolation using VXLAN and eBPF

For detailed security information, see [SECURITY.md](./SECURITY.md).
