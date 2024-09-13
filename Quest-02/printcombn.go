package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if n < 0 || n > 9 {
		return
	}
	digits := make([]int, n)
	GenerateComb(digits, n, 0, 0)
}

func GenerateComb(digits []int, n int, start int, index int) {
	if index == len(digits) {
		lastComb := 0
		for i, r := range digits {
			z01.PrintRune(rune('0' + r%10))
			if r == 10-(n-i) {
				lastComb++
			}
		}
		if lastComb != n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
			return
		} else {
			z01.PrintRune('\n')
			return
		}
	}
	for i := start; i < 10; i++ {
		digits[index] = i
		GenerateComb(digits, n, i+1, index+1)
	}
}
