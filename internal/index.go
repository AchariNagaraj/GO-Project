package internal

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
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

	// check file exists in working directory
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("file does not exist")
	}

	indexPath := filepath.Join(".minigit", "index")

	// open index file in append mode
	file, err := os.OpenFile(indexPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write file path
	_, err = file.WriteString(filePath + "\n")
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
		return nil, err
	}
	defer file.Close()

	var files []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		files = append(files, scanner.Text())
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
