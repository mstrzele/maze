package main

import (
	"fmt"
	"strings"
)

type maze struct {
	values [][]value
	paths  []path
}

func newMaze(v [][]value) maze {
	return maze{values: v}
}

func (m maze) validate(p point) value {
	switch {
	case
		p.x < 0,
		p.x >= len(m.values),
		p.y < 0,
		p.y >= len(m.values[p.x]):
		return -1
	}

	return m.values[p.x][p.y]
}

func (m *maze) walk(p point, pt path) {
	npt := make(path, len(pt))
	copy(npt, pt)

	switch m.validate(p) {
	case openSpace:
		npt = append(npt, p)

		nps := map[string]point{
			"left":  p.left(),
			"right": p.right(),
			"up":    p.up(),
			"down":  p.down(),
		}

		for _, np := range nps {
			if !pt.validated(np) {
				m.walk(np, npt)
			}
		}
	case endPoint:
		npt = append(npt, p)
		m.paths = append(m.paths, npt)
	}
}

func (m maze) mark(pt path) maze {
	v := make([][]value, len(m.values))
	for x := 0; x < len(m.values); x++ {
		v[x] = make([]value, len(m.values[x]))
		for y := 0; y < len(m.values[x]); y++ {
			v[x][y] = m.values[x][y]
		}
	}

	for _, p := range pt {
		v[p.x][p.y] = 3
	}

	return newMaze(v)
}

func (m maze) String() string {
	var sb strings.Builder
	for x := 0; x < len(m.values); x++ {
		if x != 0 {
			sb.WriteString("\n")
		}
		for y := 0; y < len(m.values[x]); y++ {
			switch m.values[x][y] {
			case openSpace, wall, endPoint:
				sb.WriteString(fmt.Sprintf("%d ", m.values[x][y]))
			default:
				sb.WriteString(fmt.Sprintf("x "))
			}
		}
	}
	return sb.String()
}
