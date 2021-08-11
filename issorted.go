package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	arr := []bool{}
	for i := 1; i < len(a); i++ {
		if Order(a[i-1], a[i]) == 1 {
			arr = append(arr, true)
		}
		if Order(a[i-1], a[i]) == -1 {
			arr = append(arr, false)
		} else {
			arr = append(arr, false)
		}
	}
	fa := 0

	for _, v := range arr {
		if !v {
			fa++
		}
	}
	if fa > 0 {
		return false
	} else {
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
