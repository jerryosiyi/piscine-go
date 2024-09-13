package piscine

func IsUpper(s string) bool {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	matches := 0
	for _, a := range s {
		for _, b := range upper {
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
