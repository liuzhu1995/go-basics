package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}
func (s sol) Hour() int {
	return 0
}

func stardate(t stardater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour())
	return 100 + day + h
}

func main() {
	day := time.Date(2023, 9, 12, 9, 51, 0, 0, time.UTC)
	fmt.Println(stardate(day))

	s := sol(798)

	fmt.Println(stardate(s))
}
