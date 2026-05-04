# Disclaimer

This project, Perspectra Fold, is a research-grade reference implementation of high-dimensional quantum and hybrid quantum-classical communication concepts. The disclaimers below are binding on anyone using, cloning, forking, deploying, or distributing the code or artifacts in this repository.

## What this project is

- A working set of references, simulations, and architecture notes for HD-QKD, qudit-based encoding, hybrid network design, and post-quantum cryptography.
- Algorithm-level alignment with NIST FIPS 203 (ML-KEM), FIPS 204 (ML-DSA), and FIPS 205 (SLH-DSA), via well-known libraries (such as liboqs and equivalents).
- A target alignment with ETSI GS QKD 014/015 API surfaces.

## What this project is not

- Not a certified or validated cryptographic product.
- Not a turnkey production system, an operational QKD network, or a deployable security control.
- Not legal, regulatory, or compliance advice.

## Certification status (none)

The following are roadmap targets only. None has been achieved by this repository:

- **FIPS 140-3** (Cryptographic Module Validation Program, CMVP). No module-level validation has been performed by an accredited lab; no module identifier has been issued.
- **Common Criteria (EAL4+).** No security target, evaluation, or certificate exists.
- **ISO 27001.** No information-security management system audit has been conducted.
- **SOC 2 / FedRAMP / regional accreditations.** None.

Algorithm-standards alignment (FIPS 203/204/205, ETSI GS QKD) is **not** equivalent to module or system certification. If a deployment requires certified modules, audited controls, or formal accreditation, you must obtain and validate them independently.

## Cryptographic and security caveats

- Cryptographic implementations in this repository have not undergone independent security audit, formal verification, or constant-time analysis. Side-channel and fault-injection resistance is out of scope.
- Random number generation, key storage, and lifecycle management in this repository are illustrative. Production deployments require certified hardware (HSMs, validated PRNGs) outside the scope of this project.
- Cryptographic agility is a design goal, but does not constitute a guarantee against future cryptanalytic advances, including against the post-quantum primitives.
- Any "defense in depth" claims describe the architectural intent (QKD plus PQC plus classical), not a measured security level.

## Quantum hardware caveats

- Where this repository discusses QKD, qudits, or photonic encoding, the implementation is a model or simulation unless an explicit hardware integration is documented. Real QKD requires real hardware (sources, detectors, channels) that this repository does not provide.
- Performance numbers in research findings are drawn from cited literature, not from in-repo measurements, unless explicitly stated.

## Export controls and dual-use

Cryptographic software and quantum-adjacent systems may be classified as dual-use under export-control regimes (for example EAR in the US, EU Dual-Use Regulation 2021/821, Wassenaar Arrangement). You are responsible for determining applicable classifications and obtaining any licenses required before distribution, transfer, or cross-border use.

## Regulatory responsibility

You are solely responsible for compliance with all laws and regulations applicable where you deploy, operate, or transport this software or any artifacts derived from it. This includes telecommunications rules, encryption restrictions, data-protection regimes (such as GDPR, CCPA, regional residency rules), and sector-specific frameworks. Non-compliance is your responsibility, not the authors'.

## No warranty, no liability

This software is provided "as is," without warranty of any kind, express or implied, including the implied warranties of merchantability, fitness for a particular purpose, and non-infringement. The authors and contributors disclaim all liability for any direct, indirect, incidental, special, consequential, or punitive damages arising from any use or misuse of this software, regardless of theory.

This disclaimer supplements the [Apache License 2.0](LICENSE) under which the code is distributed; in case of conflict, the terms of the Apache License 2.0 govern.

## Reporting concerns

For security-related concerns, see [SECURITY.md](SECURITY.md). For documentation issues with this disclaimer or the standards-alignment claims, open a regular issue or pull request.
