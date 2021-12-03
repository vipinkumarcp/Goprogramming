package main

import "fmt"

func main() {

	f := func() {

		fmt.Println("hiiii")
	}
	f()

	g := func(x int) {

		fmt.Println("hiiiiii", x)
	}

	g(180)

}
