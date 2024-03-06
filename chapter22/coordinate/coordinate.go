package main

import "fmt"

// 在 DMS 格式（degrees/minutes/seconds，度/分/秒）中，分和秒表示的是位置而不是时间，
// 其中每 60 秒（"）为一分，每 60 分（'）为一度。例如，布莱德柏利着陆点的位置用 DMS
// 格式表示就是南纬 4°35' 22.2"东经 137°26' 30.1"

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func main() {
	lat := coordinate{d: 4.0, m: 35, s: 22.2, h: 'S'}
	long := coordinate{137, 26, 30.1, 'E'}

	fmt.Println(lat.decimal())
	fmt.Println(long.decimal())

	cur := newLocation(lat, long)
	fmt.Println(cur)
}

func (c coordinate) decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat: lat.decimal(), long: long.decimal()}
}
