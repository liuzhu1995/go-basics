package main

import (
	"fmt"
	"math"
)

type world struct{ radius float64 }
type location struct {
	name      string
	lat, long float64
}
type gps struct {
	world
	current     location
	destination location
}

type rover struct {
	gps
}

func main() {
	var mars = world{radius: 3389.5}
	currLocation := location{"Bradbury Landing", -4.5895, 137.4417}
	desLocation := location{"Elysium Planitia", 4.5, 135.9}
	gps := gps{
		world:       mars,
		current:     currLocation,
		destination: desLocation,
	}

	fmt.Println(gps.message())
	curiosity := rover{
		gps: gps,
	}

	fmt.Println(curiosity.message())
}

func (l location) descripton() string {
	return fmt.Sprintf("%v lat:%v long: %v", l.name, l.lat, l.long)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}
func rad(deg float64) float64 {
	// 将角度转换为弧度
	return deg * math.Pi / 180
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.1f km to %v km\n", g.distance(), g.destination.descripton())
}
