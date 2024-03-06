package main

import "fmt"

func main() {

	const (
		hoursPerDay   = 24
		minutePerHour = 60
	)

	var (
		speed    = 100800 // km/h
		distance = 96300000
	)
	fmt.Println(distance/speed/hoursPerDay, "days")

}
