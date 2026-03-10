package repository

import (
	"fmt"
	"os"
	"strings"
)

// Hard reset implementation
func Reset(commitHash string) error {

	if !ObjectExists(commitHash) {
		return fmt.Errorf("commit does not exist")
	}

	currentBranch, err := GetCurrentBranch()
	if err != nil {
		return err
	}

	// Move branch pointer
	err = UpdateBranchCommit(currentBranch, commitHash)
	if err != nil {
		return err
	}

	return RestoreSnapshot(commitHash)
}

// Restore working directory snapshot
func RestoreSnapshot(commitHash string) error {

	commit, err := ParseCommit(commitHash)
	if err != nil {
		return err
	}

	// Delete previously tracked files
	currentBranch, err := GetCurrentBranch()
	if err != nil {
		return err
	}

	previousHash, err := GetBranchCommit(currentBranch)
	if err != nil {
		return err
	}

	if previousHash != "" && previousHash != "null" {

		previousCommit, err := ParseCommit(previousHash)
		if err != nil {
			return err
		}

		for _, filename := range previousCommit.Files {
			if err := os.Remove(filename); err != nil && !os.IsNotExist(err) {
				return err
			}
		}
	}

	// Restore files from blobs
	for _, entry := range commit.Files {

		parts := strings.Split(entry, " ")

		if len(parts) != 2 {
			return fmt.Errorf("invalid file entry: %s", entry)
		}

		blobHash := parts[0]
		filename := parts[1]

		blobData, err := ReadObject(blobHash)
		if err != nil {
			return err
		}

		err = os.WriteFile(filename, blobData, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// Check if current commit is ancestor of target
func IsAncestor(current string, target string) bool {

	for target != "null" {

		if target == current {
			return true
		}

		commit, err := ParseCommit(target)
		if err != nil {
			return false
		}

		target = commit.Parent
	}

	return false
}

// Fast-forward merge only
func Merge(branch string) error {

	if !BranchExists(branch) {
		return fmt.Errorf("branch does not exist")
	}

	currentBranch, err := GetCurrentBranch()
	if err != nil {
		return err
	}

	currentHash, err := GetBranchCommit(currentBranch)
	if err != nil {
		return err
	}

	targetHash, err := GetBranchCommit(branch)
	if err != nil {
		return err
	}

	if IsAncestor(currentHash, targetHash) {

		// Move pointer
		err = UpdateBranchCommit(currentBranch, targetHash)
		if err != nil {
			return err
		}

		return RestoreSnapshot(targetHash)

	} else {
		return fmt.Errorf("non fast-forward merge not supported")
	}
}