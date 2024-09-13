package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	res := 1
	for i := nb; i > 0; i-- {
		// if res > 9223372036854775807 {
		// 	return 0
		// }
		res *= i
	}
	return res
}
