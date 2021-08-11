package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	arr := []bool{}
	for i := 0; i < len(a); i++ {
		if f(i, i+1) == 1 {
			arr = append(arr, true)
		} else if f(i, i+1) == -1 {
			arr = append(arr, false)
		} else {
			return false
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
	if tr == len(a) {
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
