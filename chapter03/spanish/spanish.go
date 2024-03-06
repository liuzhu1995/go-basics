package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "bytes")

	for i := 0; i < len(question); i++ {
		c := question[i]
		fmt.Printf("%c", c)
	}

	fmt.Println()

	for _, v := range question {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	for i := 0; i < utf8.RuneCountInString(question); i++ {
		c := question[i]
		fmt.Printf("%c", c)
	}
}
