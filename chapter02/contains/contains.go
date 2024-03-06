package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("你来到了海边")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("你离开了海边:", exit)
}
