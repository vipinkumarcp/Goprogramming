package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {

		for j := 0; j <= 5; j++ {

			fmt.Printf("%d *  %d is %d\n", i, j, i*j)
		}
	}
}
