package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
0600：所有者有读写权限，组和其他用户无权限
0644  所有者有读写权限，组和其他用户只有读权限
0660：所有者有读写权限，组有读写权限，其他用户无权限
0666：所有者、组和其他用户都有读写权限（但通常不推荐这样设置，因为它可能会带来安全风险）
0700：所有者有读写执行权限，组和其他用户无权限
0755：所有者有读写执行权限，组有读执行权限，其他用户有读执行权限（这是很多可执行文件的默认权限设置）
*/

const accountBalanceFile = "writing_files/balance.txt"
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	if (err != nil) {
		log.Fatal(err)
	}
}

func getBalanceFromFile() float64 {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		log.Fatal(err)
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		log.Fatal(err)
	}
	return balance
}
func main() {
	// var accountBalance = 1000.0

	// writeBalanceToFile(accountBalance)

	balance := getBalanceFromFile()
	fmt.Println(balance)
}