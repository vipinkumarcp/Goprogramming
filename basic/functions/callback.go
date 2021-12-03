package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := sum(ii...)

	fmt.Println("sum of all number:", s)

	e := even(sum, ii...)
	fmt.Println("sum of all even number:", e)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, v ...int) int {

	var yi []int
	for _, i := range v {

		if i%2 == 0 {
			yi = append(yi, i)
		}
	}
	return f(yi...)
}
