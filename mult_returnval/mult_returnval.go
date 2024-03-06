package main

import "fmt"

// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值
func main() {
	fmt.Println(andProductDiff_01(1, 2))
	fmt.Println(andProductDiff_02(1, 2))
}

func andProductDiff_01(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func andProductDiff_02(a, b int) (and int, pro int, diff int) {
	and = a + b
	pro = a * b
	diff = a - b
	return
}
