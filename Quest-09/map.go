package piscine

func Map(f func(int) bool, a []int) []bool {
	lenA := len(a)
	res := make([]bool, lenA)
	for i, v := range a {
		res[i] = f(v)
	}
	return res
}
