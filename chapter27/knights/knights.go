package main

import "fmt"

// 亚瑟被一位骑士挡住了去路。正如 leftHand *item 变量的值 nil 所示，这位英雄
// 手上正空无一物。请实现一个拥有 pickup(i *item)和 give(to *character)等方法
// 的 character 结构，然后使用你在本章学到的知识编写一个脚本，使得亚瑟可以拿起一件
// 物品并将其交给骑士，与此同时为每个动作打印出适当的描述
type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v拿起了%v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {

	if c == nil || to == nil {
		return
	}

	if c.leftHand == nil {
		fmt.Printf("%v没有什么东西可给的\n", c.name)
		return
	}

	if to.leftHand != nil {
		fmt.Printf("%v手上已经有东西了\n", to.name)
		return
	}

	fmt.Printf("%v将武器%v交给了%v\n", c.name, c.leftHand.name, to.name)
	to.leftHand = c.leftHand
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v手上什么都没有拿", c.name)
	}
	return fmt.Sprintf("%v手上拿着%v", c.name, c.leftHand.name)
}

func main() {
	sword := &item{name: "剑"}
	arthur := &character{name: "亚瑟"}
	knight := &character{name: "骑士"}
	arthur.pickup(sword)

	arthur.give(knight)

	bow := &item{name: "长弓"}

	arthur.pickup(bow)
	arthur.give(knight)

	fmt.Println(arthur)
	fmt.Println(knight)
}
