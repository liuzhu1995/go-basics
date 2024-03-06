package main

import (
	"fmt"
	"time"
)

//	func stardate(t time.Time) float64 {
//		day := float64(t.YearDay())
//		h := float64(t.Hour()) / 24.0
//		return 1000 + day + h
//	}
type stardater interface {
	YearDay() int
	Hour() int
}

func stardate(t stardater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour())
	return 100 + day + h
}

func main() {
	day := time.Date(2023, 9, 12, 9, 51, 0, 0, time.UTC)

	fmt.Println(stardate(day))
}
