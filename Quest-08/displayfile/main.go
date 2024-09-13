package main

import (
	"fmt"
	"os"
)

func main() {
	Args := os.Args[1:]
	if len(Args) == 0 {
		fmt.Println("File name missing")
		return
	} else if len(Args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := os.Open(Args[0])
	if err != nil {
		fmt.Printf("the mistake is: %v\n", err.Error())
		return
	}
	arr := make([]byte, 14)
	file.Read(arr)
	fmt.Println(string(arr))
}
