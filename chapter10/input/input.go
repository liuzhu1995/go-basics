package main

import "fmt"

func main() {
	launchText := "true"

	var launch bool

	switch launchText {
	case "true", "yes", "1":
		launch = true
	case "false", "no", "0":
		launch = false
	default:
		fmt.Println("error")
	}
	fmt.Printf("%v %[1]T\n", launch)
}
