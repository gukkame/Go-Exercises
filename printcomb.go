package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for y := 0; y < 10; y++ {
				if i < j && j < y && i < y {
					// prints
					z01.PrintRune(rune(i + 48))
					z01.PrintRune(rune(j + 48))
					z01.PrintRune(rune(y + 48))

					if !(7 == i && 8 == j && 9 == y) {
						// comma
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
