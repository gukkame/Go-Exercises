package piscine

func LastRune(s string) rune {
	k := s[len(s)-1:]
	r := rune(k[0])

	return r
}
