package piscine

func FirstRune(s string) rune {
	if len(s) < 1 {
		return '\x00'
	}
	runes := []rune(s)
	return runes[0]
}
