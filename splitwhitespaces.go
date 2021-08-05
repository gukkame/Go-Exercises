package piscine

func SplitWhiteSpaces(s string) []string {
	var newString []string
	var tempString string = ""
	length := len(s) - 1

	for i, letter := range s {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			tempString = tempString + string(letter)
		} else if letter == 32 && tempString != "" {
			newString = append(newString, tempString)
			tempString = ""
		} else if i == length {
			tempString = tempString + string(letter)
			newString = append(newString, tempString)
			tempString = ""
		}
	}
	return newString
}
