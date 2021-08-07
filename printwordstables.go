package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, p := range a[0:] {

		for _, let := range p {
			z01.PrintRune(let)
		}
		z01.PrintRune('\n')
	}
}
