/* Create a for loop using this syntax
● for condition { }
Have it print out the years you have been alive. */

package main

import "fmt"

func main() {

	x := 1994
	for {
		fmt.Println(x)
		if x == 2021 {
			break
		}
		x++
	}
}
