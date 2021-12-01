package main

import "fmt"

func main() {

	m := map[string]int{

		"james": 32,
		"bond":  45,
	}

	println(m["james"])

	m["todd"] = 55
	fmt.Println(m)

	for i, v := range m {

		fmt.Println(i, v)
	}

}
