//read write objects from .minigit/objects
package repository

import (
	"os"
)

const objectsPath = ".minigit/objects/"

func WriteObject(hash string, data []byte) error {
	path := objectsPath + hash
	return os.WriteFile(path, data, 0644)
}

func ReadObject(hash string) ([]byte, error) {
	path := objectsPath + hash
	return os.ReadFile(path)
}

func ObjectExists(hash string) bool {
	path := objectsPath + hash
	_, err := os.Stat(path)
	return err == nil
}