package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 8, 8

var (
	ErrorBounds = errors.New("超出边界")
	ErrorDigit  = errors.New("无效的整数")
)

type Grid [rows][columns]int8
type SudokuError []error

func (se SudokuError) Error() string {
	var s []string

	for _, err := range se {

		s = append(s, err.Error())
	}

	return strings.Join(s, ", ")
}

func (g *Grid) set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrorBounds)
	}

	if !validDigit(digit) {
		errs = append(errs, ErrorDigit)
	}

	if len(errs) > 0 {
		return errs
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
	var s Grid
	defer fmt.Println("1")
	fmt.Println("2")

	err := s.set(7, 9, 10)
	fmt.Println(err, "err")

	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))

			for _, e := range errs {
				fmt.Printf("- %v\n", e)

				if e == ErrorBounds {
					fmt.Println("ErrorBounds")
				}

				if e == ErrorDigit {
					fmt.Println("ErrorDigit")
				}
			}
		}
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}
