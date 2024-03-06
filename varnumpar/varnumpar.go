package main

import "fmt"

func main() {

	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	greeting("New ", "efe", "fwfwe")
	greeting("New ", planets...)
	fmt.Println(planets)
}

func greeting(prefix string, list ...string) {
	for i, v := range list {
		list[i] = prefix + v
	}

}
