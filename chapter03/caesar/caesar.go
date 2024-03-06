package main

import "fmt"

// 请编写一个凯撒密码的解密程序，它需要把大写字母和小写字母向左移动 3 位，也就是
// 把字符减去 3。除此之外，程序还需要把'a'转换为'x'、'b'转换为'y'、'c'转换为'z'，
// 并对大写字母做同样的处理
func main() {
	message := "L fdph, L vdz, L frqtxhuhg."

	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c -= 3

			if c < 'a' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}

}
