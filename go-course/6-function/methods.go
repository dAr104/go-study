package main

import "fmt"

/*
Methods
*/

type person struct{
	first string
	last string
}

type secretAgent struct{
	person
	ltk bool
}

// func (r receiver) identifier(parametes (return(s)) {code}

// when specify a receiver, you attach the function to any value of that type (it is a method)
func (s secretAgent) speak() {
	fmt.Println("I'm", s.first, s.last)
}

func main(){
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Dario",
			"Mario",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	
}

