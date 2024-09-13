package piscine

func LastRune(s string) rune {
	if len(s) < 1 {
		return '\x00'
	}
	runes := []rune(s)
	length := len(runes)
	return runes[length-1]
}
