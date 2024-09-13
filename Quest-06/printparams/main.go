package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, arr := range os.Args[1:] {
		for _, r := range arr {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
