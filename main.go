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
		err := repository.AddToIndex(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}

	case "commit":
		if len(os.Args) < 3 {
			fmt.Println("Usage: minigit commit <message>")
			return
		}
		err := repository.CreateCommit(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}

	case "log":
		err := repository.ShowLog()
		if err != nil {
			fmt.Println(err)
		}

	case "branch":
		if len(os.Args) < 3 {
			fmt.Println("Usage: minigit branch <branch-name>")
			return
		}
		err := repository.CreateBranch(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}

	case "reset":
		if len(os.Args) < 3 {
			fmt.Println("Usage: minigit reset <commit-hash>")
			return
		}
		err := repository.Reset(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}

	case "merge":
		if len(os.Args) < 3 {
			fmt.Println("Usage: minigit merge <branch>")
			return
		}
		err := repository.Merge(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Println("Unknown command")
	}
}