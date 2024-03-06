package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make(chan string)
	c := make(chan string)
	go sourceGopher(a)
	go filterGopher(a, c)
	printGopher(c)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", " a bad apple", "goodbye all"} {
		downstream <- v
	}

	close(downstream)
}

func filterGopher(upstream, downstream chan string) {

	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for str := range upstream {
		fmt.Println(str)
	}
}
