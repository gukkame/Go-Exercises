package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a / *b
	*a = c
	d := *a % *b
	*b = d
}
