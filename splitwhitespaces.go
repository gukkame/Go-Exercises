package piscine

func SplitWhiteSpaces(s string) []string {
	var newString []string
	var tempString string = ""
	length := len(s) - 1

	for i, letter := range s {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			tempString = tempString + string(letter)
		}
		if letter >= '!' && letter <= '/' {
			tempString = tempString + string(letter)
		}
		if letter >= 123 && letter <= 126 {
			tempString = tempString + string(letter)
		}
		if letter >= 48 && letter <= 57 {
			tempString = tempString + string(letter)
		}
		if letter >= 91 && letter <= 96 {
			tempString = tempString + string(letter)
		}
		if (letter == 32 && tempString != "") || (letter == 9 && tempString != "") || (letter == 10 && tempString != "") {
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
