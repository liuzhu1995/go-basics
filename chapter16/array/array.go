package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// var arr01 [3]string
	// arr01[0] = "张三"
	// arr01[1] = "李四"
	// arr01[2] = "赵五"
	// i := 3
	// arr01[i] = "798"
	// fmt.Println(arr01)

	message := "¿Cómo estás?"
	fmt.Println(len(message))

	fmt.Println(utf8.RuneCountInString(message))

	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]

	fmt.Printf("dwarfArrayType:%T, dwarfSliceType: %T\n", dwarfArray, dwarfSlice)
}
