package main

import "fmt"

func main() {
	// 收入、费用、税率
	var revenue, expenses, taxRate float64
	// EBT 税前收益，profit 税后收益，ratio(EBT / profit)
	var ebt,profit, ratio float64


	fmt.Print("请输入收入：")
	if _,err := fmt.Scan(&revenue); err!= nil {
		fmt.Println("输入错误：", err)
		return 
	}
 
	fmt.Print("请输入费用：")
	if _,err := fmt.Scan(&expenses); err!= nil {
		fmt.Println("输入错误：", err)
		return 
	}
	fmt.Print("请输入税率：")
	if _,err := fmt.Scan(&taxRate); err!= nil {
		fmt.Println("输入错误：", err)
		return 
	}

	ebt = revenue - expenses

	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit

	fmt.Printf("税前收益(EBT): %.2f\n", ebt)
	fmt.Printf("税后收益:(profit)：%.2f\n", profit)
	fmt.Printf("比率(EBT/profit) ：%.2f\n", ratio)
}