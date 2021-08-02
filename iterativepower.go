package piscine

func IterativePower(nb int, power int) int {
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	if power > 0 {
		var fact int = 1
		for i := 1; i <= power; i++ {
			fact = fact * nb
		}
		return fact
	} else {
		return 0
	}
}
