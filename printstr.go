package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	st := []byte(s)
	for i := 0; i <= len(st); i++ {
		z01.PrintRune(rune(st[i]))
	}
	z01.PrintRune('\n')
}
