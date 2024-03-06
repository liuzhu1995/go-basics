package main

import "fmt"

func main() {
	var soup []string = make([]string, 3, 5)

	fmt.Println(soup == nil)

	for _, v := range soup {
		fmt.Println(v)
	}

	soup[0] = "0"
	soup[1] = "1"
	soup[2] = "2"
	soup = append(soup, "zhangs", "orange")

	fmt.Println(soup)

	fmt.Println(len(soup), cap(soup))

	s1 := mirepoix(nil, "orange", "apple", "banana")

	fmt.Println(s1)
}

func mirepoix(ingredients []string, strs ...string) []string {
	return append(ingredients, strs...)
}
