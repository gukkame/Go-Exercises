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

func main() {
	points := &point{}

	setPoint(points)
	p := points.x / 10
	o := points.x % 10
	p1 := points.y / 10
	o1 := points.y % 10
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(rune(p) + 48)
	z01.PrintRune(rune(o) + 48)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(rune(p1) + 48)
	z01.PrintRune(rune(o1) + 48)
	z01.PrintRune(10)
}
