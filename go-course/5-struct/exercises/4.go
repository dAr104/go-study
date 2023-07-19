package main	

import "fmt"

/*
Create and use an anonymous struct.
*/

func main() {

	anonymous := struct {
			doors int
			color string
		}{
			doors: 4,
			color: "white",
		}

	fmt.Println(anonymous)

	fmt.Println(anonymous.doors, anonymous.color)
}