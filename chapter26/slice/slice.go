package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:3]
}

// func reclassify(planets []string) {
// 	planets = planets[0:3]
// }

func add(s *[3]int) {
	s[0] = 798
}

type location struct {
	lat, long float64
}
type person struct {
	name string
	age  int
	location
}

func setName(p *person) {
	p.name = "熊大"
}
func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}

	reclassify(&planets)

	fmt.Printf("%v cap(%v) len(%v)\n", planets, cap(planets), len(planets))

	s := [3]int{1, 2, 3}

	add(&s)
	fmt.Println(s)

	p := person{
		name:     "张三",
		age:      27,
		location: location{-1.25, 6.54},
	}

	fmt.Printf("%+v\n", p)

	setName(&p)

	fmt.Println(p.name)
	fmt.Printf("%+v\n", p)
}
