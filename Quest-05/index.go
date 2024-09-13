package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	for pos, r := range s {
		// Check if the first rune of toFind = r (current rune of s)
		if r == FirstRune(toFind) {
			matches := 1
			// Loop through the next runes of toFind
			// Check if the match the next runes of s
			for i := 2; i <= len(toFind); i++ {
				if NRune(toFind, i) == NRune(s, pos+i) {
					matches++
				}
			}
			// If matches == length of toFind then
			// the substring has been found so return pos
			if matches == len(toFind) {
				return pos
			}
		}
	}
	return -1
}
