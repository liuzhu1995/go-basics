package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown := 10

	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."

	fmt.Println(str)
	fmt.Printf("%v %[1]T\n", strconv.Itoa(countdown))

	countdown, err := strconv.Atoi("10.66")

	if err != nil {
		fmt.Println("错误")
	}

	fmt.Println(countdown)

	number := 999.99

	fmt.Println(strconv.Itoa(int(number)))
	str02 := fmt.Sprintf("%v", number)
	fmt.Println(str02)

	fmt.Printf("%v %[1]T\n", fmt.Sprint(true)+" Hello")
}
