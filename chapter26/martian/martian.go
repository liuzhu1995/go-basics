package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{})
	shout(&martian{})

	shout(martian{})
	shout(&martian{})
}
