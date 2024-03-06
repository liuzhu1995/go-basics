package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"lat"`
	Long coordinate `json:"long"`
}
type coordinate struct {
	d, m, s float64
	h       rune
}

type DMS struct {
	DD  float64 `json:"decimal"`
	DMS string  `json:"dms"`
	D   float64 `json:"degrees"`
	M   float64 `json:"minutes"`
	S   float64 `json:"seconds"`
	H   string  `json:"hemisphere"`
}

func (c coordinate) String() string {
	// 对dms格式进行格式化
	return fmt.Sprintf("Elysium Planitia is at %v°%v'%v\"%c", c.d, c.m, c.s, c.h)
}

func (c coordinate) decimal() float64 {
	// 将dms转换为十进制
	sign := 1.0

	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v\n", l.Lat, l.Long)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(DMS{
		DD:  c.decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}
func main() {

	elysium := location{
		Name: "Elysium Planitia",
		Lat:  coordinate{d: 4.0, m: 35, s: 22.2, h: 'S'},
		Long: coordinate{137, 26, 30.1, 'E'},
	}

	bytes, err := json.Marshal(elysium)

	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(bytes))

}
