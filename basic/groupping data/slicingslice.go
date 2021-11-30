package main

import "fmt"

func main() {

	x := []int{1, 5, 8, 9, 8, 8, 9, 7}
	fmt.Println(x[1])
	fmt.Println(x[:3])
	fmt.Println(x[:])

	for i := 0; i <= 7; i++ {

		fmt.Println(i, x[i])
	}

}
