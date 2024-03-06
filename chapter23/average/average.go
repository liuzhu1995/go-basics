package main

import "fmt"

type celsius float64

type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}
type report struct {
	sol         int
	temperature temperature
	location    location
}

func main() {
	t := temperature{high: -1.0, low: -78.0}
	fmt.Println(t.average())

	report := report{sol: 15, temperature: t}

	fmt.Println(report.temperature.average())

	fmt.Println(report.average())
}

func (t temperature) average() celsius {
	return celsius((t.high + t.low) / 2)
}

func (r report) average() celsius {
	return r.temperature.average()
}
