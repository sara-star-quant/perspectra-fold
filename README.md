# Perspectra Fold

> High-dimensional quantum and hybrid quantum-classical communication systems.

[![CI](https://github.com/sara-star-quant/perspectra-fold/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/sara-star-quant/perspectra-fold/actions/workflows/ci.yml)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![Code of Conduct](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)
[![Security Policy](https://img.shields.io/badge/security-policy-blue.svg)](SECURITY.md)
[![Python](https://img.shields.io/badge/Python-3.12+-3776AB.svg?logo=python&logoColor=white)](src/python)
[![Rust](https://img.shields.io/badge/Rust-stable-DEA584.svg?logo=rust&logoColor=white)](src/rust)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8.svg?logo=go&logoColor=white)](src/go)
[![PQC: FIPS 203/204/205](https://img.shields.io/badge/PQC-FIPS%20203%2F204%2F205-1f6feb.svg)](docs/standards/01-standards-overview.md)
[![QKD: ETSI GS QKD](https://img.shields.io/badge/QKD-ETSI%20GS%20QKD-1f6feb.svg)](docs/standards/01-standards-overview.md)

## Overview

Perspectra Fold provides research findings, architecture documentation, and production-grade implementations for secure communication systems using:

- **High-dimensional quantum key distribution (HD-QKD)** - 4D and 8D qudit-based systems
- **Hybrid quantum-classical networks** - QKD and classical traffic coexistence
- **Post-quantum cryptography (PQC)** - NIST-standardized algorithms (FIPS 203/204/205)
- **Defense-in-depth security** - Layered cryptographic protection

## Key Findings

| Aspect | Finding |
|--------|---------|
| **Dimensional Advantage** | 8D systems tolerate 24% QBER vs 11% for 2D |
| **Capacity** | 3 bits/photon (8D) vs 1 bit/photon (2D) |
| **Hybrid Networks** | 110+ Tb/s classical + QKD demonstrated in same fiber |
| **Feasibility** | Production pilots ongoing globally |

## Project Structure

```text
.
├── src/                          # Implementation code
│   ├── python/                   # Research utilities (metrics, dimensionality reduction)
│   ├── rust/                     # Core compute (crypto primitives, gRPC service)
│   └── go/                       # Control plane (telemetry, routing, HTTP API)
├── proto/                        # Protocol buffer definitions
├── docs/                         # Documentation
│   ├── research/                 # Literature review and analysis
│   ├── architecture/             # System design and tech stack
│   ├── guides/                   # Implementation and deployment guides
│   └── standards/                # Compliance and certification
├── scripts/                      # Build and test automation
└── .github/                      # CI/CD and GitHub templates
```

## Quick Start

### Prerequisites

- Python 3.12+
- Rust (stable toolchain)
- Go 1.22+
- Protocol Buffers compiler (`protoc`)

### Build and Test

```bash
# Python research utilities
cd src/python && pip install -e .[dev] && pytest -q

# Rust core library
cargo test --manifest-path src/rust/mdqc-core/Cargo.toml

# Rust compute service
cargo build --manifest-path src/rust/compute-service/Cargo.toml

# Go control plane
cd src/go/control-plane && go test ./...

# Integration test
make verify-metrics
```

### Run Services

```bash
# Start Rust compute service
PFOLD_CORE_ADDR=127.0.0.1:50051 cargo run --manifest-path src/rust/compute-service/Cargo.toml

# Start Go control plane (in another terminal)
cd src/go/control-plane
PFOLD_CONTROL_ADDR=:8080 PFOLD_CORE_ADDR=127.0.0.1:50051 go run ./cmd/control-plane
```

## Technology Stack

| Layer | Technology | Purpose |
|-------|------------|---------|
| Quantum Simulation | Qiskit + QuTiP | Research and modeling |
| Performance-Critical | Rust | Cryptographic primitives |
| Control Plane | Go (gRPC) | Orchestration and routing |
| Data Science | Python + NumPy | Metrics and analysis |
| Dimensionality Reduction | UMAP / cuML | High-dimensional data processing |
| Post-Quantum Crypto | liboqs | PQC algorithms |

## Architecture

### Security Model

```text
┌─────────────────────────────────────────┐
│     DEFENSE-IN-DEPTH SECURITY           │
├─────────────────────────────────────────┤
│  Layer 1: QKD Keys (quantum secure)     │
│  Layer 2: PQC Keys (post-quantum)       │
│  Layer 3: Classical Keys (ECDH)         │
├─────────────────────────────────────────┤
│  Combined via HKDF → Session Keys       │
│  Secure if ANY layer remains secure     │
└─────────────────────────────────────────┘
```

### Dimensional Encoding

```text
Input Dimension → Encoding Decision:

D ≤ 8      → Direct quantum encoding (optimal)
8 < D ≤ 17 → Reduce to 8D via UMAP
D > 17     → Hierarchical reduction
```

## Documentation

| Category | Document | Description |
|----------|----------|-------------|
| **Research** | [Research Findings](docs/research/01-research-findings.md) | Literature review and feasibility analysis |
| **Architecture** | [Technology Stack](docs/architecture/01-technology-stack.md) | Languages, frameworks, and tools |
| | [Zero Trust Environment](docs/architecture/02-zero-trust-environment.md) | ZTA/ZTE deployment principles |
| | [ADR: Language Split](docs/architecture/adr/0001-language-split.md) | Architecture decision record |
| **Guides** | [Implementation Guide](docs/guides/01-implementation-guide.md) | Phased implementation approach |
| | [Development Plan](docs/guides/02-public-development-plan.md) | Public development strategy |
| | [Vision 2026-2075](docs/guides/03-vision-2026-2075.md) | Long-horizon roadmap |
| | [Deployment Playbook](docs/guides/04-deployment-playbook.md) | Operational deployment guidance |
| **Standards** | [Standards Overview](docs/standards/01-standards-overview.md) | ETSI, NIST, ISO compliance |
| | [Conformance Suite](docs/standards/02-conformance-suite.md) | Interoperability testing |
| | [Conformance Roadmap](docs/standards/03-conformance-suite-roadmap.md) | Certification pathway |

## Standards Compliance

This project targets compliance with:

- **ETSI QKD** - GS QKD 004, 008, 014, 015, 018
- **NIST PQC** - FIPS 203 (ML-KEM), FIPS 204 (ML-DSA), FIPS 205 (SLH-DSA)
- **ISO 27001** - Information Security Management
- **FIPS 140-3** - Cryptographic Module Validation
- **Common Criteria** - EAL4+ certification path

See [Standards Overview](docs/standards/01-standards-overview.md) for detailed compliance guidance.

## Contributing

Contributions are welcome. Please read our [Contributing Guidelines](CONTRIBUTING.md) before submitting pull requests.

### Areas for Contribution

- Additional encoding scheme implementations
- Performance benchmarks and optimizations
- Integration examples with quantum hardware
- Documentation improvements
- Conformance test cases

## License

Copyright 2026 SARA STAR QUANT LLC

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.

## References

### Key Papers

1. Cozzolino et al. "High-Dimensional Quantum Communication" (2019)
2. NTT Research "High-Dimensional Quantum Dits" (2025)
3. Nature "QKD + 110 Tb/s Classical Coexistence" (2025)

### Standards

1. ETSI GS QKD 014 - Key Delivery API
2. NIST FIPS 203 - ML-KEM (Kyber)
3. NIST FIPS 204 - ML-DSA (Dilithium)

## Acknowledgments

- ETSI ISG QKD for standardization work
- NIST for post-quantum cryptography standards
- Open source quantum computing community
