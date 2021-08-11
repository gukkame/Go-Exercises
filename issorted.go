package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	arr := []bool{}
	for i := 0; i < len(a); i++ {
		if f(i, i+1) == 1 {
			arr = append(arr, true)
		}
		if f(i, i+1) == -1 {
			arr = append(arr, false)
		}
	}
	for _, v := range arr {
		if v == false {
			return false
		} else {
			return true
		}
	}
}

func Order(ft, sc int) int {
	for ft := 0; ft <= sc; ft++ {
		return 1
	}
	for ft := 0; ft > sc; ft++ {
		return -1
	}
}
