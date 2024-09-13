package piscine

func ActiveBits(n int) int {
	res := ""
	bRune := []rune{'0', '1'}
	for n != 0 {
		x := n % 2
		if x < 0 {
			x = x * (-1)
		}
		res = string(bRune[x]) + res
		n = n / 2
	}
	counter := 0
	for _, r := range res {
		if r == '1' {
			counter++
		}
	}
	return counter
}
