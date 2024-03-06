package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	p := person{
		name:       "Jason",
		superpower: "",
		age:        10,
	}
	birthday(&p)

	fmt.Printf("%+v\n", p)
}
