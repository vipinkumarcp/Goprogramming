package main

import "fmt"

func main() {

	jb := []string{"james", "bond", "icecream"}
	fmt.Println(jb)
	mp := []string{"miss", "moneypenny", "strawberry"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
