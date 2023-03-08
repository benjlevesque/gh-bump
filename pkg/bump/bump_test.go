package bump_test

import (
	"testing"

	"github.com/benjlevesque/gh-bump/pkg/bump"
)

func runTest(t *testing.T, current string, bumpType bump.BumpType, expected string) {
	result := bump.GetNewTag(current, bumpType)
	if result != expected {
		t.Fatalf("Expected %q, got %q", expected, result)
	}
}

func TestGetNewTag(t *testing.T) {
	runTest(t, "1.2.3", bump.Patch, "1.2.4")
	runTest(t, "1.2.3", bump.Minor, "1.3.0")
	runTest(t, "1.2.3", bump.Major, "2.0.0")

	runTest(t, "v1.2.3", bump.Patch, "v1.2.4")
	runTest(t, "v1.2.3", bump.Minor, "v1.3.0")
	runTest(t, "v1.2.3", bump.Major, "v2.0.0")
}
