package main

import (
	"os"
)

func Itoa(n int) string {
	str := ""
	var negative string
	if n < 0 {
		negative = "-"
		n *= -1
	}
	for i := n; i > 0; i /= 10 {
		str = string(rune(i%10+48)) + str
	}
	return negative + str
}

func isNumeric(s string) bool {
	for _, i := range s {
		if (i >= 48 && i <= 57) || i == 43 || i == 45 {
			return true
		}
	}
	return false
}

func main() {
	var answer int
	if len(os.Args[1:]) != 3 || (!(isNumeric(os.Args[1])) || !(isNumeric(os.Args[3]))) {
		return
	}
	if len(os.Args[1]) >= 19 || len(os.Args[3]) >= 19 {
		return
	}
	n1 := Atoi(os.Args[1])
	n2 := Atoi(os.Args[3])
	operator := os.Args[2]
	if operator == "+" || operator == "-" || operator == "*" || operator == "/" || operator == "%" {
		if operator == "+" {
			answer = n1 + n2
		} else if operator == "-" {
			answer = n1 - n2
		} else if operator == "*" {
			answer = n1 * n2
		} else if operator == "%" {
			if n2 == 0 {
				os.Stdout.WriteString("No modulo by 0")
				os.Stdout.WriteString("\n")
				return
			} else {
				answer = n1 % n2
			}
		} else if operator == "/" {
			if n2 == 0 {
				os.Stdout.WriteString("No division by 0")
				os.Stdout.WriteString("\n")
				return
			} else {
				answer = n1 / n2
			}
		}
		if answer == 0 {
			os.Stdout.WriteString("0")
			os.Stdout.WriteString("\n")
		} else if answer > -9223372036854775808 && 9223372036854775807 > answer {
			os.Stdout.WriteString(Itoa(answer))
			os.Stdout.WriteString("\n")
		}
	}
}

func Atoi(s string) int {
	runes := []rune(s)
	numb := 0
	for _, s := range runes {
		for i := '0'; i <= '9'; i++ {
			if s == i {
				numb = numb*10 + int(i-'0')
			}
		}
	}

	if runes[0] == '-' {
		numb *= -1
	}
	return numb
}
