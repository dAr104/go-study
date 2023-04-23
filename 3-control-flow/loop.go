package main

import "fmt"

func main() {
	// inner loop
	// loop with for clause
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("The outer loop: %d\t the inner loop: %d\n", i, j)
		}
	}

	// while loop
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")

	// leave with break

	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("Done")

}