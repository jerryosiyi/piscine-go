package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n < 10 {
		z01.PrintRune(rune(n + 48))
		return
	}
	s := ToString(n)
	sr := []rune(s)
	for i := 0; i < len(sr); i++ {
		for j := 0; j < len(sr)-1; j++ {
			if sr[j] > sr[j+1] {
				sr[j], sr[j+1] = sr[j+1], sr[j]
			}
		}
	}
	for _, r := range sr {
		z01.PrintRune(r)
	}
}

func ToString(nb int) string {
	if nb == 0 {
		return "0"
	}
	result := ""
	for nb > 0 {
		result = string(rune('0'+nb%10)) + result
		nb /= 10
	}
	return result
}
