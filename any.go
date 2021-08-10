package piscine

func Any(f func(string) bool, a []string) bool {
	ct := 0

	for _, s := range a {
		if f(s) {
			ct++
		}
	}
	if ct > 0 {
		return true
	} else {
		return false
	}
}
