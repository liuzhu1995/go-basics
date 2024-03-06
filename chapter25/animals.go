package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 你的任务是为火星上的第一个动物避难所创建模拟器。请创建几种不同类型的动物，为
// 每种动物命名并通过实现 Stringer 接口来返回这些名字。
// 每种动物都应该拥有相应的移动方法 move 和进食方法 eat，其中前者用于返回动物的
// 行动描述，而后者则用于随机返回动物喜欢的某种食物的名字。
// 请实现一个昼夜循环，并用它模拟 3 个 24 小时（共 72 小时）的火星日。所有动物从日
// 出开始活动，日落之后开始睡觉。在白天的时候，每过 1 小时随机挑选一种动物随机完成一
// 种动作（移动或进食），并在动作完成之后打印出相应的描述说明动物做了什么。
// 请在实现这个模拟器的时候使用结构和接口

type honeyBee struct {
	name string
}

func (h honeyBee) String() string {
	return h.name
}

func (h honeyBee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "嗡嗡作响"
	default:
		return "飞向远方采蜜"
	}
}

func (h honeyBee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "花粉"
	default:
		return "花蜜"
	}
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v\n", a, a.move())
	default:
		fmt.Printf("%v eats the %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	rand.Seed(time.Now().UnixNano())

	animal := []animal{
		honeyBee{name: "小蜜蜂"},
	}

	var sol, hour int
	for {
		fmt.Printf("%2d:00 ", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("动物们在睡觉")
		} else {
			i := rand.Intn(len(animal))
			step(animal[i])
		}
		time.Sleep(500 * time.Millisecond)
		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
	// fmt.Printf("%d\n", runtime.MemStats.Alloc/1024)
	// 此处代码在 Go 1.5.1下不再有效，更正为
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)

}
