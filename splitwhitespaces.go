package piscine

func SplitWhiteSpaces(s string) []string {
	var newString []string
	var tempString string = ""
	length := len(s) - 1

	for i, letter := range s {
		if letter >= 33 && letter <= 126 {
			tempString = tempString + string(letter)
		}
		if (letter == 32 && tempString != "") || (letter == 9 && tempString != "") || (letter == 10 && tempString != "") {
			newString = append(newString, tempString)
			tempString = ""
		} else if i == length {

			newString = append(newString, tempString)
			tempString = ""
		}
	}
	return newString
}
