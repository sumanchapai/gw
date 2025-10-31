package scripts

import (
	"fmt"
	"io"
	"log"

	"github.com/sumanchapai/gw/cerrors"
	"github.com/sumanchapai/gw/git"
)

// CreatePR commits, pushes, and creates a PR if one doesn't exist
// Current branch is the head (example. feature) and default branch
// of the default remote is considered base (example. origin/main)
// If the head and base are the same, aborts asking to create PR
// from a different branch. If a PR already exists on the base branch,
// aborts creating PR (only pushes).
func CreatePR(repo string, commitMsg string, stdout, stderr io.Writer) error {
	currentBranch, err := git.CurrentBranch(repo)
	if err != nil {
		log.Println("err getting current branch", err)
		fmt.Fprintf(stderr, "error getting current branch")
		return err
	}
	// TODO:
	// Get default remote branch name
	var defaultRemoteBranchName string
	if currentBranch == defaultRemoteBranchName {
		fmt.Fprintf(stderr, "%s", cerrors.ErrCantCreatePRFromDefaultBranch.Error())
		return cerrors.ErrCantCreatePRFromDefaultBranch
	}
	err = CommitAndPush(repo, commitMsg, stdout, stderr)
	if err != nil {
		return err
	}
	// TODO:
	// Create PR
	return nil
}
