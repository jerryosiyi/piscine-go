package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	for _, s := range arr {
		for _, r := range s {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
