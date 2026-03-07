package repository

import (
	"crypto/sha1"
	"fmt"
)

func Hash(data []byte) string {
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x", hash)
}