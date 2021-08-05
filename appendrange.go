package piscine

func AppendRange(min, max int) []int {
	var s []int
	if min < max {
		for i := min - 1; i < max-1; i++ {
			s = append(s, i+1)
		}
		return s
	} else {
		return nil
	}
}
