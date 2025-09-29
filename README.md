# AGILira Workflow Templates
### Standardized GitHub Actions workflows and CI/CD templates

Reusable GitHub Actions workflows and configuration templates for AGILira projects. Designed for Go projects with professiona-grade quality gates, security scanning, and automated dependency management.

## Quick Start

### 1. Use Reusable Workflows

Add to your project's `.github/workflows/ci.yml`:

```yaml
name: CI/CD
on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  ci:
    uses: agilira/workflow/.github/workflows/go-ci-full.yml@v1
    with:
      go-version: 'stable'
      codecov-token: ${{ secrets.CODECOV_TOKEN }}
```

Add to your project's `.github/workflows/pr.yml`:

```yaml
name: PR Check
on:
  pull_request:
    branches: [ main ]

jobs:
  pr-check:
    uses: agilira/workflow/.github/workflows/go-pr-quick.yml@v1
    with:
      go-version: 'stable'
      gosec-excludes: 'G104,G306,G301'
```

### 2. Copy Configuration Templates

```bash
# Copy dependabot configuration
curl -o .github/dependabot.yml https://raw.githubusercontent.com/agilira/workflow/main/templates/dependabot.yml

# Copy gosec security configuration
curl -o .gosec.json https://raw.githubusercontent.com/agilira/workflow/main/templates/.gosec.json

# Copy Go gitignore
curl -o .gitignore https://raw.githubusercontent.com/agilira/workflow/main/templates/.gitignore
```

### 3. Enable Auto-merge (Optional)

Add to `.github/workflows/dependabot-auto-merge.yml`:

```yaml
name: Dependabot auto-merge
on: pull_request_target

jobs:
  dependabot:
    uses: agilira/workflow/.github/workflows/dependabot-auto-merge.yml@v1
    with:
      ci-check-name: "ci"
```

## Workflows

### Go CI/CD Full (`go-ci-full.yml`)

Comprehensive CI/CD pipeline for Go projects with:

- **Quality Gates**: Format check, vet, staticcheck
- **Security**: Vulnerability scan, gosec security analysis
- **Testing**: Race detection, coverage reporting
- **Cross-platform**: Linux, Windows, macOS builds
- **Dependencies**: Module verification

**Inputs:**
- `go-version` (default: `'stable'`)
- `working-directory` (default: `'.'`)
- `codecov-token` (optional)
- `gosec-config` (default: `'.gosec.json'`)

### Go PR Quick Check (`go-pr-quick.yml`)

Fast validation for pull requests:

- **Quick Quality**: Format, vet, short tests
- **Security Scan**: Basic gosec analysis
- **Fast Feedback**: Optimized for speed

**Inputs:**
- `go-version` (default: `'stable'`)
- `gosec-excludes` (default: `'G104,G306,G301'`)
- `working-directory` (default: `'.'`)

### Dependabot Auto-merge (`dependabot-auto-merge.yml`)

Intelligent auto-merging for dependency updates:

- **Safe Updates**: Auto-merge patch/minor versions
- **CI Integration**: Waits for CI to pass
- **Manual Review**: Major versions require approval

**Inputs:**
- `ci-check-name` (default: `'ci'`)
- `timeout-seconds` (default: `600`)

## Templates

### Dependabot Configuration (`templates/dependabot.yml`)

Standard dependency management:
- Daily Go module updates
- GitHub Actions updates
- AGILira package grouping
- Italian timezone scheduling

### Security Configuration (`templates/.gosec.json`)

Gosec security scanner settings:
- Medium severity/confidence baseline
- Common false positive exclusions
- Generated code filtering

### Go Gitignore (`templates/.gitignore`)

Comprehensive Go project exclusions:
- Build artifacts and binaries
- IDE and OS files
- Test coverage and security reports
- AGILira-specific patterns

## Examples

### Standard Go Library Project

```yaml
# .github/workflows/ci.yml
name: CI/CD
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  ci:
    uses: agilira/workflow/.github/workflows/go-ci-full.yml@v1

  pr-check:
    if: github.event_name == 'pull_request'
    uses: agilira/workflow/.github/workflows/go-pr-quick.yml@v1
```

### Multi-module Project

```yaml
# .github/workflows/ci.yml
name: CI/CD
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  core:
    uses: agilira/workflow/.github/workflows/go-ci-full.yml@v1
    with:
      working-directory: './core'
  
  provider:
    uses: agilira/workflow/.github/workflows/go-ci-full.yml@v1
    with:
      working-directory: './provider'
```

## Best Practices

- **Pin versions** in production: `@v1.2.0` instead of `@v1`
- **Use appropriate workflow** per project type (quick for PRs, full for CI)
- **Configure gosec excludes** based on your project's security requirements
- **Enable codecov integration** for coverage tracking
- **Review major dependency updates** manually before merging

## Versioning

We use semantic versioning for releases:

- **v1.x.x**: Stable API, backward compatible
- **v1.x.0**: New features, minor breaking changes possible
- **v1.0.x**: Bug fixes, fully backward compatible

Pin to specific versions for production:
```yaml
uses: agilira/workflow/.github/workflows/go-ci-full.yml@v1.2.0
```

## License

This project is licensed under the [MIT License](./LICENSE).

---

AGILira Workflow • Standardized DevOps Excellence
