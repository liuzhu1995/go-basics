package main

import (
	"fmt"
	"strings"
)

type laser int
type martian struct{}

func main() {
	var t interface {
		talk() string
	}

	t = martian{}

	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}

func (m martian) talk() string {
	return "nack nack"
}

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}
