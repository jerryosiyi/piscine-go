package piscine

func ToUpper(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] > 96 && r[i] < 123 {
			r[i] = r[i] - 32
		}
	}
	return string(r)
}
