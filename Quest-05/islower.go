package piscine

func IsLower(s string) bool {
	lower := "abcdefghijklmnopqrstuvwxyz"
	matches := 0
	for _, a := range s {
		for _, b := range lower {
			if a == b {
				matches++
			}
		}
	}
	if matches == len(s) {
		return true
	} else {
		return false
	}
}
