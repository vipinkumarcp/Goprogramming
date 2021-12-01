package main

import "fmt"

func main() {

	//composite literal

	m := map[string]int{

		"james":       32,
		"money penny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["james"])

	//checking the statement
	if v, ok := m["james"]; ok {
		fmt.Println("This is the test print", v)
	}

}
