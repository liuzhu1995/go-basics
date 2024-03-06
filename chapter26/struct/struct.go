package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	timmy := &person{
		name: "Jason",
		age:  10,
	}

	fmt.Println(timmy)
	fmt.Printf("%T\n", timmy)
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
}
