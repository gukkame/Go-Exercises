package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for i := s[0]; i <= s[11]; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
