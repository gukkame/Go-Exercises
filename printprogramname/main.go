package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	name := arguments[0]

	runes := []rune(name)
	for i := range runes {
		if i > 1 {
			z01.PrintRune(runes[i])
		}
	}
	z01.PrintRune('\n')
}
