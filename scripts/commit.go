package scripts

import (
	"fmt"
	"io"
	"os/exec"
)

// CommitAll stages all changes and commits with the provided message.
// `stdout` and `stderr` can be used to stream output (for example, to websocket).
func CommitAll(repo string, commitMsg string, stdout, stderr io.Writer) error {
	if commitMsg == "" {
		return fmt.Errorf("commit message not provided")
	}

	// Stage all changes
	fmt.Fprintln(stdout, "git add -A")
	addCmd := exec.Command("git", "add", "-A")
	addCmd.Dir = repo
	addCmd.Stdout = stdout
	addCmd.Stderr = stderr
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("git add failed: %w", err)
	}

	// Commit
	fmt.Fprintf(stdout, "git commit -m %q\n", commitMsg)
	commitCmd := exec.Command("git", "commit", "-m", commitMsg)
	commitCmd.Dir = repo
	commitCmd.Stdout = stdout
	commitCmd.Stderr = stderr
	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
