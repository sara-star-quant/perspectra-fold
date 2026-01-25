# Public Development Plan (GitHub $0) with Paid Alternatives

**Version:** 0.1
**Date:** 2026-01-21
**License:** Apache 2.0

---

## Scope and Assumptions

- Repository is public and hosted on GitHub.
- Documentation is the primary public deliverable.
- Heavy compute (GPU, long simulations) stays on local or lab hardware.
- CI uses CPU-only checks and lightweight validation.
- Long-horizon vision is captured in [Vision 2026-2075](03-vision-2026-2075.md).

---

## $0 GitHub Stack (Public Repo)

**Core**
- GitHub repository with branch protection and CODEOWNERS.
- Issues, Projects, and Discussions for planning and public feedback.
- Releases and tags for versioned deliverables.

**CI and Automation**
- GitHub Actions for lint, tests, markdown checks, and link validation.
- CodeQL code scanning (free for public repos).
- Dependabot for dependency updates.
- Secret scanning (free for public repos).

**Docs**
- README with linked markdown files (no site required).
- Optional GitHub Pages for a lightweight landing page.

**Storage**
- Avoid large binaries and datasets in git.
- Keep artifacts local or in external research storage.

---

## Roadmap Phase to Instruments (Free)

| Phase | Instruments (Free) | Outputs |
|------|---------------------|---------|
| Phase 1: Foundation | GitHub repo, Actions, ADRs, OpenAPI/protobuf | CI/CD baseline, interface definitions |
| Phase 2: Core Dev | Python + Rust locally, gRPC compute service, liboqs | Dimensionality reduction, PQC primitives |
| Phase 3: Quantum Layer | gRPC/protobuf, Go control plane, LDPC libraries | ETSI QKD API, sifting + EC modules |
| Phase 4: Integration | Docker Compose, kind/minikube | Hybrid integration prototypes |
| Phase 5: Deployment | Local lab, GitHub Releases | Pilot documentation and benchmarks |

---

## Conformance Program Alignment

Use the [Conformance Suite Roadmap](../standards/03-conformance-suite-roadmap.md) to plan standards-facing deliverables. Detailed scope lives in [Conformance Suite](../standards/02-conformance-suite.md).

| Public Plan Phase | Conformance Milestone | Evidence Targets |
|-------------------|-----------------------|------------------|
| Phase 1: Foundation | Conformance Phase 1 (L0) | API schemas, test vectors, L0 report template |
| Phase 2: Core Dev | Conformance Phase 1 (L0) | Protocol harness baseline, golden traces |
| Phase 3: Quantum Layer | Conformance Phase 2 (L1) | Security controls, audit evidence pack |
| Phase 4: Integration | Conformance Phase 2 (L1) | Interop report format, plugfest readiness |
| Phase 5: Deployment | Conformance Phase 3 (L2) | Performance baselines, resilience drills |

Note: Conformance Phase 4 (registry and long-horizon profiles) is tracked in [Vision 2026-2075](03-vision-2026-2075.md).

---

## Evidence Checklist by Phase (Public)

**Phase 1: Foundation**
- ADRs for language split and service boundaries.
- API schema baseline (OpenAPI/protobuf).
- Initial test vectors for core compute and control endpoints.
- Zero Trust baseline policy and trust boundary diagram.

**Phase 2: Core Dev**
- Golden traces for core workflows.
- API conformance harness running in CI.
- Metrics schema and sample datasets.

**Phase 3: Quantum Layer**
- Security control mappings (FIPS 140-3 and ISO 27001).
- Hybrid key derivation evidence and negative tests.
- Incident response and audit log procedures.

**Phase 4: Integration**
- Interop report template and sample report.
- Failure drills and recovery runbooks.
- Load/latency baselines for routing and key delivery.

**Phase 5: Deployment**
- L2 conformance report and performance baselines.
- Regional compliance matrix (targeted).
- Pilot readiness checklist and risk register.

---

## Optional Paid Alternatives

**GitHub and DevOps**
- GitHub Team or Enterprise for SSO, audit logs, and tighter policy control.
- Larger or self-hosted runners for GPU or long-running jobs.

**Cloud and Infrastructure**
- AWS/GCP/Azure for VPC-based pilots and managed services.
- Managed KMS/HSM (AWS CloudHSM, Azure Dedicated HSM, GCP Cloud HSM).
- Container registries with higher storage limits.

**Security and Compliance**
- Centralized logging (Datadog, Splunk).
- Compliance automation (Drata, Vanta) for ISO 27001 or SOC 2.
- Pen testing engagements and formal third-party audits.

---

## Deployment Recommendations by Stage

**Documentation and Research**
- GitHub only, no runtime deployments.

**Development and Simulation**
- Local or lab machines with GPU.
- Optional self-hosted runners for reproducible CI on lab hardware.

**Pilot**
- On-prem lab with secure VLANs and hardware QKD devices.
- If cloud is needed, isolate in a single-region VPC with VPN.

**Production**
- On-prem or hybrid, with HSM-backed key storage.
- Centralized logging, change control, and formal incident response.

---

## Digital Transformation Considerations

- Establish a cross-functional governance group (security, networking, research, ops).
- Integrate with existing OSS/BSS and network monitoring stacks where applicable.
- Define data governance and ownership for telemetry and audit artifacts.
- Plan training and change management for operators and partner teams.
- Track transformation metrics (availability, security posture, onboarding speed).

---

## Regional Compliance Guidance (High Level)

**Target sectors (abstract)**
- Telecom and network infrastructure
- Finance (requires DSS-PII compliance)
- Government and defense
- Healthcare and life sciences
- R&D in technology

| Region | Primary Baselines | Notes |
|--------|-------------------|-------|
| US | NIST FIPS 140-3, NIST SP 800-53, SOC 2 | FedRAMP if deploying on US Gov cloud |
| EU | GDPR, NIS2, ETSI QKD, ISO 27001 | Data residency and DPA requirements |
| UK | UK GDPR, NIS Regulations, NCSC guidance | Align with UK Gov security model if public sector |
| Canada | PIPEDA, ITSG-33 (gov) | Provincial privacy adds constraints |
| APAC | PDPA (SG), APPI (JP), Privacy Act (AU) | Map to local regulators for sector-specific rules |

**Note:** This is a planning guide, not legal advice. Use a compliance matrix per target market. If the target sector includes finance, require DSS-PII controls in addition to regional baselines.

---

## Export Control and Crypto Notes

- Quantum hardware and strong crypto may be dual-use.
- Confirm export classifications before shipping hardware or binaries.
- Public-source crypto code is generally eligible for license exceptions.

---

## Immediate Next Actions

- Decide on data classification and target regions.
- Add ADRs for toolchain and CI policy.
- Create a minimal GitHub Actions pipeline.
- Add a SECURITY.md and disclosure policy.
