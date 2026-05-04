# Standards and Certifications Overview

**Version:** 1.0
**Date:** 2026-01-21
**License:** Apache 2.0

---

## Overview

This document provides a comprehensive overview of applicable standards, protocols, and certification requirements for implementing high-dimensional quantum and hybrid quantum-classical communication systems.

---

## How to Use This Standards Guidance

- Use this document to decide which external standards apply.
- Use [Conformance Suite](02-conformance-suite.md) to validate implementation evidence.
- Use [Conformance Suite Roadmap](03-conformance-suite-roadmap.md) to plan phased deliverables.
- For long-horizon mapping, see [Vision Appendix A](../guides/03-vision-2026-2075.md).

---

## 1. Applicable Standards

### 1.1 Quantum Key Distribution Standards

| Standard | Organization | Scope | Status |
|----------|--------------|-------|--------|
| **ETSI GS QKD 004** | ETSI ISG QKD | Application Interface | Published |
| **ETSI GS QKD 008** | ETSI ISG QKD | QKD Module Security | Published |
| **ETSI GS QKD 014** | ETSI ISG QKD | Protocol & Data Format | Published |
| **ETSI GS QKD 015** | ETSI ISG QKD | Control Interface | Published |
| **ETSI GS QKD 018** | ETSI ISG QKD | Orchestration Interface | Published |
| **ITU-T Y.3800** | ITU-T SG13 | QKD Network Framework | Published |
| **ITU-T Y.3801** | ITU-T SG13 | QKD Network Functional Req | Published |
| **ITU-T Y.3802** | ITU-T SG13 | QKD Network Architecture | Published |
| **ISO/IEC 23837** | ISO/IEC JTC1 | Security Requirements | In Development |

### 1.2 Post-Quantum Cryptography Standards

| Standard | Organization | Algorithm | Status |
|----------|--------------|-----------|--------|
| **FIPS 203** | NIST | ML-KEM (Kyber) | Published 2024 |
| **FIPS 204** | NIST | ML-DSA (Dilithium) | Published 2024 |
| **FIPS 205** | NIST | SLH-DSA (SPHINCS+) | Published 2024 |
| **FIPS 206** | NIST | FN-DSA (Falcon) | Expected 2025 |
| **RFC 9180** | IETF | HPKE | Published |
| **RFC 9370** | IETF | Hybrid Key Exchange (draft) | In Progress |

### 1.3 Classical Cryptography Standards

| Standard | Scope | Requirement |
|----------|-------|-------------|
| **FIPS 140-3** | Cryptographic Modules | Level 2+ recommended |
| **FIPS 186-5** | Digital Signatures | ECDSA/EdDSA |
| **FIPS 197** | AES | AES-256 recommended |
| **FIPS 198-1** | HMAC | HMAC-SHA-384+ |
| **SP 800-56A** | Key Establishment | Applicable |
| **SP 800-56C** | Key Derivation | HKDF recommended |
| **SP 800-133** | Key Generation | RBG requirements |

### 1.4 Zero Trust Architecture Standards

| Standard | Organization | Scope | Status |
|----------|--------------|-------|--------|
| **NIST SP 800-207** | NIST | Zero Trust Architecture | Published |

---

## 2. High-Dimensional QKD Protocol Concepts

### 2.1 Protocol Stack Overview

```text
┌─────────────────────────────────────────────────────────────────┐
│                    HD-QKD PROTOCOL STACK                        │
├─────────────────────────────────────────────────────────────────┤
│  L5: Key Management Interface (ETSI QKD 014/015)                │
├─────────────────────────────────────────────────────────────────┤
│  L4: Privacy Amplification                                      │
│      └── Universal Hash Functions (e.g., Toeplitz)              │
├─────────────────────────────────────────────────────────────────┤
│  L3: Error Correction                                           │
│      └── LDPC Codes (optimized for qudits)                      │
├─────────────────────────────────────────────────────────────────┤
│  L2: Parameter Estimation & Sifting                             │
│      └── d-dimensional basis reconciliation                     │
├─────────────────────────────────────────────────────────────────┤
│  L1: Quantum Transmission                                       │
│      └── d-dimensional quantum states (d = 4, 8, ...)           │
└─────────────────────────────────────────────────────────────────┘
```

### 2.2 Mutually Unbiased Bases (MUBs)

For dimension d (prime power), there exist d+1 MUBs:

| Dimension | MUBs Available | Security Benefit |
|-----------|----------------|------------------|
| d=2 | 3 | Baseline |
| d=4 | 5 | Enhanced |
| d=8 | 9 | Maximum practical |

### 2.3 Error Tolerance Comparison

| Dimension | QBER Tolerance | Capacity (bits/photon) |
|-----------|----------------|------------------------|
| 2D (qubit) | ~11% | 1 |
| 4D (qudit) | ~18% | 2 |
| 8D (qudit) | ~24% | 3 |

---

## 3. ETSI QKD Interfaces

### 3.1 ETSI GS QKD 014 - Key Delivery API

Standard REST API for key delivery between QKD systems and applications.

**Key Endpoints:**

| Operation | Endpoint | Method |
|-----------|----------|--------|
| Get Keys | `/api/v1/keys/{SAE_ID}/enc_keys` | GET |
| Get Key by ID | `/api/v1/keys/{SAE_ID}/enc_keys/{key_ID}` | GET |
| Get Status | `/api/v1/keys/{SAE_ID}/status` | GET |

### 3.2 ETSI GS QKD 015 - Control Interface

SDN control interface for QKD network management.

**Key Operations:**

| Operation | Endpoint | Method |
|-----------|----------|--------|
| Get Status | `/api/v1/status` | GET |
| Get QKD Links | `/api/v1/links` | GET |
| Configure Link | `/api/v1/links/{id}` | PUT |
| Get Metrics | `/api/v1/metrics` | GET |

### 3.3 ETSI GS QKD 018 - Orchestration

Network orchestration interface for multi-domain QKD networks.

---

## 4. Certification Pathways

### 4.1 Draft Compliance Timeline (Public)

PFOLD-CP = MDQC Compliance Program milestone.

| Milestone | Target Window | Focus | Outputs |
|-----------|---------------|-------|---------|
| PFOLD-CP1: ISO + ETSI baseline | 0-12 months | ISO 27001 ISMS scope and gap analysis + ETSI QKD L0 conformance | ISMS scope, gap report, L0 conformance evidence |
| PFOLD-CP2: ISO + ETSI uplift | 12-24 months | ISO 27001 certification + ETSI QKD L1 conformance | ISO 27001 certificate, L1 security evidence |
| PFOLD-CP3: Crypto validation | 18-36 months | FIPS 140-3 module validation + Common Criteria planning | Lab testing, validation plan, security target |
| PFOLD-CP4: Operational audits | 24-36 months | SOC 2 Type II + ETSI interop reports | Audit report, plugfest results |

### 4.2 FIPS 140-3 Considerations

**Recommended Security Levels:**

| Area | Target Level |
|------|--------------|
| Cryptographic Module Specification | Level 2 |
| Roles, Services, Authentication | Level 3 |
| Physical Security (HSM) | Level 3 |
| Sensitive Security Parameters | Level 3 |

### 4.3 Common Criteria (ISO 15408)

**Target:** EAL4+ (augmented)

**Relevant Protection Profiles:**

- PP_ND_V2.2e - Network Device
- CPP_FW_V2.0e - Firewall
- PP_APP_V1.3 - Application Software

### 4.4 ISO 27001:2022

**Key Controls for Quantum Systems:**

| Control | Title | Relevance |
|---------|-------|-----------|
| A.5.15 | Access Control | Key management access |
| A.8.24 | Use of Cryptography | Core requirement |
| A.8.28 | Secure Coding | Implementation security |

---

## 5. Regulatory Compliance

### 5.1 Regional Requirements

| Region | Regulation | Applicability |
|--------|------------|---------------|
| EU | GDPR | Personal data encryption |
| EU | NIS2 | Critical infrastructure |
| EU | DORA | Financial sector |
| USA | CMMC 2.0 | Defense contracts |
| USA | HIPAA | Healthcare data |
| Germany | BSI TR-03116 | Government systems |
| Global | PCI DSS 4.0 | Payment systems |

### 5.2 Export Control

| Item | Classification | Notes |
|------|----------------|-------|
| Quantum hardware | Dual-use | May require license |
| Encryption software | Cat 5 Part 2 | License exceptions available |
| PQC implementations | Publicly available | Generally exempt |

---

## Public Interest and International Coordination

- Align deployments with international law, human rights, and regional public-interest obligations.
- Coordinate cross-border pilots with regulators and standards bodies.
- Avoid use cases that bypass telecom restrictions or enable unlawful surveillance.

---

## Global Regulation and Liability (Read Carefully)

> The repository-wide statement is in [DISCLAIMER.md](../../DISCLAIMER.md). The terms below are reproduced here for convenience and apply to use of this repository.

By using, cloning, or forking this repository, you accept and agree to all terms below:

1. **Mandatory Legal Compliance**  
   You are solely responsible for complying with all applicable local, national, and international laws and regulations where you deploy, operate, or transport this software or related artifacts. This includes (but is not limited to) telecommunications rules, encryption restrictions, and export/import controls. Non-compliance is entirely your responsibility.

2. **Certification and Controls Are Not Provided**  
   This project is research-grade and not FIPS 140-3 validated. If your deployment requires certified modules, audited controls, or formal accreditation, you must obtain and validate them independently. You are solely responsible for data sovereignty and residency obligations (e.g., GDPR, CCPA, or regional data-localization rules).

3. **Export and Dual-Use Obligations**  
   Cryptographic software and quantum-adjacent systems may be classified as dual-use. You must determine applicable classifications and licensing requirements and obtain any required approvals before distribution or cross-border use.

4. **No Warranty; No Liability**  
   This software is provided "as is," without warranty of any kind. The authors disclaim all liability for any use, misuse, or regulatory non-compliance.

---

## 6. Implementation Checklist

### 6.1 Protocol Implementation

- [ ] ETSI QKD 014 API implementation
- [ ] ETSI QKD 015 Control interface
- [ ] High-dimensional sifting protocol
- [ ] Error correction (LDPC for qudits)
- [ ] Privacy amplification
- [ ] Hybrid fallback mechanism (PQC)

See [Conformance Suite](02-conformance-suite.md) for interoperability testing and certification guidance.

### 6.2 Certification Preparation

- [ ] FIPS 140-3 module design documentation
- [ ] Security Target for Common Criteria
- [ ] ISO 27001 ISMS documentation
- [ ] SOC 2 control documentation
- [ ] Risk assessment
- [ ] Incident response procedures

---

## 7. References

1. ETSI GS QKD 014 V1.1.1 - Protocol and data format of REST-based key delivery API
2. ETSI GS QKD 015 V2.1.1 - Control Interface for Software Defined Networks
3. NIST FIPS 203 - Module-Lattice-Based Key-Encapsulation Mechanism Standard
4. NIST FIPS 204 - Module-Lattice-Based Digital Signature Standard
5. NIST SP 800-56C Rev. 2 - Key-Derivation Methods in Key-Establishment Schemes
6. ISO/IEC 27001:2022 - Information security management systems
7. ISO/IEC 15408 - Common Criteria for IT Security Evaluation

---

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

## IP and Publication Sequencing

If you plan to publish, present, or release new material, follow the internal IP sequencing checklist before external disclosure to preserve patentability and trade secrets. Coordinate with maintainers for review timing.

## License

This document is released under the MIT License.
