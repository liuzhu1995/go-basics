package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

var ErrorBounds = errors.New("超出边界")
var ErrorDigit = errors.New("无效的整数")

type Grid [rows][columns]int8

func (g *Grid) set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrorBounds
	}

	d := int(digit)

	if d == 0 {
		return ErrorDigit
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if column < 0 || column >= columns {
		return false
	}

	return true
}

func validDigit(digit int8) bool {
	return digit > 0 && digit <= 9
}

func main() {
	var a Grid

	err := a.set(8, 8, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(a)
}
