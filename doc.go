// Package workflow provides standardized CI/CD workflow templates for Go projects.
//
// This repository contains production-ready GitHub Actions workflows specifically
// designed for Go applications and libraries. The workflows include:
//
//   - Comprehensive testing with race detection
//   - Multi-platform builds (Linux, Windows, macOS)
//   - Static analysis (go vet, staticcheck)
//   - Enhanced security scanning (gosec + govulncheck)
//   - Dependency integrity verification (go mod verify)
//   - Code coverage reporting
//   - Dependabot integration with auto-merge
//
// Usage:
//
// Copy the desired workflow files from the templates/ directory to your
// project's .github/workflows/ directory and customize as needed.
//
// The workflows are designed to work out-of-the-box for most Go projects
// while being easily customizable for specific requirements.
//
// For more information, see the README.md file.
package workflow
