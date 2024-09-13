package piscine

func AlphaCount(s string) int {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	i := 0
	for _, r := range s {
		for _, a := range alpha {
			if r == a {
				i++
				break
			}
		}
	}
	return i
}
