# Changelog

All notable changes to AGILira Workflow Templates will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-09-29

### Added
- **Reusable Workflows:**
  - `go-ci-full.yml` - Comprehensive CI/CD pipeline for Go projects
  - `go-pr-quick.yml` - Fast PR validation workflow
  - `dependabot-auto-merge.yml` - Intelligent dependency auto-merging
  
- **Configuration Templates:**
  - `dependabot.yml` - Standard dependabot configuration for Go projects
  - `.gosec.json` - Security scanner configuration with AGILira standards
  - `.gitignore` - Comprehensive Go project gitignore template
  
- **Documentation:**
  - Complete README with usage examples
  - Quick start guide for immediate adoption
  - Template and workflow documentation

### Features
- **Go Version Management:** Always use latest stable Go version
- **Security Integration:** Vulnerability scanning, gosec analysis, dependency verification
- **Quality Gates:** Format checking, static analysis, race detection
- **Cross-platform Testing:** Linux, Windows, macOS build matrix
- **Intelligent Auto-merge:** Safe dependency updates with CI validation
- **AGILira Standards:** Consistent configuration across all projects
- **European Timezone:** Dependabot scheduling optimized for Italian development

### Technical Details
- Enterprise-grade workflow design
- Configurable inputs for customization
- Optimized for AGILira project structure
- MIT licensed for maximum compatibility
- Semantic versioning for stable API

[Unreleased]: https://github.com/agilira/workflow/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/agilira/workflow/releases/tag/v1.0.0