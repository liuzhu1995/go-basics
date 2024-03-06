package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string) {
	for i, v := range worlds {
		worlds[i] = strings.TrimSpace(v)
	}
}
func main() {
	planets := []string{" Venus ", "Earth ", " Mars"}
	fmt.Println(planets)
	hyperspace(planets)
	fmt.Println(planets)

	s1 := "Venus"
	s2 := strings.Split(s1, "")
	fmt.Println(s2)
	s3 := strings.Join(s2, "--")
	fmt.Println(s3)
}
