package main

import "fmt"

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna", "Salacia", "Quaoar", "Sedna")

	fmt.Printf("dwarfs1: cap(%v), len(%v)\n", cap(dwarfs1), len(dwarfs1))
	fmt.Printf("dwarfs2: cap(%v), len(%v)\n", cap(dwarfs2), len(dwarfs2))
	fmt.Printf("dwarfs3: cap(%v), len(%v)\n", cap(dwarfs3), len(dwarfs3))
}
