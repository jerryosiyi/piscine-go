package piscine

func ToLower(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] > 64 && r[i] < 91 {
			r[i] = r[i] + 32
		}
	}
	return string(r)
}
