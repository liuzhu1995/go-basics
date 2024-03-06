package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	const startDistance = 56_000_000
	const endDistance = 401_000_000
	var distance = rand.Intn(endDistance-startDistance) + startDistance

	fmt.Println(distance)
}
