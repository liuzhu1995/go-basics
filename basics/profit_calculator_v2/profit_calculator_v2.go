package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

/*
Goals
1) Validate user input
	显示错误消息并在提供无效输入时退出
	没有负数
	Not 0
2)  将计算结果存储到文件中
*/

func main() {
	// revenue 收入、expenses 费用、taxRate 税率
	revenue, err := getUserInput("请输入收入：")
	if err != nil {
		fmt.Println(err)
		return 
	}
	expenses, err := getUserInput("请输入费用：")
	if err != nil {
		fmt.Println(err)
		return 
	}
	taxRate, err := getUserInput("请输入税率：")
	if err != nil {
		fmt.Println(err)
		return 
	}
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("税前收益(EBT): %.2f\n", ebt)
	fmt.Printf("税后收益:(profit): %.2f\n", profit)
	fmt.Printf("比率(EBT/profit) : %.2f\n", ratio)
	storeResult(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	// EBT 税前收益，profit 税后收益，ratio(EBT / profit)
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func getUserInput(prompt string) (float64, error) {
	var userInput float64
	fmt.Print(prompt)

	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatalln(err)
	}

	if userInput <= 0 {
		return 0, errors.New("输入值必须大于0")
	}
	return userInput, nil

}

func storeResult(ebt, profit, ratio float64) {
	 result := fmt.Sprintf("EBT:%.1f\nprofit:%.1f\nratio:%.3f\n", ebt, profit, ratio)

	 err := os.WriteFile("profit_calculator_v2/result.txt", []byte(result), 0644)

	 if err != nil {
		log.Fatalln(err)
	 
	 }
}