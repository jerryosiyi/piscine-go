package main

import (
	"os"

	"github.com/01-edu/z01"
)

func TrimAtoi(s string) int {
	arr := []rune{}
	digit := false
	negative := false
	for _, r := range s {
		if r >= 48 && r <= 57 {
			arr = append(arr, r)
			digit = true
		} else if r == 45 && !digit {
			negative = true
		}
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		value := int(arr[i]) - 48
		res += value * Power(10, len(arr)-1-i)
	}
	if negative {
		res = -res
	}
	return res
}

func Power(nb int, power int) int {
	if power < 0 {
		return 0
	}
	res := 1
	for i := power; i > 0; i-- {
		res *= nb
	}
	return res
}

func main() {
	arr := os.Args[1:]
	if len(arr) == 0 {
		return
	}
	upper := false
	if os.Args[1] == "--upper" {
		upper = true
		arr = os.Args[2:]
	}
	for _, s := range arr {
		i := TrimAtoi(s)
		if i < 1 || i > 26 {
			z01.PrintRune(' ')
			continue
		}
		if upper {
			z01.PrintRune(rune(i + 64))
		} else {
			z01.PrintRune(rune(i + 96))
		}
	}
	z01.PrintRune('\n')
}
