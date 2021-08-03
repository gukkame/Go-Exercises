package piscine

func IsAlpha(s string) bool {
	c := 0
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' || letter > 47 && letter < 58 {
			c++
		}
	}
	return c == len(s)
}
