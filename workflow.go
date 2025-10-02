package workflow

// Version represents the current version of the workflow templates
const Version = "v1.0.6"

// SupportedGoVersions lists the Go versions these workflows support
var SupportedGoVersions = []string{
	"1.21",
	"1.22",
	"1.23",
	"1.24",
	"1.25",
}

// WorkflowType represents different types of workflows available
type WorkflowType string

const (
	// ContinuousIntegration represents the main CI workflow
	ContinuousIntegration WorkflowType = "ci"

	// PullRequest represents the PR-specific workflow
	PullRequest WorkflowType = "pr"

	// DependabotAutoMerge represents the Dependabot automation workflow
	DependabotAutoMerge WorkflowType = "dependabot-auto-merge"
)

// IsSupported checks if a Go version is supported by these workflows
func IsSupported(goVersion string) bool {
	for _, supported := range SupportedGoVersions {
		if supported == goVersion {
			return true
		}
	}
	return false
}
