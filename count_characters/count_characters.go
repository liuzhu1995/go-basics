package main

import "fmt"

func main() {
	s1 := "asSASA ddd dsjkdsjs dk"
	s2 := "asSASA ddd dsjkdsjsこん dk"

	count1 := 0
	count2 := 0
	fmt.Printf("s1.length = %v, s2.length=%v\n", len(s1), len(s2))

	for i := 0; i < len(s1); i++ {
		if s1[i] >= 'a' || s1[i] <= 'z' && s1[i] >= 'A' || s1[i] <= 'Z' {
			count1++
		}
	}
	for i := 0; i < len(s2); i++ {
		if s2[i] >= 'a' || s2[i] <= 'z' && s2[i] >= 'A' || s2[i] <= 'Z' {
			count2++
		}
	}

	fmt.Printf("s1 runes: %v, s2 runes: %v\n", count1, count2)
}
