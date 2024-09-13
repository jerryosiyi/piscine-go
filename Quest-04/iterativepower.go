package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	res := 1
	for i := power; i > 0; i-- {
		res *= nb
	}
	return res
}
