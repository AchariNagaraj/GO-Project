package repository

import (
	"fmt"
	"os"
	"strings"
)

// Check if a branch exists
func BranchExists(name string) bool {
	path := ".minigit/refs/" + name
	_, err := os.Stat(path)
	return err == nil
}

// Get latest commit hash of a branch
func GetBranchCommit(branch string) (string, error) {
	data, err := os.ReadFile(".minigit/refs/" + branch)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// Update branch pointer to new commit
func UpdateBranchCommit(branch string, hash string) error {
	path := ".minigit/refs/" + branch
	return os.WriteFile(path, []byte(hash), 0644)
}

// Create a new branch from current branch
func CreateBranch(name string) error {

	if BranchExists(name) {
		return fmt.Errorf("branch already exists")
	}

	currentBranch, err := GetCurrentBranch()
	if err != nil {
		return err
	}

	currentHash, err := GetBranchCommit(currentBranch)
	if err != nil {
		return err
	}

	return os.WriteFile(".minigit/refs/"+name, []byte(currentHash), 0644)
}


//Switching branch
func CheckoutBranch(name string) error {

	if !BranchExists(name) {
		return fmt.Errorf("branch does not exist")
	}

	hash, err := GetBranchCommit(name)
	if err != nil {
		return err
	}

	err = SetCurrentBranch(name)
	if err != nil {
		return err
	}

	if hash != "" && hash != "null" {
		err = RestoreSnapshot(hash)
		if err != nil {
			return err
		}
	}

	fmt.Println("Switched to branch:", name)
	return nil
}