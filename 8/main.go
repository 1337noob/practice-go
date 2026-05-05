package main

import "fmt"

func Inc(i *int) {
	*i += 1
}

func Dec(i *int) {
	*i -= 1
}
func main() {
	num := 1
	fmt.Println(num)

	Inc(&num)
	fmt.Println(num)

	Dec(&num)
	fmt.Println(num)
}
