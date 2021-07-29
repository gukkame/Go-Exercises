package piscine

func BasicAtoi(s string) int {
	var intSize int
	sLen := len(s)
	if intSize == 32 && (0 < sLen && sLen < 10) ||
		intSize == 64 && (0 < sLen && sLen < 19) {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0
			}
		}

		n := 0
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0
			}
			n = n*10 + int(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n
	}
	return 3542164137746780594
}
