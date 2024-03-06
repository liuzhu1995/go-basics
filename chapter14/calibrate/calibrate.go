package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offsset kelvin) sensor {
	return func() kelvin {
		return s() + offsset
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	var k kelvin = 5
	sensor := calibrate(realSensor, k)

	fmt.Println(sensor())

	sensor = calibrate(fakeSensor, k)
	fmt.Println(sensor())
	sensor = calibrate(fakeSensor, k)
	fmt.Println(sensor())

}
