package main

import "fmt"

func main() {
	q := 798
	doSomething(q)
}

func doSomething(a int) {
	b := &a

	fmt.Println(b)
}
