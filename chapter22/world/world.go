package main

import (
	"fmt"
	"math"
)

// 在 DMS 格式（degrees/minutes/seconds，度/分/秒）中，分和秒表示的是位置而不是时间，
// 其中每 60 秒（"）为一分，每 60 分（'）为一度。例如，布莱德柏利着陆点的位置用 DMS
// 格式表示就是南纬S 4°35' 22.2"东经E 137°26' 30.1"

// 计算出英国伦敦（北纬 51°30'，西经 0°08'）至法国巴黎（北纬 48°51'，东经 2°21'）
// 之间的距离

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}
type world struct {
	radius float64
}

func main() {
	var mars = world{radius: 3389.5}

	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)

	yonqi := newLocation(coordinate{d: 14.0, m: 34.0, s: 6.2, h: 'S'}, coordinate{d: 175, m: 28, s: 21.5, h: 'E'})
	jiyu := newLocation(coordinate{d: 1, m: 56, s: 46.3, h: 'S'}, coordinate{d: 354, m: 28, s: 24.2, h: 'E'})
	haoqi := newLocation(coordinate{d: 4, m: 35, s: 22.2, h: 'S'}, coordinate{d: 137, m: 26, s: 30.1, h: 'E'})
	dongcha := newLocation(coordinate{d: 4, m: 30, s: 0.0, h: 'S'}, coordinate{d: 135, m: 54, s: 0, h: 'E'})

	fmt.Printf("%.2f km\n", mars.distance(yonqi, jiyu))
	fmt.Printf("%.2f km\n", mars.distance(yonqi, haoqi))
	fmt.Printf("%.2f km\n", mars.distance(yonqi, dongcha))

}

func (c coordinate) decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func newLocation(lat, long coordinate) location {
	return location{lat: lat.decimal(), long: long.decimal()}
}

func rad(deg float64) float64 {
	// 将角度转换为弧度
	return deg * math.Pi / 180
}
