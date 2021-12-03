package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretagent struct {
	person
	ltk bool
}

// func (r reciver) identifier(parameters(s)) (return(s)) {code}

func (s secretagent) speak() {

	fmt.Println("Iam ", s.first, s.last)
}

func main() {

	sa1 := secretagent{

		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	sa2 := secretagent{

		person: person{
			first: "miss",
			last:  "moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sa1, sa2)
	sa1.speak()
	sa2.speak()

}
