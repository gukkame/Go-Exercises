package piscine

func MakeRange(min, max int) []int {
	if min < max {
		s := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			s[i] = i + min
		}
		return s
	} else {
		return nil
	}
}
