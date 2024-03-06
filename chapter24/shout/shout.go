package main

import (
	"fmt"
	"strings"
)

type laser int
type martian struct{}

type starship struct{ laser }

type rover struct{}
type talker interface {
	talk() string
}

func main() {
	shout(martian{})
	shout(laser(2))

	s := starship{laser(3)}
	shout(s)

	shout(rover{})
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())

	fmt.Println(louder)
}

func (m martian) talk() string {
	return "nack nack"
}

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func (r rover) talk() string {
	return "whir whir"
}
