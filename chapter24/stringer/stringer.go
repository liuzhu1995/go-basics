package main

import "fmt"

type location struct {
	lat, long coordinate
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {

	cur := location{lat: coordinate{d: 4.0, m: 35, s: 22.2, h: 'S'}, long: coordinate{137, 26, 30.1, 'E'}}

	fmt.Println(cur)
}

func (c coordinate) String() string {
	return fmt.Sprintf("Elysium Planitia is at %v°%v'%v\"%c", c.d, c.m, c.s, c.h)
}
