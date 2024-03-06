package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}
	time.Sleep(4 * time.Second)
	fmt.Println("798")
}

func sleepGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore...")
	c <- id
}
