package piscine

func IsPrintable(s string) bool {
	c := 0
	for _, letter := range s {
		if letter >= 31 {
			c++
		}
	}
	return c == len(s)
}
