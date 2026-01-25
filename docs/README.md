# Perspectra Fold Documentation

This directory contains all project documentation organized by category.

## Structure

```
docs/
├── research/           # Literature review and analysis
├── architecture/       # System design and decisions
│   └── adr/            # Architecture Decision Records
├── guides/             # Implementation and deployment guides
└── standards/          # Compliance and certification
```

## Quick Navigation

### Research

| Document | Description |
|----------|-------------|
| [Research Findings](research/01-research-findings.md) | Literature review on high-dimensional quantum communication |

### Architecture

| Document | Description |
|----------|-------------|
| [Technology Stack](architecture/01-technology-stack.md) | Recommended languages, frameworks, and tools |
| [Zero Trust Environment](architecture/02-zero-trust-environment.md) | ZTA/ZTE deployment principles |
| [ADR: Language Split](architecture/adr/0001-language-split.md) | Decision record for multi-language approach |

### Guides

| Document | Description |
|----------|-------------|
| [Implementation Guide](guides/01-implementation-guide.md) | Phased implementation approach |
| [Development Plan](guides/02-public-development-plan.md) | Public development strategy |
| [Vision 2026-2075](guides/03-vision-2026-2075.md) | Long-horizon infrastructure roadmap |
| [Deployment Playbook](guides/04-deployment-playbook.md) | Operational deployment guidance |

### Standards

| Document | Description |
|----------|-------------|
| [Standards Overview](standards/01-standards-overview.md) | ETSI, NIST, ISO compliance overview |
| [Conformance Suite](standards/02-conformance-suite.md) | Interoperability testing framework |
| [Conformance Roadmap](standards/03-conformance-suite-roadmap.md) | Phased certification pathway |

## Contributing to Documentation

See [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines on contributing to documentation.

### Style Guide

- Use clear, concise language
- Include code examples where appropriate
- Add references for technical claims
- Follow existing markdown formatting
- Run markdown lint before submitting: `markdownlint-cli2 docs/**/*.md`
