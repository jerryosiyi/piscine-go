package main

import "github.com/01-edu/z01"

func main() {
	arr := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for i := 0; i < len(arr); i++ {
		z01.PrintRune(arr[i])
	}
	z01.PrintRune('\n')
}
