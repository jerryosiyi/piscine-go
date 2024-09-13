package piscine

func BasicAtoi2(s string) int {
	if IsValidStr(s) {
		return BasicAtoi(s)
	}
	return 0
}

func IsValidStr(s string) bool {
	digits := "0123456789"
	for _, r := range s {
		if IndexRune(r, digits) == -1 {
			return false
		}
	}
	return true
}

func IndexRune(r rune, s string) int {
	for i, c := range s {
		if r == c {
			return i
		}
	}
	return -1
}
