package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	total := 20.0
	count := 0.0

	for count < total {
		switch num := rand.Intn(3); num {
		case 0:
			count += 0.05

		case 1:
			count += 0.1
		case 2:
			count += 0.25
		}

		// fmt.Printf("%5.2f\n", count)

	}

	var a int = 1
	b := "dd"
	c := true
	var d int32 = 1
	var e int64 = 1

	fmt.Printf("%T %T %T %T %T\n", a, b, c, d, e)

	fmt.Println(math.MaxInt)
}
