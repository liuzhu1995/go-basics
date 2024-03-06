package main

import "fmt"

func main() {

	message := "uv vagreangvbany fcnpr fgngvba"

	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c = c + 13

			if c > 'z' {
				c -= 26
			}
		}

		fmt.Printf("%c", c)

	}
}
