/* Take the code from the previous exercise, then store the values of type person in a map with the
key of last name. Access each value in the map. Print out the values, ranging over the slice. */

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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m) // output will be map[dim:{ravi dim [pista strawberry badam]} krishna:{ram krishna [vanilla strawberry badam]}]

	for _, v := range m {

		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.icecream {
			fmt.Println(i, val)
		}
	}

}
