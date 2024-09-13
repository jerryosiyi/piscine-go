package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i, s := range strs {
		if i != 0 {
			res += sep
		}
		res += s
	}
	return res
}
