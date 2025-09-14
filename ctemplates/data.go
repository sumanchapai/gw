package ctemplates

import (
	"log"

	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/git"
)

// GetRootData builds the datatype that the root template requires.
func GetRootData() (*RootData, error) {
	repo := env.GW_REPO()

	currentBranch, err := git.CurrentBranch(repo)
	if err != nil {
		log.Println("could not get current branch")
		return nil, err
	}

	branches, err := git.Branches(repo)
	if err != nil {
		log.Println("could not get git branches")
		return nil, err
	}
	otherBranches := make([]string, 0)
	for _, b := range branches {
		if b != currentBranch {
			otherBranches = append(otherBranches, b)
		}
	}
	return &RootData{
		CurrentBranch: currentBranch,
		OtherBranches: otherBranches,
		BackLink:      env.BACK_LINK(),
	}, nil
}
