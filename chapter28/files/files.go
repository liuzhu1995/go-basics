package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir("/etc/hosts")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {

		if file.IsDir() {
			fmt.Println(file.Name())
		}
	}
}
