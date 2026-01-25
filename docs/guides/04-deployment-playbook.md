# Deployment Playbook: Extreme Environments

**Version:** 0.1
**Date:** 2026-01-21
**License:** Apache 2.0

---

## Purpose

Define site-selection criteria, operational checklists, and resilience practices for deploying Perspectra Fold infrastructure in harsh or remote environments.

---

## Site Selection Criteria

- **Thermal envelope:** Ambient ranges, peak loads, and cooling availability.
- **Power stability:** Grid quality, backup duration, and surge protection.
- **Connectivity:** Fiber presence, line-of-sight options, and redundancy.
- **Physical security:** Controlled access, tamper response, and monitoring.
- **Regulatory constraints:** Data residency, export controls, and permits.

---

## Environment Profiles and Strategies

**Hot regions**
- High-efficiency cooling plants or containerized cooling nodes.
- Dust and humidity filtration; sealed optics.

**Polar and extra cold**
- Heater budgets, cold-start procedures, and condensation control.
- Materials qualified for thermal cycling.

**Marine and subsea**
- Pressure-rated housings and corrosion-resistant materials.
- Long maintenance cycles with remote diagnostics.

**Airborne and high altitude**
- Vibration isolation, pressure variance tolerance, and auto-alignment.
- Thermal swings handled by active regulation.

---

## Operational Checklists

**Pre-Deploy**
- Validate site thermal and power profiles.
- Establish physical security controls.
- Stage spares and calibration tools.
- Validate Zero Trust Environment controls and identity enforcement (see [Zero Trust Environment](../architecture/02-zero-trust-environment.md)).

**Install and Commission**
- Verify optical alignment and detector calibration.
- Run baseline QBER and key rate benchmarks.
- Record initial conformance traces.

**Operations**
- Continuous monitoring of QBER, SKR, and drift metrics.
- Scheduled recalibration and alignment checks.
- Automated failover to classical or PQC-only mode.

**Incident Response**
- Trigger thresholds for thermal excursions and link loss.
- Rapid isolation and rollback procedures.
- Audit log capture and post-incident review.

**Maintenance**
- Periodic cleaning and optical inspection.
- Firmware and security patch cadence.
- Spare parts lifecycle planning.

---

## Resilience and Continuity

- **Offline-first control plane** for remote sites with delayed sync.
- **Multi-path routing** across terrestrial and free-space links.
- **Local key caching** with strict rotation policies.
- Routing policies: see [Routing Policy Matrix](#routing-policy-matrix).

---

## Routing Policy Matrix

| Policy | Primary Goal | Thresholds | Weighting (QBER/SKR/Operational) | Notes |
|--------|---------------|------------|----------------------------------|-------|
| Security-first | Minimize errors | QBER <= 0.08, SKR >= 5 | 0.6 / 0.3 / 0.1 | Use for regulated data |
| Balanced | Mixed objectives | QBER <= 0.15, SKR >= 3 | 0.5 / 0.3 / 0.2 | Default policy |
| Latency-first | Minimize latency | QBER <= 0.20, SKR >= 2 | 0.3 / 0.2 / 0.5 | Use for control traffic |
| Availability-first | Maximize uptime | QBER <= 0.24, SKR >= 1 | 0.4 / 0.4 / 0.2 | Prefer larger route sets |

Operational weighting covers loss, latency, and jitter penalties when available.

**Control plane mapping**
- Set `MDQC_ROUTE_POLICY` to `security-first`, `balanced`, `latency-first`, or `availability-first`.
- Override thresholds and weights via `MDQC_ROUTE_*` environment variables if needed.

---

## Evidence and Reporting

- Maintain deployment logs, calibration records, and conformance results.
- Capture environmental telemetry with retention policies.
- Align reports with regional compliance baselines.

---

## Sustainability and Community Impact

- Optimize cooling and power usage for local grid constraints.
- Prefer reusable site infrastructure and modular upgrades.
- Engage local operators and regulators early for site approvals.
- Plan for end-of-life disposal and hardware recycling.
