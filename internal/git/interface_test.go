package git_test

import (
	"github.com/xcawolfe-amzn/gastown/internal/beads"
	"github.com/xcawolfe-amzn/gastown/internal/git"
)

// Compile-time assertion: Git must satisfy BranchChecker.
var _ beads.BranchChecker = (*git.Git)(nil)
