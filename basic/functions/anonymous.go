package main

import "fmt"

func main() {

	func(x int) {
		fmt.Println("hiii", x)
	}(42)
}
