package main

import "fmt"

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")

	fmt.Println(twoWorlds)
}

func terraform(prefix string, worlds ...string) []string {

	newWorlds := make([]string, len(worlds))

	for i, v := range worlds {
		newWorlds[i] = prefix + v
	}

	return newWorlds
}
