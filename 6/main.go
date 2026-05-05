package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	fmt.Println(sum)
}
