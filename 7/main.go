package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years) from %v", p.Name, p.Age, p.City)
}

func main() {
	bob := Person{"Bob", 23, "New York"}
	fmt.Println(bob)

	jane := Person{"Jane", 31, "Alabama"}
	fmt.Println(jane)

	rachel := Person{"Rachel", 47, "Detroit"}
	fmt.Println(rachel)
}
