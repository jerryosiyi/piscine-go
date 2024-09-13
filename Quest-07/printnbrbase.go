package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	counter := 0
	fail := false
	minus := false
	for i, v := range base {
		if v == '-' || v == '+' {
			fail = true
		}
		counter++
		for j, v2 := range base {
			if v == v2 && i != j {
				fail = true
			}
		}
	}
	if counter < 2 {
		fail = true
	}
	if fail {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nbr < 0 {
			minus = true
			nbr = nbr * (-1)
		}
		res := ""
		bRune := []rune(base)
		for nbr != 0 {
			x := nbr % counter
			if x < 0 {
				x = x * (-1)
			}
			res = string(bRune[x]) + res
			nbr = nbr / counter
		}
		if minus {
			z01.PrintRune('-')
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
	}
}
