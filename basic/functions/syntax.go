package main

import "fmt"

func main() {

	foo()
	bar("car")
	x, y := barg("honey", "sweet")
	fmt.Println(x, y)
}

func foo() {

	fmt.Println("THis is from foo")
}

func bar(s string) {

	fmt.Println("this is bar function", s)
}

//specify return values as string a bool
func barg(s string, g string) (string, bool) {

	a := fmt.Sprint(s, " ", g, `, says "Hello"`)
	b := true

	return a, b

}
