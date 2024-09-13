package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			PrintRune(Format2D(i))
			PrintRune(" ")
			PrintRune(Format2D(j))
			if i == 98 && j == 99 {
				z01.PrintRune('\n')
				return
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}

func Format2D(nb int) string {
	if nb < 10 {
		return "0" + ToString(nb)
	} else {
		return ToString(nb)
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

func PrintRune(str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
}
