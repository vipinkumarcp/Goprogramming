package main

import "fmt"

func main() {

	//specify type,length,capacity
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
