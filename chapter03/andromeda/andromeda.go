package main

import (
	"fmt"
)

func main() {
	const distance = 236e15
	const lightSpeed = 299792

	const secondsPerDay = 86400
	const daysPerYear = 365
	fmt.Print(distance / lightSpeed / secondsPerDay / daysPerYear)

	message := "shalom"
	s := message[0]

	fmt.Printf("%c %[1]v %[1]T\n", s)

	for _, v := range message {
		fmt.Printf("%c %[1]v %[1]T\n", v)
	}

}
