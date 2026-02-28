package main

import (
	"GO-Project/internal"
	"fmt"
)

func main() {
	err := internal.AddToIndex("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = internal.CreateCommit("first commit")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = internal.ShowLog()
	if err != nil {
		fmt.Println(err)
	}
}
