# Zero Trust Environment (ZTE) and Zero Trust Architecture (ZTA)

**Document ID:** PFOLD-ARCH-002
**Version:** 0.1
**Date:** 2026-01-21
**Classification:** Public
**License:** Apache 2.0

---

## Purpose

Define the Zero Trust posture for Perspectra Fold deployments and explain how Zero Trust Architecture (ZTA) and Zero Trust Environment (ZTE) apply to research, pilots, and production systems.

---

## Scope and Definitions

- **ZTA**: The system-level model for verifying identity, device posture, and policy before access (NIST SP 800-207).
- **ZTE**: The operational environment where Perspectra Fold runs, including sites, labs, field nodes, and remote stations.
- **Zero Trust Encryption** and **Zero Trust Edge** can be applied as subdomains when data-plane encryption or edge enforcement needs explicit focus.

---

## Zero Trust Edge and Zero Trust Encryption (Public Scope)

- **Zero Trust Edge**: Enforce policy at remote sites and field nodes with local decision points.
- **Zero Trust Encryption**: Tie key usage to identity, policy, and telemetry thresholds.
- Keep control-plane authorization independent from the underlying transport or hardware vendor.

---

## Principles

1. **Assume breach**: Treat every network and site as untrusted by default.
2. **Verify explicitly**: Authenticate every request with service identity and device posture.
3. **Least privilege**: Grant only the minimal scope required per request.
4. **Continuous evaluation**: Re-evaluate access based on telemetry and policy changes.

---

## Architecture Mapping

- **Identity and access**: Per-service identities for control plane and compute services, with short-lived credentials.
- **Device and workload**: Signed builds, runtime attestation where feasible, and immutable deployment images.
- **Network segmentation**: Deny-by-default policies, micro-segmentation by plane (control, compute, telemetry).
- **Data protection**: Encrypt in transit and at rest; derive session keys from QKD + PQC + classical sources.
- **Policy control**: Central policy definitions enforced in the control plane; local enforcement for disconnected sites.
- **Observability**: Continuous telemetry for QBER, SKR, link health, and policy compliance.

---

## Reference Architecture (High Level)

```
                +----------------------+
                |   Policy Control     |
                | (identity + policy)  |
                +----------+-----------+
                           |
                           v
    +----------------------+----------------------+
    |              Control Plane                  |
    |  authz, policy eval, routing, telemetry     |
    +-----------+--------------+------------------+
                |              |
                v              v
     +----------+---+   +------+-----------+
     | Compute svc |   | Telemetry store  |
     | (gRPC)      |   | (read-only ingest)|
     +------+------+   +------+-----------+
            |                 |
            v                 v
        +---+-----------------+---+
        |   Data Plane / Links    |
        | QKD + PQC + classical   |
        +-------------------------+
```

This diagram is intentionally high-level and omits enforcement details, key material handling, and site-specific controls.

---

## ZTE Focus Areas

- **Site isolation**: Separate lab, staging, and production environments with independent credentials.
- **Environmental gating**: Condition-based route selection and key usage based on telemetry quality.
- **Offline-first operations**: Allow local policy enforcement with delayed sync for remote sites.
- **Physical controls**: Tamper detection, secure enclosures, and controlled access for high-assurance sites.

---

## Long-Haul Experiment Compatibility (Public Reference)

Reference case: China-South Africa microsatellite QKD link (12,900 km, 2024).

MDQC applicability and compatibility targets:
- Link-agnostic control plane with standardized telemetry and policy gating.
- Support for high-loss, high-latency links via adaptive routing and local key caching.
- ETSI-aligned key delivery interface for interoperability with emerging hardware.
- Modular hardware integration to support new detectors, sources, and ground stations.

---

## Public vs Private Guidance

- **Public**: Principles, interface expectations, and high-level compatibility guidance.
- **Private**: Enforcement policies, identity and segmentation maps, hardware integration playbooks, and site-specific thresholds.

---

## References

- NIST SP 800-207: Zero Trust Architecture
- ETSI GS QKD 014/015/018
- ITU-T Y.3800/Y.3801/Y.3802
