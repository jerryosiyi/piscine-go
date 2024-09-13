package piscine

func Rot14(s string) string {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rotAlpha := "OPQRSTUVWXYZABCDEFGHIJKLMNopqrstuvwxyzabcdefghijklmn"
	res := []rune(s)
	for i, r := range res {
		for j, rr := range alpha {
			if r == rr {
				res[i] = rune(rotAlpha[j])
				break
			}
		}
	}
	return string(res)
}
