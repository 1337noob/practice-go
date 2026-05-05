package main

import "fmt"

func main() {
	i := 44
	s := "message"
	b := true

	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", b, b)
}
