package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}

	timeout := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopherID", gopherID)
		case <-timeout:
			fmt.Println("超时")
			return
		}
	}
	time.Sleep(3 * time.Second)
}

func sleepGopher(id int, c chan int) {
	// time.Sleep(2 * time.Second)
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}
