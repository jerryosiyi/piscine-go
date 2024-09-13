package main

import (
	"fmt"
	"os"
)

func main() {
	Args := os.Args[1:]
	for _, args := range Args {
		if args == "01" || args == "galaxy" || args == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
