package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(12, 21))
	fmt.Println(Sum(112, 221))
	fmt.Println(Sum(0x11, 0b11))
}

func Sum(a, b int) int {
	return a + b
}
