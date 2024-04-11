package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

const (
	msg = "x = 42, y = 21\n"
)

func main() {
	points := &point{}

	setPoint(points)

	for _, x := range msg {
		z01.PrintRune(x)
	}
}
