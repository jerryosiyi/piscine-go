package piscine

func IsNumeric(s string) bool {
	digits := "0123456789"
	matches := 0
	for _, a := range s {
		for _, b := range digits {
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
