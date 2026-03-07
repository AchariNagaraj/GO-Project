package repository

import (
	"os"
	"strings"
)

const headPath = ".minigit/HEAD"

func GetCurrentBranch() (string, error) {
	data, err := os.ReadFile(headPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func SetCurrentBranch(branch string) error {
	return os.WriteFile(headPath, []byte(branch), 0644)
}