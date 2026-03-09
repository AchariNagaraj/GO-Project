package internal

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// helper function to check repo initialized
func checkRepo() error {
	if _, err := os.Stat(".minigit"); os.IsNotExist(err) {
		return errors.New("repository not initialized. Run init first")
	}
	return nil
}

// AddToIndex adds a file path to .minigit/index
func AddToIndex(filePath string) error {

	// check repo exists
	if err := checkRepo(); err != nil {
		return err
	}

	cleanPath := filepath.Clean(filePath)

	// check file exists in working directory
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		return errors.New("file does not exist")
	}

	// prevent duplicate entries
	existing, err := ReadIndex()
	if err != nil {
		return err
	}
	for _, f := range existing {
		if filepath.Clean(strings.TrimSpace(f)) == cleanPath {
			return nil
		}
	}

	indexPath := filepath.Join(".minigit", "index")

	// open index file in append mode (create if missing)
	file, err := os.OpenFile(indexPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write file path
	_, err = file.WriteString(cleanPath + "\n")
	return err
}

// ReadIndex reads all staged files
func ReadIndex() ([]string, error) {

	if err := checkRepo(); err != nil {
		return nil, err
	}

	indexPath := filepath.Join(".minigit", "index")

	file, err := os.Open(indexPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var files []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		files = append(files, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

// ClearIndex empties the index file
func ClearIndex() error {

	if err := checkRepo(); err != nil {
		return err
	}

	indexPath := filepath.Join(".minigit", "index")

	// overwrite with empty content
	return os.WriteFile(indexPath, []byte(""), 0644)
}
