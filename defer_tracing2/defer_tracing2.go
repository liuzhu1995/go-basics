package main

import "fmt"

func main() {
	var arr1 = new([5]int)
	fmt.Printf("%p\n", arr1)
	a()
	b()
}

func trace(s string) string {
	fmt.Println("进入", s)
	return s
}
func un(s string) {
	fmt.Println("离开", s)
}
func a() {
	defer un(trace("a"))
}

func b() {
	defer un(trace("b"))
}
