# Massed Compute Security Notes

Massed Compute authenticates API requests with bearer tokens over HTTPS and supports SSH public-key injection through account SSH-key resources.

The published API does not expose firewall or security-group operations, nor does it document the default inbound policy or storage encryption guarantees. This adapter therefore does not advertise firewall capability and rejects non-empty firewall rules instead of silently ignoring them. These provider behaviors must be verified before production use against the repository's [security requirements](../../../docs/SECURITY.md).
