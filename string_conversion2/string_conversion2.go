package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	// s := "798"

	// a, err := strconv.Atoi(s)

	// if err != nil {
	// 	fmt.Println("变量s的值不能转换为数字")
	// 	os.Exit(1)
	// }

	// fmt.Println(a)
	// openTxt()
	t, e := mySqrt(-1)
	fmt.Println(t, e)
}

func openTxt() {
	file, err := os.Open("string_conversion2/abc.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		fmt.Println(err, "err")
		if err != io.EOF {
			break
		}

		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	} // error case
	return math.Sqrt(f), true
}
