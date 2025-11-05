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

// Push stages all changes, commits with the provided message,
// then pushes to remote. If upstream is not set, defaults to origin/<branch>.
func Push(repo string, commitMsg string, stdout, stderr io.Writer) error {
	// Step 1: Try push
	fmt.Fprintln(stdout, "git push")
	pushCmd := exec.Command("git", "push")
	pushCmd.Dir = repo
	pushCmd.Stdout = stdout
	pushCmd.Stderr = stderr
	if err := pushCmd.Run(); err == nil {
		return nil // success on first push
	}

	// Step 2: Handle missing upstream
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
		// Try to set origin if not exists
		fmt.Fprintln(stdout, "no origin found. trying to set origin")
		githubUsername := env.GITHUB_USERNAME()
		if githubUsername == "" {
			return cerrors.ErrNoGithubUsernameSet
		}
		githubRepo := env.GITHUB_REPO()
		if githubRepo == "" {
			return cerrors.ErrNoGithubRepoSet
		}
		remoteFullName := fmt.Sprintf("git@github.com:%v/%v.git", githubUsername, githubRepo)
		fmt.Fprintf(stdout, "git remote add origin %v\n", remoteFullName)
		setOriginCmd := exec.Command("git", "remote", "add", "origin", remoteFullName)
		setOriginCmd.Dir = repo
		setOriginCmd.Stdout = stdout
		setOriginCmd.Stderr = stderr
	}

	// Retry push with upstream set
	fmt.Fprintf(stdout, "git push --set-upstream origin %s\n", branchName)
	upstreamCmd := exec.Command("git", "push", "--set-upstream", "origin", branchName)
	upstreamCmd.Dir = repo
	upstreamCmd.Stdout = stdout
	upstreamCmd.Stderr = stderr
	if err := upstreamCmd.Run(); err != nil {
		return fmt.Errorf("failed to push with upstream: %w", err)
	}

	return nil
}
