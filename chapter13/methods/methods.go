package main

import "fmt"

// 编写一个程序，使它带有 celsius、fahrenheit 和 kelvin 这 3 种类型以及在这
// 3 种温度类型之间互相转换的方法

type celsius float64
type fahrenheit float64
type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func main() {
	var k kelvin = 0
	fmt.Printf("0ºK is %.2f°F\n", k.fahrenheit())
}
