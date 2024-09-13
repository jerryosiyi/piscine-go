package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
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

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	str := "x = " + ToString(points.x) + ", y = " + ToString(points.y)
	printStr(str)
}
