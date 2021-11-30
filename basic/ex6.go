package main

import "fmt"

var x int

func main() {

	x = 45
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)

}
