package piscine

func NRune(s string, n int) rune {
	c := 0
	for range s {
		c++
	}

	if n > 1 && c+1 > n {
		k := ([]rune(s))[n-1]

		return k

	} else {
		return 0
	}
}
