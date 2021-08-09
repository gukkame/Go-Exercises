package main

import (
	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) int {
	if nbr == 2 {
		return 1
	} else {
		return 0
	}
}

func main() {
	if isEven(a) == 0 {
		printStr("EvenMsg")
	} else {
		printStr("OddMsg")
	}
}
