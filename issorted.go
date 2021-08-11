package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	arr := []bool{}
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) == 1 {
			arr = append(arr, true)
		}
		if f(a[i-1], a[i]) == -1 {
			arr = append(arr, false)
		} else {
			arr = append(arr, false)
		}
	}
	fa := 0
	tr := 0
	for _, v := range arr {
		if !v {
			fa++
		} else {
			tr++
		}
	}
	if tr == len(arr) {
		return true
	} else {
		return false
	}
}

func f(ft, sc int) int {
	if ft <= sc {
		return 1
	}
	if ft > sc {
		return -1
	}
	return 0
}
