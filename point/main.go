package main

import "fmt"

type point struct {
	x int
	y int
}

func setPoint(ptr *point, x int, y int) {
	ptr.x = x
	x = 42
	ptr.y = y
	y = 21
}

func main() {
	points := point{}

	setPoint(&points, 42, 21)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
