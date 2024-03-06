package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%#v", e)
		}
	}()

	url, err := url.Parse("https://a b.com/")

	if err != nil {
		fmt.Println(err)

		// if e, ok := err.(*url.Error); ok {

		// 	fmt.Println(e.Op)
		// }
		os.Exit(1)
	}

	fmt.Println(url)
}
