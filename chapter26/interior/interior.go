package main

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

type character struct {
	name  string
	stats stats
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

func main() {
	player := character{name: "熊大"}
	levelUp(&player.stats)

	fmt.Println(player)
	fmt.Printf("%+v\n", player)
}
