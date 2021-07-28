package main

import "github.com/01-edu/z01"

func main() {

	var alph string = "abcdefghijklmnopqrstuvwxyz"

	for i := alph[0]; i < alph[25]; i++ {

		z01.PrintRune(rune(i))

	}
	z01.PrintRune('\n')
}
