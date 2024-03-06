package main

import (
	"fmt"
	"math/rand"
)

// 请编写一个“猜数字”程序，让它重复地在 1～100 中随机选择一个数字，直到这个数
// 字跟你在程序开头声明的数字相同为止。请打印出程序随机选择的每个数字，并说明该数字
// 是大于还是小于你声明的数字
func main() {
	var num = 7

	for {
		randNum := rand.Intn(99) + 1
		if randNum == num {
			fmt.Println("正确", randNum)
			break
		}

		if randNum > num {
			fmt.Println("大")
		} else {
			fmt.Println("小")
		}
	}

}
