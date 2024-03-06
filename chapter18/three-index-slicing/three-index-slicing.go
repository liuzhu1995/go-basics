package main

import "fmt"

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	fmt.Printf("planets: cap(%v), len(%v)\n", cap(planets), len(planets))

	terrestrial := planets[0:4:4]

	worlds := append(terrestrial, "798")
	fmt.Printf("terrestrial(%v) cap(%v), len(%v)\n", terrestrial, cap(terrestrial), len(terrestrial))

	fmt.Printf("worlds(%v) cap(%v), len(%v)\n", worlds, cap(worlds), len(worlds))

	fmt.Printf("planets(%v) cap(%v), len(%v)\n", planets, cap(planets), len(planets))

}
