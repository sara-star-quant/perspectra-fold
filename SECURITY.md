# Security Policy

## Reporting a Vulnerability

The Perspectra Fold project takes security seriously. If you discover a security vulnerability, please report it responsibly.

### How to Report

**Do not open a public GitHub issue for security vulnerabilities.**

Instead, please email the maintainers directly or use GitHub's private vulnerability reporting feature:

1. Go to the repository's Security tab
2. Click "Report a vulnerability"
3. Provide detailed information about the vulnerability

### What to Include

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

### Response Timeline

- **Acknowledgment**: Within 48 hours
- **Initial Assessment**: Within 1 week
- **Resolution Timeline**: Communicated after assessment

## Supported Versions

| Version | Supported          |
|---------|--------------------|
| main    | :white_check_mark: |
| < 1.0   | :x:                |

## Security Considerations

### Cryptographic Components

This project implements cryptographic primitives. Users should:

- Use only stable, reviewed releases in production
- Follow the deployment guidelines in the documentation
- Keep dependencies updated
- Conduct independent security audits for production deployments

### Key Material

- Never commit key material, secrets, or credentials
- Use environment variables or secure secret management
- Follow the zero-trust principles outlined in the architecture documentation

### Dependencies

We monitor dependencies for known vulnerabilities using:

- GitHub Dependabot
- `cargo audit` for Rust
- `pip-audit` for Python
- `govulncheck` for Go

## Disclosure Policy

We follow a coordinated disclosure process:

1. Reporter submits vulnerability privately
2. We acknowledge and assess the report
3. We develop and test a fix
4. We release the fix and publish an advisory
5. We credit the reporter (unless anonymity is requested)

## Security Best Practices

When deploying Perspectra Fold components:

1. **Network Security**: Use TLS for all communications
2. **Access Control**: Implement proper authentication and authorization
3. **Monitoring**: Enable logging and anomaly detection
4. **Updates**: Keep all components and dependencies current
5. **Isolation**: Run services with minimal privileges

See [Zero Trust Environment](docs/architecture/02-zero-trust-environment.md) for detailed security architecture guidance.
