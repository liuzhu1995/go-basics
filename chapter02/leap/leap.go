package main

import "fmt"

func main() {

	var year = 2100

	var leap = (year%4 == 0 && year%100 != 0) || year%400 == 0
	if leap {
		fmt.Println("是闰年")
	} else {
		fmt.Println("不是闰年")
	}
}
