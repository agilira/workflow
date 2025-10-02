package workflow

import "testing"

func TestIsSupported(t *testing.T) {
	tests := []struct {
		version  string
		expected bool
	}{
		{"1.21", true},
		{"1.22", true},
		{"1.23", true},
		{"1.24", true},
		{"1.25", true},
		{"1.19", false},
		{"1.26", false},
		{"invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			result := IsSupported(tt.version)
			if result != tt.expected {
				t.Errorf("IsSupported(%q) = %v, want %v", tt.version, result, tt.expected)
			}
		})
	}
}

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}

	if Version != "v1.0.6" {
		t.Errorf("Version = %q, want %q", Version, "v1.0.6")
	}
}

func TestWorkflowTypes(t *testing.T) {
	expectedTypes := []WorkflowType{
		ContinuousIntegration,
		PullRequest,
		DependabotAutoMerge,
	}

	for _, wt := range expectedTypes {
		if string(wt) == "" {
			t.Errorf("WorkflowType %v should not be empty", wt)
		}
	}
}
