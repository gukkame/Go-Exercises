package piscine

func Sqrt(nb int) int {
	upper := nb
	lower := 0
	var z int

	for z*z != nb {
		z = int((upper + lower) / 2)
		if z*z > nb {
			upper = z
		}
		if z*z < nb {
			lower = z
		}
		if (upper-lower)*(upper-lower) == 1 {
			z = 0
			break
		}
		if z == 1 {
			break
		}
	}

	return z
}
