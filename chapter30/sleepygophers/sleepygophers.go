package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}

	time.Sleep(4 * time.Second)
	go func() {
		fmt.Println("end")
	}()
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore...")
}
