package main

import "fmt"

func main() {
	users := [...]string{"Bret", "Antonette", "Samantha", "Karianne"}

	fmt.Println(len(users))
	users02 := [3][4]string{{"Bret", "Antonette", "Samantha", "Karianne"}, {"Bret", "Antonette", "Samantha", "Karianne"}, {"Bret", "Antonette", "Samantha", "Karianne"}}

	fmt.Println(users02)

}
