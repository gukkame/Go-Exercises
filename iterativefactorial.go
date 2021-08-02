package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	} else {
		var fact int = 1
		for i := 1; i <= nb; i++ {
			fact = fact * i
		}
		return fact
	}
}
