package main

import "fmt"

func main() {
	fn1()
	defer fn2()
	a := fn3()
	fmt.Println(a)

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func fn1() {
	fmt.Println("fn1")
}
func fn2() {
	fmt.Println("fn2")
}
func fn3() int {
	fmt.Println("fn3")
	return 798
}
