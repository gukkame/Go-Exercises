package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i := len(args); i > 1; i++ {
		for j := 0; j < len(string(i)); j++ {
			z01.PrintRune(j)
		}

		z01.PrintRune('\n')
	}
}
