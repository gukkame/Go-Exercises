package piscine

func NRune(s string, n int) rune {
	c := 0
	for range s {
		c++
	}

	if n > 0 && c+1 > n {
		k := ([]rune(s))[n-1]
		r := rune(k)
		return r

	} else {
		return 0
	}
}
