package piscine

func Atoi(s string) int {
	slice := []rune(s)
	negative := false
	positive := false
	var res int
	if len(slice) == 0 {
		return 0
	}
	if slice[0] == '-' {
		negative = true
	} else if slice[0] == '+' {
		positive = true
	}
	if negative || positive {
		res = BasicAtoi2(string(slice[1:]))
	} else {
		res = BasicAtoi2(s)
	}
	if positive {
		return res
	} else if negative {
		return -res
	} else {
		return res
	}
}
