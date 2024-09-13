package piscine

func CountIf(f func(string) bool, tab []string) int {
	res := 0
	for _, v := range tab {
		if f(v) {
			res += 1
		}
	}
	return res
}
