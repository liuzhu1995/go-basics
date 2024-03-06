package main

import "fmt"

func main() {
	var arr1 [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var slice1 []int = arr1[2:5] //

	fmt.Println(len(arr1), cap(arr1))
	fmt.Println(len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 798)
	fmt.Println(len(slice1), cap(slice1), slice1)
}
