package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	fmt.Printf("p: %+v\n", p)
	if p == nil {
		return
	}
	p.age++
}

func main() {
	p := person{name: "张三", age: 27}
	p.birthday()
	fmt.Printf("%+v\n", p)

	var p1 *person
	p1.birthday()
	fmt.Printf("%+v\n", p1)
}
