package piscine

func ToUpper(s string) string {
	var newString string

	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			newString = newString + string(letter-32)
		} else {
			newString = newString + string(letter)
		}
	}
	return newString
}
