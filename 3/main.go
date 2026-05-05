package main

import "fmt"

func main() {
	fmt.Println(0, Parity(0))
	fmt.Println(1, Parity(1))
	fmt.Println(2, Parity(2))
	fmt.Println(-10, Parity(-10))
}

func Parity(n int) string {
	if n%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
