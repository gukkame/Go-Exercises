package piscine

func IterativeFactorial(nb int) int {
	var fact int = 1

	if nb < 0 {
	}

	for i := 1; i <= nb; i++ {
		fact = fact * i
	}

	return fact
}
