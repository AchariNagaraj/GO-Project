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

// read HEAD
func readHEAD() (string, error) {
	headPath := filepath.Join(".minigit", "HEAD")
	data, err := os.ReadFile(headPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// update HEAD
func updateHEAD(commitHash string) error {
	headPath := filepath.Join(".minigit", "HEAD")
	return os.WriteFile(headPath, []byte(commitHash), 0644)
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

	// get parent commit
	parent, _ := readHEAD()

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return err // or handle as needed
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

	// update HEAD
	err = updateHEAD(commitHash)
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
