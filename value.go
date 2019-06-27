package main

const (
	openSpace value = 0
	wall      value = 1
	endPoint  value = 2
)

type value int

func (v value) String() string {
	switch v {
	default:
		return "out of bound"
	case 0:
		return "open space"
	case 1:
		return "wall"
	case 2:
		return "end point"
	}
}
