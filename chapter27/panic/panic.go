package main

import "fmt"

type talker interface {
	talk()
}

type man struct {
	name string
}

func (m man) talk() {
	fmt.Println("Hello my name is ", m.name)
}

func main() {

	var nowhere *int

	fmt.Println(nowhere == nil)

	if nowhere != nil {
		*nowhere = 233
		fmt.Println(*nowhere)
	}

	m := man{name: "张三"}
	m.talk()

	var m1 *man = &man{name: "刘大"}
	fmt.Println(m1 == nil)
	if m1 != nil {
		*m1 = man{name: "李四"}
		fmt.Printf("%+v\n", *m1)
	}

}
