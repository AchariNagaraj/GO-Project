package main

import (
	"fmt"
	"os"

	"goproject/cmd"
	"goproject/repository"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: minigit <command>")
		return
	}

	switch os.Args[1] {

	case "init":
		err := cmd.Init()
		if err != nil {
			fmt.Println(err)
		}

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: minigit add <file>")
			return
		}
		repository.AddToIndex(os.Args[2])

	case "commit":
		if len(os.Args) < 3 {
			fmt.Println("Usage: minigit commit <message>")
			return
		}
		repository.CreateCommit(os.Args[2])

	case "log":
		repository.ShowLog()

	default:
		fmt.Println("Unknown command")
	}
}