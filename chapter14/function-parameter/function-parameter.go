package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Printf("%v K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

type getRowFn func() (string, string)

func drawTable(rows int, getRow getRowFn) {
	for i := 0; i < rows; i++ {
		a, _ := getRow()
		fmt.Println(a)
	}
}

func getRow() (string, string) {
	return "hello-", "|world"
}
func main() {
	measureTemperature(3, fakeSensor)

	drawTable(10, getRow)
}
