package main

import "fmt"

func main() {
	var admin *string

	scolese := "张三"
	admin = &scolese
	fmt.Printf("%v %v\n", admin, *admin)

	bolden := "李四"
	admin = &bolden
	fmt.Printf("%v %v\n", admin, *admin)

	bolden = "王麻子"
	fmt.Printf("%v %v\n", admin, *admin)

	*admin = "赵二"

	fmt.Printf("%v %v\n", bolden, *admin)

	major := admin

	fmt.Println(major == admin)
}
