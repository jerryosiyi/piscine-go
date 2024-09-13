package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Power(nb int, ex int) int {
	if ex < 0 {
		return 0
	}
	if ex == 0 {
		return 1
	}
	if ex > 0 {
		return nb * Power(nb, ex-1)
	}
	return 0
}

func main() {
	Args := os.Args[1:]
	if len(Args) != 1 {
		return
	}
	nb, _ := strconv.Atoi(Args[0])
	v := 0
	ex := 0
	for v <= nb {
		v = Power(2, ex)
		if v == nb {
			printStr("true")
			return
		}
		ex++
	}
	printStr("false")
}
