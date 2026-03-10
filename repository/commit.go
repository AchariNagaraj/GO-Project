package repository

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Commit structure
type Commit struct {
	Parent    string
	Timestamp string
	Message   string
	Files     []string
}

// helper to hash content
func hashContent(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

// CreateCommit creates a new commit
func CreateCommit(message string) error {

	// check repo
	if err := checkRepo(); err != nil {
		return err
	}

	// read staged files
	files, err := ReadIndex()
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("nothing to commit")
	}

	var blobEntries []string

	// create blobs
	for _, file := range files {

		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		hash := hashContent(content)

		objectPath := filepath.Join(".minigit", "objects", hash)

		// write blob if not exists
		if _, err := os.Stat(objectPath); os.IsNotExist(err) {
			err = os.WriteFile(objectPath, content, 0644)
			if err != nil {
				return err
			}
		}

		blobEntries = append(blobEntries, fmt.Sprintf("%s %s", hash, file))
	}

	// get current branch
	currentBranch, err := GetCurrentBranch()
	if err != nil {
		return err
	}

	// get parent commit from branch
	parent, _ := GetBranchCommit(currentBranch)

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return err
	}

	timestamp := time.Now().In(loc).Format("2006-01-02 15:04:05 MST")

	commitContent := fmt.Sprintf(
		"parent:%s\n"+
			"timestamp:%s\n"+
			"message:%s\n"+
			"files:\n%s\n",
		parent,
		timestamp,
		message,
		strings.Join(blobEntries, "\n"),
	)

	commitHash := hashContent([]byte(commitContent))

	commitPath := filepath.Join(".minigit", "objects", commitHash)

	err = os.WriteFile(commitPath, []byte(commitContent), 0644)
	if err != nil {
		return err
	}

	// update branch pointer instead of HEAD
	err = UpdateBranchCommit(currentBranch, commitHash)
	if err != nil {
		return err
	}

	// clear index
	err = ClearIndex()
	if err != nil {
		return err
	}

	fmt.Println("Committed as:", commitHash)
	return nil
}