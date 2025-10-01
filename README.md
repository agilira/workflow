# AGILira Workflow Templates

Standardized GitHub Actions workflows and configuration templates for AGILira projects. Designed for Go projects with professional-grade quality gates, security scanning, and automated dependency management.

## Overview

This repository provides:

- **GitHub Actions Workflows** - Complete CI/CD pipelines
- **Configuration Templates** - Standardized configuration files
- **Documentation** - Usage guides and examples

## Available Workflows

### Complete CI/CD Pipeline (`ci.yml`)

Comprehensive continuous integration and deployment workflow:

```yaml
# Copy .github/ci.yml to your project
# Provides full test suite with quality gates
```

Features:
- Test execution with race condition detection
- Security scanning using gosec
- Code coverage reporting with Codecov integration
- Static analysis with staticcheck
- Cross-platform build matrix (Linux, Windows, macOS)
- Latest stable Go version support

### Pull Request Validation (`pr.yml`)

Fast validation workflow for pull requests:

```yaml
# Quick validation including:
# - Code formatting verification
# - Go vet analysis
# - Basic test execution
# - Security scanning
```

Optimized for:
- Rapid developer feedback
- Essential quality checks
- Basic security validation

### Dependabot Auto-merge (`dependabot-auto-merge.yml`)

Intelligent automated dependency management:

```yaml
# Automated merging for:
# - Patch and minor updates after CI validation
# - Manual review requirement for major updates
```

Capabilities:
- Automated merge for safe dependency updates
- Manual review required for major version changes
- Mandatory CI validation before merge
- Automated informational comments

## Configuration Templates

### Dependabot Configuration (`templates/dependabot.yml`)

Standardized dependency management configuration:

```yaml
# Copy to .github/dependabot.yml
# Optimized for Go projects with AGILira standards
```

Configuration includes:
- Daily updates scheduled at 09:00 Europe/Rome timezone
- Intelligent dependency grouping
- Automatic assignment to AGILira team
- Pull request limits to prevent noise

### Gosec Security Configuration (`templates/.gosec.json`)

Custom security scanning template:

```json
// Optimized configuration for AGILira projects
// Excludes common false positives
// Includes custom security pattern rules
```

### Go Gitignore Template (`templates/.gitignore`)

Comprehensive gitignore for Go projects:

```gitignore
# Complete template for AGILira Go projects
# Includes IDE, OS, and development tool exclusions
```

## Usage

### Quick Setup

Copy workflows to your project:

```bash
# Create required directory structure
mkdir -p .github/workflows

# Copy primary workflows
cp workflow/.github/ci.yml .github/workflows/
cp workflow/.github/pr.yml .github/workflows/
cp workflow/.github/dependabot-auto-merge.yml .github/workflows/

# Copy configuration templates
cp workflow/templates/dependabot.yml .github/
cp workflow/templates/.gosec.json .
cp workflow/templates/.gitignore .
```

### Customization

Adapt templates to your requirements:

```yaml
# In ci.yml, customize if needed:
env:
  CGO_ENABLED: 1  # Modify if your project doesn't use CGO
  
# In dependabot.yml:
schedule:
  time: "09:00"  # Adjust for your timezone if necessary
```

### Codecov Configuration (Optional)

For enhanced coverage reporting, add to your repository:

```yaml
# .codecov.yml
coverage:
  range: "50...80"
  status:
    project:
      default:
        target: 70%
```

## Requirements

- **Go**: Stable version (managed automatically by workflows)
- **Repository**: Standard Go project structure
- **Secrets**: `GITHUB_TOKEN` (provided automatically by GitHub)
- **Codecov**: Optional token for advanced coverage reporting

## Implementation Examples

### Standard Go Project

For a typical Go project setup:

```bash
# 1. Apply base templates
cp workflow/templates/.gitignore .
cp workflow/templates/.gosec.json .

# 2. Setup essential workflows
cp workflow/.github/ci.yml .github/workflows/
cp workflow/.github/pr.yml .github/workflows/

# 3. Enable dependabot
cp workflow/templates/dependabot.yml .github/
```

### Go Project with Automated Dependency Management

For projects requiring full automation:

```bash
# Complete setup with auto-merge capabilities
cp workflow/.github/*.yml .github/workflows/
cp workflow/templates/* .github/
```

## Security Features

The workflows implement multiple security layers:

- **Gosec**: Code vulnerability scanning
- **Dependabot**: Secure automated dependency updates
- **Go vet**: Built-in static analysis
- **Staticcheck**: Advanced code analysis
- **Race detection**: Concurrent execution validation

## Contributing

To improve the templates:

1. Fork the repository
2. Create a feature branch
3. Test with real projects
4. Submit pull request with detailed description

## License

This project is licensed under the [MIT License](./LICENSE).

---

workflow • An AGILira tool
