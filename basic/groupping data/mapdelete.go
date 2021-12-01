package main

import "fmt"

func main() {

	m := map[string]int{

		"bond": 53,
		"kim":  56,
		"dim":  45,
	}
	fmt.Println(m)
	delete(m, "dim")
	fmt.Println(m)

	if i, v := m["bond"]; v {

		fmt.Println("value:", i)
		delete(m, "bond")
	}
	fmt.Println(m)
}
