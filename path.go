package main

type path []point

func (pt path) validated(p point) bool {
	for _, pp := range pt {
		if pp == p {
			return true
		}
	}
	return false
}
