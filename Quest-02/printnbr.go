package piscine

// Quest 02

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	for _, r := range NbrToString(n) {
		z01.PrintRune(r)
	}
}

func NbrToString(nb int) string {
	if nb == 0 {
		return "0"
	}
	if nb == -9223372036854775808 {
		return "-9223372036854775808"
	}
	result := ""
	negative := false
	if nb < 0 {
		negative = true
		nb *= -1
	}
	for nb > 0 {
		result = string(rune('0'+nb%10)) + result
		nb /= 10
	}
	if negative {
		result = "-" + result
	}
	return result
}
