package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "bond",
		age:   32,
	}

	fmt.Println(p1)
}
