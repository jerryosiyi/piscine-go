package main

import (
	"fmt"
	"os"
)

func HasPrefix(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func main() {
	Args := os.Args[1:]
	res := ""
	insert := ""
	order := false
	if len(Args) == 0 || Args[0] == "--help" || Args[0] == "-h" {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("	 This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
		return
	}
	for _, s := range Args {
		// Get insert string
		if HasPrefix(s, "--order") || HasPrefix(s, "-o") {
			order = true
		} else if HasPrefix(s, "--insert=") {
			insert = s[len("--insert="):]
		} else if HasPrefix(s, "-i=") {
			insert = s[len("-i="):]
		} else {
			res = s
		}
	}
	res += insert
	if order {
		resr := []rune(res)
		for i := 0; i < len(resr); i++ {
			for j := 0; j < len(resr)-1; j++ {
				if resr[j] > resr[j+1] {
					resr[j], resr[j+1] = resr[j+1], resr[j]
				}
			}
		}
		res = string(resr)
	}
	fmt.Println(res)
}
