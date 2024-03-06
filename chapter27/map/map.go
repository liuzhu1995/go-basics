package main

import "fmt"

func main() {
	var soup map[string]int

	fmt.Println(soup == nil)

	text, ok := soup["name"]

	if ok {
		fmt.Println(text)
	}

	var v interface{}

	fmt.Printf("%T %v %v %#v\n", v, v, v == nil, v)

	var n *int

	v = n
	fmt.Printf("%T %v %v %#v\n", v, v, v == nil, v)
}
