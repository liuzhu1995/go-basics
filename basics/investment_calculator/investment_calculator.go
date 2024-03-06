package main

import (
	"fmt"
	"math"
)

func main() {
	//通货膨胀率
	const inflationRate = 2.5
	// 用户输入的投资金额
	var investmentAmount float64
	// 投资年数
	years := 10.0
	// 预期的年回报率
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	// 读取用户输入并将其存储在 investmentAmount 变量中
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("ExpectedReturn Rate: ")
	fmt.Scan(&expectedReturnRate)



	// 计算投资在 years 年后的未来价值。这是通过复利公式 (本金 * (1 + 回报率)^年数) 计算的
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// 虑通货膨胀后的实际价值。这是通过将未来价值除以 (1 + 通货膨胀率)^年数 来计算的
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	formattedFV :=  fmt.Sprintf("Futuew Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Futuew Real Value: %.1f\n", futureRealValue)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

}