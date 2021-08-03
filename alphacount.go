package piscine

func AlphaCount(s string) int {
	c := 0
	for _, letter := range s {
		if !(letter == ' ') {
			if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
				c++
			}
		}
	}
	return c
}
