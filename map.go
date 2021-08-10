package piscine

func Map(f func(int) bool, a []int) []bool {
	for i := range a {
		f(a[i])
	}
}

func IsPrime(value int) bool {
	if value <= 1 {
		return false
	}
	for i := 2; i < value; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
