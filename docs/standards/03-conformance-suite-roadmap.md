# Conformance Suite Roadmap

**Version:** 0.1
**Date:** 2026-01-21
**License:** Apache 2.0

---

## Purpose

Provide a phased, public roadmap to evolve the Perspectra Fold conformance suite into an industry-standard interoperability and assurance program aligned to ETSI/ITU-T/ISO guidance.

---

## Guiding Principles

- **Vendor-neutral**: Tests must be implementable without proprietary dependencies.
- **Reproducible**: Deterministic artifacts and stable versions for audit-grade results.
- **Security-first**: Controls cover cryptography, key lifecycle, and operational integrity.
- **Operational realism**: Performance and resilience tests reflect real-world networks.
- **Global readiness**: Evidence maps to regional compliance obligations.

---

## Roadmap Phases

| Phase | Horizon | Focus | Public Outputs |
|-------|---------|-------|----------------|
| Phase 1 | 2026-2030 | L0 API compliance and baseline interop | L0 suite, API schemas, test vectors, public reports |
| Phase 2 | 2031-2040 | L1 security controls and plugfests | L1 suite, audit evidence pack, interop events |
| Phase 3 | 2041-2055 | L2 performance and resilience | L2 suite, performance baselines, failure drills |
| Phase 4 | 2056-2075 | Global registry and continuous alignment | Public registry, long-horizon certification profiles |

---

## Compliance Alignment (Public)

This roadmap aligns to the milestone timeline in [Standards Overview](01-standards-overview.md) and is intended for planning, not as a certification guarantee.
PFOLD-CP = MDQC Compliance Program milestone.

| Compliance Milestone | Conformance Phase | Evidence Focus |
|----------------------|-------------------|----------------|
| PFOLD-CP1: ISO + ETSI baseline | Phase 1 (L0) | API schemas, test vectors, public reports, ZTA/ZTE baseline |
| PFOLD-CP2: ISO + ETSI uplift | Phase 2 (L1) | Security controls, audit evidence pack, policy validation |
| PFOLD-CP3: Crypto validation | Phase 3 (L2) | Performance baselines, resilience drills, crypto boundary readiness |
| PFOLD-CP4: Operational audits | Phase 3/4 | Continuous monitoring reports, interop reports, conformance registry |

---

## Workstreams

1. **API Compliance**: ETSI-aligned interface behavior, schema validation, and error handling.
2. **Protocol Behavior**: Sifting, error correction, privacy amplification, and key lifecycle.
3. **Cryptographic Controls**: Hybrid key derivation, algorithm agility, and fallback policies.
4. **Performance and Resilience**: Key rate, QBER tolerance, latency, and failover timing.
5. **Security and Operations**: Audit logging, access controls, and configuration integrity.
6. **Tooling and Automation**: CI-friendly harnesses, reporters, and reproducible environments.

---

## Key Artifacts

- **Test vectors** for API and protocol conformance.
- **Golden traces** for interoperability verification.
- **Simulator harness** for channel condition replay.
- **Report templates** aligned to ISO 27001, SOC 2, and regional controls.
- **Public registry** of conformance results with versioned profiles.
- **ZTA/ZTE evidence pack**: trust boundary diagram, policy snapshots, identity inventory, and segmentation rules.

---

## Governance and Releases

- **Versioned profiles** with clear deprecation windows.
- **Annual refresh** aligned to standards updates.
- **Interop events** to validate multi-vendor behavior.

---

## Dependencies

- Reference implementations for API and protocol behaviors.
- Published telemetry schemas and logging formats.
- Access to controlled lab and field test environments.
- Standards tracking via [Standards Overview](01-standards-overview.md).
