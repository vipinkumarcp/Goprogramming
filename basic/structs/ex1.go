/* Hands-on exercise #1
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.*/

package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {

	p1 := person{
		first:    "ram",
		last:     "krishna",
		icecream: []string{"vanilla", "strawberry", "badam"},
	}

	p2 := person{
		first:    "ravi",
		last:     "dim",
		icecream: []string{"pista", "strawberry", "badam"},
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.icecream {
		fmt.Println(i, v)
	}
	fmt.Println(p1, p2)
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.icecream {
		fmt.Println(i, v)
	}

}
