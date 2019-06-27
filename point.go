package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func (p point) left() point {
	return point{
		x: p.x,
		y: p.y - 1,
	}
}

func (p point) right() point {
	return point{
		x: p.x,
		y: p.y + 1,
	}
}

func (p point) up() point {
	return point{
		x: p.x - 1,
		y: p.y,
	}
}

func (p point) down() point {
	return point{
		x: p.x + 1,
		y: p.y,
	}
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}
