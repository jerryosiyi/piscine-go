package piscine

func BasicAtoi(s string) int {
	arr := []rune(s)
	res := 0
	for i := 0; i < len(arr); i++ {
		value := int(arr[i]) - 48
		res += value * Power(10, len(arr)-1-i)
	}
	return res
}

func Power(b int, e int) int {
	res := 1
	for e > 0 {
		res *= b
		e--
	}
	return res
}
