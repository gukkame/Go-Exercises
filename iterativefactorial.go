package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 && nb < 25 {
		var fact int = 1
		for i := 1; i <= nb; i++ {
			fact = fact * i
		}
		return fact
	} else {
		return 0
	}
}
