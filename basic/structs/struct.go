package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "james",
		last:  "bond",
		age:   25,
	}

	p2 := person{
		first: "sonny",
		last:  "raj",
		age:   65,
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p2.last, p2.age)
}
