package cmd

import (
	"fmt"
	"os"
)

func Init() error {

	os.MkdirAll(".minigit/objects", 0755)
	os.MkdirAll(".minigit/refs", 0755)

	os.WriteFile(".minigit/refs/main", []byte(""), 0644)

	os.WriteFile(".minigit/HEAD", []byte("main"), 0644)

	os.WriteFile(".minigit/index", []byte(""), 0644)

	fmt.Println("Initialized empty MiniGit repository")

	return nil
}