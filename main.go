package main

import (
	"fmt"
	"os"

	"minigit/cmd"
	"minigit/repository"
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

	case "test":
		data := []byte("hello world")

		hash := repository.Hash(data)

		repository.WriteObject(hash, data)

		obj, _ := repository.ReadObject(hash)

		fmt.Println("Stored object:", string(obj))

	default:
		fmt.Println("Unknown command")
	}

}