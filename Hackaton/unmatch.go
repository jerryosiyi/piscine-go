package piscine

func Unmatch(a []int) int {
	for i, x := range a {
		pair := 1
		for j, y := range a {
			if x == y && i != j {
				pair++
			}
		}
		if pair%2 != 0 {
			return x
		}
	}
	return -1
}
