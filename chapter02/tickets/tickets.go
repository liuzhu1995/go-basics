package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const dayPerSeconds = 86400
	const distance = 62_100_000
	company := ""
	trip := ""
	fmt.Printf("%-16v%4v%8v%-10v\n", "太空航空公司", "飞行天数", "飞行类型", "价格（百万美元）")
	for i := 0; i < 10; i++ {

		switch num := rand.Intn(3); num {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		speed := rand.Intn(15) + 16
		duration := distance / speed / dayPerSeconds
		price := 20 + speed

		if rand.Intn(2) == 1 {
			trip = "往返"
			price = price * 2
		} else {
			trip = "单程"
		}
		fmt.Printf("%-16v%8v%-10v%-10v\n", company, duration, trip, price)
	}

}
