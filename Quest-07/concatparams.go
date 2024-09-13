package piscine

func ConcatParams(args []string) string {
	res := ""
	for i, s := range args {
		res += s
		if i != len(args)-1 {
			res += "\n"
		}
	}
	return res
}
