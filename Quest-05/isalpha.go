package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return true
	}
	alphanumeric := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	matches := 0
	for _, a := range s {
		for _, b := range alphanumeric {
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
