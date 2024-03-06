package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	answer := 798

	address := &answer
	fmt.Println(address)
	fmt.Println(*address)
	fmt.Printf("%T\n", address)

	locations := [3]location{
		location{lat: 1.55, long: 4.151},
		location{lat: 1.55, long: 4.151},
		location{lat: 1.55, long: 4.151},
	}

	// add_local := &locations
	var add_local *[3]location = &locations
	fmt.Println(add_local)
	fmt.Printf("%T\n", add_local)

	fmt.Println(*&add_local[0].lat)
}
