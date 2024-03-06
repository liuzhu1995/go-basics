package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {

	if !inBounds(row, column) {
		return errors.New("超出边界")
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
func main() {
	var g Grid

	err := g.Set(18, 0, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(g)
}
