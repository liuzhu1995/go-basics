package main

import "fmt"

type sol int
type celsius float64

type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}
type report struct {
	sol
	temperature
	location
}

func main() {
	t := temperature{high: -1.0, low: -78.0}

	report := report{sol: 15, temperature: t}

	fmt.Println(report.average())
	fmt.Println(report.high)
	report.high = 32
	fmt.Println(report.high)
	report.days(100)
}
func (t temperature) average() celsius {
	return celsius((t.high + t.low) / 2)
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	return 5
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}
