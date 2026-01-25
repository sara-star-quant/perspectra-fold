# ADR 0001: Language Split (Python/Rust/Go)

**Status:** Accepted
**Date:** 2026-01-21
**Owner:** Principal Software Engineer
**Reviewers:** Director of Digital Transformation, Security Lead, Research Lead, Standards and Policy Advisor

## Context

The system requires fast research iteration, high-performance computation for high-dimensional operations, and reliable service orchestration for control plane components.

## Decision

- **Python 3.12+** for research, analysis, and simulation orchestration.
- **Rust** for core compute and cryptographic primitives.
- **Go** for control plane services and network orchestration.
- Track the latest stable Python releases (3.13/3.14 when available) for throughput improvements.

## Consequences

- Use gRPC/protobuf between the Go control plane and Rust compute services.
- Reserve FFI bindings for local research workflows (Python calling Rust) when needed.
- CI will require separate lint/test stages per language when code is introduced.
