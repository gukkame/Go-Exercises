package piscine

func IsNumeric(s string) bool {
	c := 0
	for _, letter := range s {
		if letter > 47 && letter < 58 {
			c++
		}
	}
	return c == len(s)
}
