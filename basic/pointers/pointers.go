package main

import "fmt"

func main() {

	x := 45
	fmt.Println("x before", x)
	fmt.Println("x adress", &x)
	foo(&x)
	fmt.Println("x adress after", &x)
	fmt.Println("x after", x)

}

func foo(y *int) {

	fmt.Println("recived y", y)
	fmt.Println("recived y", *y)

	*y = 58

}
