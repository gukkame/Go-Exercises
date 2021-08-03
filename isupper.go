package piscine

func IsUpper(s string) bool {
	c := 0
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			c++
		}
	}
	return c == len(s)
}
