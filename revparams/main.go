package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args                       //
	for _, arg := range args[len(args)] { // words
		s := string([]rune{arg})
		for _, r := range s { // print character
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
