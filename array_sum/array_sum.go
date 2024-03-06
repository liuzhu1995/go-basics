package main

import "fmt"

func main() {

	arr1 := [5]int{7, 8, 5, 6, 3}
	r := sum(&arr1)

	fmt.Println(r)

}

func sum(list *[5]int) (sum int) {

	for _, v := range list {
		sum += v
	}
	return
}
