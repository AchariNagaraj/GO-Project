package main

import (
	"fmt"
	"os"

	"minigit/cmd"
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

	default:
		fmt.Println("Unknown command")
	}
}