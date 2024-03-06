package main

import (
	"fmt"
)

// 请编写一个维吉尼亚解密程序，并解密图 11-2 中展示的加密文本。为简单起见，我们
// 暂时只考虑加密文本和关键字都是大写英文字母的情形。
// cipherText := "CSOITEUIWUIZNSROCNKFD"
// keyword := "GOLANG"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	keyIndex := 0
	message := ""

	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		// fmt.Printf("%c %c\n", c, k)
		c = (c-k+26)%26 + 'A'
		message += string(c)

		keyIndex++
		keyIndex %= len(keyword)
	}

	fmt.Println(message)

	for i := 0; i < len(cipherText); i++ {

		c := cipherText[i]

		c = c - 'A'

		fmt.Printf("%v", c)
	}
}
