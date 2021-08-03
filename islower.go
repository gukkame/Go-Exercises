package piscine

func IsLower(s string) bool {
	c := 0
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			c++
		}
	}
	return c == len(s)
}
