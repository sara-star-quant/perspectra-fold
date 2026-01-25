# Contributing to Perspectra Fold

Thank you for your interest in contributing to the Perspectra Fold project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [How to Contribute](#how-to-contribute)
- [Pull Request Process](#pull-request-process)
- [Style Guidelines](#style-guidelines)
- [Areas for Contribution](#areas-for-contribution)

## Code of Conduct

This project follows a standard code of conduct. Please be respectful, inclusive, and professional in all interactions. See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details.

## Getting Started

1. Fork the repository
2. Clone your fork locally
3. Set up the development environment (see below)
4. Create a feature branch from `main`
5. Make your changes
6. Submit a pull request

## Development Setup

### Prerequisites

- Python 3.12+
- Rust (stable toolchain with `rustfmt` and `clippy`)
- Go 1.22+
- Protocol Buffers compiler (`protoc`)

### Environment Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/perspectra-fold.git
cd perspectra-fold

# Python setup
cd src/python
python -m venv .venv
source .venv/bin/activate  # or .venv\Scripts\activate on Windows
pip install -e .[dev]
cd ../..

# Rust setup (ensure toolchain is installed)
rustup component add rustfmt clippy

# Go setup
cd src/go/control-plane
go mod download
cd ../../..
```

### Running Tests

```bash
# Python
cd src/python && pytest -q

# Rust core
cargo test --manifest-path src/rust/mdqc-core/Cargo.toml

# Rust compute service
cargo test --manifest-path src/rust/compute-service/Cargo.toml

# Go
cd src/go/control-plane && go test ./...

# Integration tests
make verify-metrics
```

## How to Contribute

### Reporting Issues

- Search existing issues before creating a new one
- Use the appropriate issue template
- Provide clear reproduction steps for bugs
- Include relevant environment information

### Submitting Code

1. Create a feature branch: `git checkout -b feature/your-feature`
2. Write tests for new functionality
3. Ensure all tests pass
4. Follow the style guidelines for the relevant language
5. Commit with clear, descriptive messages
6. Push to your fork and open a pull request

### Submitting Documentation

- Follow existing markdown formatting
- Include code examples where appropriate
- Add references for technical claims
- Ensure links are valid (checked by CI)

## Pull Request Process

1. **Title**: Use a clear, descriptive title
2. **Description**: Explain what changes you made and why
3. **Tests**: Ensure all CI checks pass
4. **Review**: Address reviewer feedback promptly
5. **Merge**: Maintainers will merge after approval

### Commit Message Format

```
<type>: <short description>

<optional body explaining the change>

<optional footer with references>
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

Example:
```
feat: add UMAP dimensionality reduction support

Implements GPU-accelerated UMAP for high-dimensional data
reduction to 8D qudit encoding space.

Closes #42
```

## Style Guidelines

### Python

- Follow PEP 8
- Use type hints
- Format with `black` or `ruff format`
- Lint with `ruff`

### Rust

- Run `cargo fmt` before committing
- Address all `clippy` warnings
- Follow Rust API guidelines

### Go

- Run `gofmt` before committing
- Follow Effective Go guidelines
- Use meaningful variable names

### Documentation

- Use clear, concise language
- Keep line length reasonable (< 120 characters)
- Use fenced code blocks with language specifiers
- Include diagrams for complex concepts (ASCII art or Mermaid)

## Areas for Contribution

### High Priority

- **Conformance Tests**: ETSI QKD protocol compliance tests
- **Benchmarks**: Performance measurements for core operations
- **Hardware Integration**: Examples with real quantum hardware

### Medium Priority

- **Documentation**: Tutorials, examples, clarifications
- **Additional Encodings**: New qudit encoding schemes
- **Tooling**: Developer experience improvements

### Research

- Literature review updates
- New protocol analysis
- Security proofs

## Security

If you discover a security vulnerability, please follow our [Security Policy](SECURITY.md). Do not open a public issue for security vulnerabilities.

## Questions?

- Open a [Discussion](https://github.com/sarastarquant/perspectra-fold/discussions) for questions
- Open an [Issue](https://github.com/sarastarquant/perspectra-fold/issues) for bugs or feature requests

## License

By contributing, you agree that your contributions will be licensed under the Apache License 2.0.
