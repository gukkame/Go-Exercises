package piscine

func ConcatParams(args []string) string {
	str := ""

	for i, res := range args {
		str = str + string(res)
		if i != len(args)-1 {
			str = str + "\n"
		}
	}
	return str
}
