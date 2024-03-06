package main

import (
	"fmt"
	"time"
)

func main() {
	// go sleepGopher()
	// time.Sleep(4 * time.Second)

	go running()

	var input string
	fmt.Scanln(&input)
}

func sleepGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}

func running() {
	var times int

	for {
		times++
		fmt.Println("tick", times)

		time.Sleep(time.Second)
	}
}
