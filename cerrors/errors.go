package cerrors

import "errors"

var ErrPathDoesntExist = errors.New("Path doesn't exist")
var ErrPathNotAFolder = errors.New("Path is not a folder")
var ErrPathNotAFile = errors.New("Path is not a file")
var ErrPathNotAGitRepo = errors.New("Path is not a git repo")
var ErrEmptyGitCommand = errors.New("No git command provided")
var ErrRestrictedCommand = errors.New("Restricted command")
var ErrCantDetectCurrentPath = errors.New("Cannot detect path")
var ErrNoOriginRemoteExists = errors.New("No 'origin' remote found, please add one manually")
var ErrNoGithubUsernameSet = errors.New("No GITHUB_USERNAME set in .env")
var ErrNoGithubRepoSet = errors.New("No GITHUB_REPO set in .env")
var ErrCouldNotSetRemote = errors.New("Could not set remote url")
var ErrNoGithubRepoTokenSet = errors.New("No GITHUB_REPO_TOKEN set in .env")
var ErrCantCreatePRFromDefaultBranch = errors.New("Cannot create PR from a default branch. Switch to a different branch instead.")
