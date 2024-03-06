package main

import "fmt"

func main() {
	const distance = 56_000_000
	const hoursPerDay = 24

	fmt.Println(distance/hoursPerDay/28, "days")
}
