package piscine

func ToLower(s string) string {
	var newString string

	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			newString = newString + string(letter+32)
		} else {
			newString = newString + string(letter)
		}
	}
	return newString
}
