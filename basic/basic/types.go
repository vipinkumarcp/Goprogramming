package main

import "fmt"

var a int

//declaring our own type
// declaring type of hotdog as int
type hotdog int

var b hotdog

func main() {

	a = 42
	b = 45
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//assin a=b not possible so converting b to a type
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
