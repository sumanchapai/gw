package git

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"

	"github.com/sumanchapai/gw/cerrors"
	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/utils"
)

type CommandResult struct {
	result string
	err    error
}

// returns error if the provided path is not a git repo, nil otherwise
func IsGitRepo(name string) error {
	if err := utils.FolderExists(name); err != nil {
		return err
	}
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = name
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v %v", string(output), err)
	}
	if !strings.Contains(string(output), "true") {
		return cerrors.ErrPathNotAGitRepo
	}
	return nil
}

// Run arbitrary git command in the provided repo
func run(repo string, args ...string) CommandResult {
	if err := IsGitRepo(repo); err != nil {
		return CommandResult{"", cerrors.ErrPathNotAGitRepo}
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = repo
	output, err := cmd.CombinedOutput()
	return CommandResult{string(output), err}
}

// Run a git command that's not restricted.
// Returns error if restricted command is run.
func SafeRun(repo string, args ...string) CommandResult {
	if len(args) == 0 {
		return CommandResult{"", cerrors.ErrEmptyGitCommand}
	}
	cmd := args[0]
	if slices.Contains(env.RESTRICTED_COMMANDS(), cmd) {
		return CommandResult{"", cerrors.ErrRestrictedCommand}
	}
	return run(repo, args...)
}
