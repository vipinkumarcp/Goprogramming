package main

import "fmt"

func main() {

	x := []int{4, 8, 9, 6, 3}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Print(x[3])

	// i and v are range and value

	for i, v := range x {

		fmt.Println(i, v)
	}
}
