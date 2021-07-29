package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string {
	for i := len(s); i >= 0; i-- {
		z01.PrintRune(rune(i))
	}
}
