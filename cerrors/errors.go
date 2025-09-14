package cerrors

import "errors"

var ErrPathDoesntExist = errors.New("Path doesn't exist")
var ErrPathNotAFolder = errors.New("Path is not a folder")
var ErrPathNotAGitRepo = errors.New("Path is not a git repo")
var ErrEmptyGitCommand = errors.New("No git command provided")
var ErrRestrictedCommand = errors.New("Restricted command")
