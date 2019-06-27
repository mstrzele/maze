package main

import (
	"fmt"
)

func main() {
	m := newMaze([][]value{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 1, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	sp := point{4, 0}

	fmt.Println("Maze:")
	fmt.Println(m)
	fmt.Println("Starting point:", sp)

	spt := path{}
	m.walk(sp, spt)
	for i, pt := range m.paths {
		if i == 0 || len(spt) > len(pt) {
			spt = pt
		}
	}

	fmt.Println("---")

	fmt.Println("Shortest paths:")
	for _, pt := range m.paths {
		if len(pt) == len(spt) {
			fmt.Println("Path:", pt)
			fmt.Println("Maze:")
			fmt.Println(m.mark(pt))
		}
	}
	fmt.Println("Shotest path length:", len(spt))
}
