package main

import (
	"fmt"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name)

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "不要只是检查错误，要优雅地处理它们")
	return err
}
func main() {
	err := proverbs("proverbs.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
