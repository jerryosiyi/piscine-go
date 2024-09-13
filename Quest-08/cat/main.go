package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	Args := os.Args[1:]
	if len(Args) == 0 {
		file, err := io.ReadAll(os.Stdin)
		if err != nil {
			str := "ERROR: " + err.Error()
			printStr(str)
			os.Exit(1)
		}
		printStr(string(file))
	} else {
		for _, arg := range Args {
			file, err := os.Open(arg)
			if err != nil {
				str := "ERROR: " + err.Error()
				printStr(str)
				os.Exit(1)
			}
			defer file.Close()
			f, err := io.ReadAll(file)
			if err != nil {
				str := "ERROR: " + err.Error()
				printStr(str)
				os.Exit(1)
			}
			printStr(string(f))
		}
	}
}
