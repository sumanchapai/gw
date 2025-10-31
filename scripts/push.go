package scripts

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/sumanchapai/gw/cerrors"
	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/git"
)

// CommitAndPush stages all changes, commits with the provided message,
// then pushes to remote. If upstream is not set, defaults to origin/<branch>.
func CommitAndPush(commitMsg string, stdout, stderr io.Writer) error {
	repo := env.GW_REPO()

	// Step 1: Commit
	if err := CommitAll(commitMsg, stdout, stderr); err != nil {
		return err
	}

	// Step 2: Try push
	fmt.Fprintln(stdout, "git push")
	pushCmd := exec.Command("git", "push")
	pushCmd.Stdout = stdout
	pushCmd.Stderr = stderr
	if err := pushCmd.Run(); err == nil {
		return nil // success on first push
	}

	// Step 3: Handle missing upstream
	log.Println("Push failed, attempting to set upstream...")

	// Detect current branch
	branchName, err := git.CurrentBranch(repo)
	if err != nil {
		log.Println("err getting current branch", err)
		fmt.Fprintf(stderr, "error getting current branch")
		return err
	}

	// Detect if "origin" exists
	if !git.RemoteExists(repo, "origin") {
		return cerrors.ErrNoOriginRemoteExists
	}

	// Retry push with upstream set
	fmt.Fprintf(stdout, "git push --set-upstream origin %s\n", branchName)
	upstreamCmd := exec.Command("git", "push", "--set-upstream", "origin", branchName)
	upstreamCmd.Stdout = stdout
	upstreamCmd.Stderr = stderr
	if err := upstreamCmd.Run(); err != nil {
		return fmt.Errorf("failed to push with upstream: %w", err)
	}

	return nil
}
