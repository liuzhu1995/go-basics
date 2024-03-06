package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	f, err := mySqrt(-9)

	fmt.Println(f, err)

	q := 1

	multiply(5, 6, &q)
	fmt.Println(q)
}

func mySqrt(f float64) (a float64, err error) {
	if f < 0 {
		return a, errors.New("参数不能是负数")
	}
	return math.Sqrt(f), nil
}

func multiply(a, b int, reply *int) {
	*reply = a * b
}
