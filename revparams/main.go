package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, arg := range args[1+len(args)] {
		s := string([]rune{arg})
		for _, r := range s {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
