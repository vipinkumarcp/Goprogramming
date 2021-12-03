// func (r reciver) identifier(parameters(s)) (return(s)) {code}

package main

import "fmt"

func main() {

	x := num(1, 2, 3, 4, 5, 6)
	fmt.Println("sum", x)

}

func num(x ...int) int {

	sum := 0
	for _, v := range x {

		sum += v

	}

	fmt.Println(sum)
	return sum

}
