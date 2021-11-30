package main

import "fmt"

func main() {

	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do no":
		fmt.Println("miss money or bond or no")
	case "M":
		fmt.Println("bond james")
	default:
		fmt.Println("this is default")
	}
}
