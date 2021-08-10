package piscine

func CountIf(f func(string) bool, tab []string) int {
	ct := 0

	for _, s := range tab {
		if f(s) {
			ct++
		}
	}
	if ct > 0 {
		return ct
	} else {
		return 0
	}
}
