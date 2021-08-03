package piscine

func TrimAtoi(s string) int {
	i := 0
	for _, letter := range s {
		if letter >= '0' && letter <= '9' {
			if letter == '0' {
				i += 0
				i *= 10
			} else if letter == '1' {
				i += 1
				i *= 10
			} else if letter == '2' {
				i += 2
				i *= 10
			} else if letter == '3' {
				i += 3
				i *= 10
			} else if letter == '4' {
				i += 4
				i *= 10
			} else if letter == '5' {
				i += 5
				i *= 10
			} else if letter == '6' {
				i += 6
				i *= 10
			} else if letter == '7' {
				i += 7
				i *= 10
			} else if letter == '8' {
				i += 8
				i *= 10
			} else if letter == '9' {
				i = i + 9
				i = i * 10
			}
		}
	}
	num := i / 10
	mind := 100
	lin := 0
	for index, letter := range s {
		if letter == '-' {
			mind = index
		}
		if letter == '1' || letter == '2' || letter == '3' || letter == '4' || letter == '5' || letter == '6' || letter == '7' || letter == '8' || letter == '9' {
			lin = index
			if index > mind {
				return -num
			} else if lin > 0 {
				return num
			}
		}
	}
	return num
}
