package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	celsius := kelvinToCelsius(k)
	fahrenheit := celsiusToFahrenheit(celsius)

	return fahrenheit
}

func main() {
	f := kelvinToFahrenheit(0)

	fmt.Printf("0ºK is %.2f°F\n", f)
}
