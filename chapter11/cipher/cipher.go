package main

import (
	"fmt"
	"strings"
)

// 为了发送加密消息，请编写一个使用关键字对纯文本消息进行加密的维吉尼亚加密程序：

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"
	keyIndex := 0

	cipherText := ""

	plainText = strings.ToUpper(strings.ReplaceAll(plainText, " ", ""))

	for i := 0; i < len(plainText); i++ {

		c := plainText[i]

		if c > 'A' && c < 'Z' {
			c -= 'A'

			k := keyword[keyIndex] - 'A'

			c = (c+k)%26 + 'A'

			keyIndex++
			keyIndex %= len(keyword)
		}

		cipherText += string(c)
	}

	fmt.Println(cipherText)

}
