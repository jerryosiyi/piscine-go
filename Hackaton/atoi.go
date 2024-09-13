package piscine

func PerfAtoi(s string) int {
	res := 0
	for _, r := range s {
		if !(r >= '0' && r <= '9') {
			return 0
		}
		res = res*10 + int(r-'0')
	}
	return res
}

func Atoi(s string) int {
	if len(s) == 0 || s == "-" || s == "+" {
		return 0
	}
	res := 0
	switch s[0] {
	case '-':
		res = PerfAtoi(s[1:])
		return -res
	case '+':
		res = PerfAtoi(s[1:])
		return res
	default:
		res = PerfAtoi(s)
		return res
	}
}
