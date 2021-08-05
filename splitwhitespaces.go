package piscine

func SplitWhiteSpaces(s string) []string {
	var newString []string
	var tempString string = ""
	length := len(s) - 1

	for i, letter := range s {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') || letter == '?' {
			tempString = tempString + string(letter)
		} else if (letter == 32 && tempString != "") || (i == length) {
			newString = append(newString, tempString)
			tempString = ""
		}
	}
	return newString
}
