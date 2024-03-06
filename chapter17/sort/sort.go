package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)

	s1 := strings.Join(planets, "*")

	fmt.Println(s1)
}
