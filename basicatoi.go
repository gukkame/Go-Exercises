package piscine

import "strconv"

func BasicAtoi(s string) int {
	intVar, err := strconv.Atoi(s)
	return intVar
}
