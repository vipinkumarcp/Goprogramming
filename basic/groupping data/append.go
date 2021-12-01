package main

import "fmt"

func main() {

	x := []int{4, 5, 8, 9, 10}
	fmt.Println(x)
	x = append(x, 5, 8, 7, 9)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	//delete the certain element in the pervious slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}
