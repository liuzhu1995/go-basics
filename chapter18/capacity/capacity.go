package main

import "fmt"

// 编写一个程序，使它通过循环持续地将元素追加至切片，并在切片的容量发生变化时
// 打印出切片的容量。请判断 append 函数在底层数组的空间被填满之后，是否总会将数组的
// 容量增加一倍？
func main() {

	worlds := make([]int, 5000)
	lastCap := cap(worlds)

	for i := 0; i < 5000; i++ {
		worlds = append(worlds, 798)
		if cap(worlds) != lastCap {
			fmt.Println(worlds)
			fmt.Println(cap(worlds))
			lastCap = cap(worlds)
		}
	}

}
