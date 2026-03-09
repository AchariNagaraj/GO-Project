package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ParseCommit reads commit file and returns Commit struct
func ParseCommit(hash string) (*Commit, error) {

	commitPath := filepath.Join(".minigit", "objects", hash)

	data, err := os.ReadFile(commitPath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	commit := &Commit{}
	var files []string
	readFiles := false

	for _, line := range lines {

		if strings.HasPrefix(line, "parent:") {
			commit.Parent = strings.TrimPrefix(line, "parent:")
		} else if strings.HasPrefix(line, "timestamp:") {
			commit.Timestamp = strings.TrimPrefix(line, "timestamp:")
		} else if strings.HasPrefix(line, "message:") {
			commit.Message = strings.TrimPrefix(line, "message:")
		} else if line == "files:" {
			readFiles = true
		} else if readFiles && line != "" {
			files = append(files, line)
		}
	}

	commit.Files = files
	return commit, nil
}

// ShowLog prints commit history
func ShowLog() error {

	if err := checkRepo(); err != nil {
		return err
	}

	current, err := readHEAD()
	if err != nil {
		return err
	}

	if current == "" {
		fmt.Println("No commits yet")
		return nil
	}

	for current != "" {

		commit, err := ParseCommit(current)
		if err != nil {
			return err
		}

		fmt.Println("commit:", current)
		fmt.Println("Date:", commit.Timestamp)
		fmt.Println("Message:", commit.Message)
		fmt.Println("--------------------------")

		current = commit.Parent
	}

	return nil
}
