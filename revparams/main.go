package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for arg := len(args); arg > 0; arg-- {

		for _, r := range string(arg) { // print character
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
