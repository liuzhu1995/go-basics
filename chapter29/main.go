package main

import (
	"errors"
	"fmt"
	"os"
)

// 数独是一个发生在 9 × 9 网格上的逻辑谜题，整个网格会被分割成 9 个相邻的 3 × 3 子网格。
// 网格中的每个方格可以容纳一个介于1至9之间的数字，数字 0 则表示空白。在方格中放置数字需要满足
// 特定的约束条件。具体来说，每个被放置的数字必须符合以下规则。
//  它没有在同一行中出现过。
//  它没有在同一列中出现过
//  它没有在同一个子网格中出现过。
// 请使用固定大小的 9 × 9 数组来表示数独网格，并在函数或者方法需要修改数组的时候传递指向数组的指针而不是数组的副本。
// 请实现一个方法，它能够将一个数字放置到指定位置的方格里面，并且在该数字不符合上述任一规则时返回相应的错误。

// 除此之外，请实现一个方法用于清除指定方格中的数字。考虑到被清除方格附近可能会
// 有其他空白方格（值为 0），所以这个方法在执行操作的时候不需要遵守上面提到的规则。
// 数独谜题在开始的时候通常会有一些预设的数字。请编写一个创建数独谜题的构造函
// 数，并在函数中通过复合字面量为数独谜题设置初始值，就像这样
// [rows][columns]int8{
// 	{5, 3, 0,   0, 7, 0,   0, 0, 0},  0,3 0,4 0,5
// 	{6, 0, 0,   1, 9, 5,   0, 0, 0},  1,3 1,4 1,5
// 	{0, 9, 8,   0, 0, 0,   0, 6, 0},  2,3 2,4 2,5

// 	{8, 0, 0,   0, 6, 0,   0, 0, 3},
// 	{4, 0, 0,   8, 0, 3,   0, 0, 1},
// 	{7, 0, 0,   0, 2, 0,   0, 0, 6},

// 	{0, 6, 0,   0, 0, 0,   2, 8, 0},
// 	{0, 0, 0,   4, 1, 9,   0, 0, 5},
// 	{0, 0, 0,   0, 8, 0,   0, 7, 9},
// }
// 预设的数字无法移动、修改或者清除。请修改程序，使得它可以区分无法修改的预设数
// 字和人为放置的数字，并添加一条检查规则，使得程序可以在用户尝试设置或者清除任何预
// 设数字时返回相应的错误。至于那些被初始化为 0 的数字则可以随意设置、修改和清除。
// 你不需要为这个练习编写数独解题器，但请通过测试确保上述所有规则都能够正确实现

const (
	rows, columns = 9, 9
	empty         = 0
)

// 可能发生的错误
var (
	ErrBounds     = errors.New("超出数组边界")
	ErrDigit      = errors.New("无效的数字")
	ErrInRow      = errors.New("该行中已存在此数字")
	ErrColumn     = errors.New("该列中已存在此数字")
	ErrInRegion   = errors.New("该字网格已存在此数字")
	ErrFixedDigit = errors.New("无法修改预设数字")
)

type Cell struct {
	digit int8
	fixed bool
}
type Grid [rows][columns]Cell

func NewSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]
			// 初始化为0的数字才能修改
			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}
		}
	}
	return &grid
}

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}

	if !validDigit(digit) {
		return ErrDigit
	}

	if g.inRow(row, digit) {
		return ErrInRow
	}

	if g.inColmn(column, digit) {
		return ErrColumn
	}

	if g.inRegion(row, column, digit) {
		return ErrInRegion
	}
	return nil
}

func (g *Grid) Clear(row, column int) error {
	if !inBounds(row, column) {
		return ErrBounds
	}

	if !g.isFixed(row, column) {
		return ErrFixedDigit
	}
	g[row][column].digit = empty
	return nil
}

func (g *Grid) inRow(row int, digit int8) bool {
	for _, v := range g[row] {
		if v.digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inColmn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}

	return false
}

func (g *Grid) inRegion(row, column int, digit int8) bool {
	startRow, startColumn := row-row%3, column-column%3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
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
	return digit >= 1 || digit <= 9
}

func main() {
	grid := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},

		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},

		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := grid.Set(0, 2, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("错误")

}
