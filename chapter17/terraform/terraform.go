package main

import (
	"fmt"
	"strings"
)

type Planets []string

type Worlds string

func main() {
	var planets Planets = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planets.terraform("??? ")

	fmt.Println(planets)

	var m Worlds = "Mercury"
	fmt.Println(m.splice())
	fmt.Println(&planets[0])

	var intP *Worlds = &m
	fmt.Printf("变量m的值是 %s, 内存地址是%v\n", m, intP)
}

func (planets Planets) terraform(s string) {
	fmt.Println(planets, "planets--")
	for i, v := range planets {
		planets[i] = s + v
	}
}

func (w Worlds) splice() []string {
	fmt.Println(w, "w")
	return strings.Split(string(w), "")
}
