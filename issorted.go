package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	arr := []bool{}
	for i := 1; i < len(a); i++ {
		if Order(a[i-1], a[i]) == 1 {
			arr = append(arr, true)
		}
		if Order(a[i-1], a[i]) == -1 {
			arr = append(arr, false)
		}
	}

	for _, v := range arr {
		if !v {
			return false
		}
		return true
	}
}

func Order(a, b int) int {
	if a <= b {
		return 1
	}
	if a > b {
		return -1
	}
	return 0
}
