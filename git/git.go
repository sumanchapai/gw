package git

import (
	"fmt"
	"log"
	"os/exec"
	"slices"
	"strings"

	"github.com/sumanchapai/gw/cerrors"
	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/utils"
)

type CommandResult struct {
	Result string
	Err    error
}

// IsGitRepo returns error if the provided path is not a git repo, nil otherwise
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

// run exectues arbitrary git command in the provided repo
func run(repo string, args ...string) CommandResult {
	if err := IsGitRepo(repo); err != nil {
		return CommandResult{"", cerrors.ErrPathNotAGitRepo}
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = repo
	output, err := cmd.CombinedOutput()
	return CommandResult{string(output), err}
}

// SafeRun exectues a git command that's not restricted.
// returns error if restricted command is run.
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

// Branches returns the list of Git branches
func Branches(repo string) ([]string, error) {
	if ok, _ := hasCommits(repo); !ok {
		return []string{}, nil
	}
	x := SafeRun(repo, "branch", "--list", "--format=%(refname:short)")
	if x.Err != nil {
		log.Println(x.Result)
		log.Println(x.Err)
		return nil, x.Err
	}
	branches := strings.Split(strings.TrimSpace(x.Result), "\n")
	return branches, nil
}

// CurrentBranch returns the name of the current Git branch
func CurrentBranch(repo string) (string, error) {
	if ok, _ := hasCommits(repo); !ok {
		return defaultBranch(repo)
	}

	x := SafeRun(repo, "rev-parse", "--abbrev-ref", "HEAD")
	if x.Err != nil {
		log.Println(x.Result)
		log.Println(x.Err)
		return "", x.Err
	}

	branch := strings.TrimSpace(string(x.Result))
	return branch, nil
}

// defaultBranch returns the default branch in the repo
func defaultBranch(repo string) (string, error) {
	defaultRun := SafeRun(repo, "symbolic-ref", "--short", "HEAD")
	return strings.TrimSpace(defaultRun.Result), defaultRun.Err
}

// hasCommits returns true if the repo any has commits
func hasCommits(repo string) (bool, error) {
	x := SafeRun(repo, "rev-parse", "HEAD")
	return x.Err == nil, x.Err
}
