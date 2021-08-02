package piscine

func IterativePower(nb int, power int) int {
	if nb > 0 && nb < 25 {
		var fact int = 1
		for i := 1; i <= power; i++ {

			fact1 := nb * nb
			fact = fact1 * nb

		}
		return fact
	}
	if power == 1 {
		return nb
	}
	if nb < 0 || power < 0 {
		return 0
	} else {
		return 0
	}
}
