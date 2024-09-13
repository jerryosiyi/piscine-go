package piscine

func TrimAtoi(s string) int {
	arr := []rune{}
	digit := false
	negative := false
	for _, r := range s {
		if r >= 48 && r <= 57 {
			arr = append(arr, r)
			digit = true
		} else if r == 45 && !digit {
			negative = true
		}
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		value := int(arr[i]) - 48
		res += value * Power(10, len(arr)-1-i)
	}
	if negative {
		res = -res
	}
	return res
}

func Power(nb int, power int) int {
	if power < 0 {
		return 0
	}
	res := 1
	for i := power; i > 0; i-- {
		res *= nb
	}
	return res
}
