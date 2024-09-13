package main

import (
	"os"
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

func ToString(nb int) string {
	if nb == 0 {
		return "0"
	}
	negative := false
	if nb < 0 {
		negative = true
		nb = -nb
	}
	result := ""
	for nb > 0 {
		result = string(rune('0'+nb%10)) + result
		nb /= 10
	}
	if negative {
		result = "-" + result
	}
	return result
}

func main() {
	Args := os.Args[1:]
	if len(Args) < 3 {
		return
	}
	opperands := "+-*/%"
	valid := false
	for _, r := range opperands {
		if Args[1] == string(r) {
			valid = true
			break
		}
	}
	if !valid {
		return
	}
	for i, r := range Args[0] {
		if r == '-' && i == 0 {
			continue
		}
		if r < 48 || r > 57 {
			return
		}
	}
	for i, r := range Args[2] {
		if r == '-' && i == 0 {
			continue
		}
		if r < 48 || r > 57 {
			return
		}
	}
	num1 := TrimAtoi(Args[0])
	num2 := TrimAtoi(Args[2])
	if num1 >= 9223372036854775807 || num2 >= 9223372036854775807 {
		return
	}
	if num1 <= -9223372036854775800 || num2 <= -9223372036854775800 {
		return
	}
	res := 0
	switch Args[1] {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "%":
		if Args[2] == "0" {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		res = num1 % num2
	case "/":
		if Args[2] == "0" {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		res = num1 / num2
	}
	os.Stdout.Write([]byte(ToString(res) + "\n"))
}
