# Changelog

All notable changes to AGILira Workflow Templates will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2025-10-01

### Added
- **Core GitHub Actions Workflows:**
  - `ci.yml` - Complete CI/CD pipeline with comprehensive testing
  - `pr.yml` - Fast pull request validation workflow
  - `dependabot-auto-merge.yml` - Intelligent dependency update management
  
- **Configuration Templates:**
  - `dependabot.yml` - Standardized dependabot configuration for Go projects
  - `.gosec.json` - Security scanner configuration with project-specific rules
  - `.gitignore` - Comprehensive Go project gitignore template
  
- **Documentation:**
  - Professional README with technical specifications
  - Usage examples and implementation guides
  - Security and customization documentation

### Features
- **CI/CD Pipeline:** Multi-stage testing with race detection, security scanning, and coverage reporting
- **Security Integration:** gosec vulnerability scanning with custom rule configuration
- **Quality Assurance:** Code formatting validation, static analysis with staticcheck, Go vet integration
- **Cross-platform Support:** Build matrix testing across Linux, Windows, and macOS
- **Automated Dependency Management:** Smart auto-merge for patch/minor updates with manual review for major versions
- **Codecov Integration:** Automated test coverage reporting and analysis
- **European Timezone Optimization:** Scheduled updates aligned with Italian development hours

### Technical Specifications
- Go stable version support with automatic version management
- CGO_ENABLED configuration for projects requiring C bindings
- Configurable timeout settings for different operating systems
- JSON-formatted gosec output for integration with development tools
- MIT license for broad compatibility and adoption

[Unreleased]: https://github.com/agilira/workflow/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/agilira/workflow/releases/tag/v0.1.0