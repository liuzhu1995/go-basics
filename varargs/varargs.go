package main

import "fmt"

func main() {
	typecheck(798, "hello", false)
}

func typecheck(values ...interface{}) {
	for _, value := range values {

		switch value.(type) {
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float64")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("未知类型")

		}

	}
}
