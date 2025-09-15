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
