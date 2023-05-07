package main

import "fmt"

/*
Callback: passing a function like argument
*/

func main() {
	ii := []int{1, 2, 3, 4, 5, 6}
	s := sum(ii...)
	fmt.Println("All numbers", s)

	// when pass the callback, i pass the identifier
	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("odd numbers", s3)

}

func sum(xi ...int) int{
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// take like parameter a function f and a variadic parameter vi: take only even number and then return their sum
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v %2 == 0{
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v %2 != 0{
			yi = append(yi, v)
		}
	}
	return f(yi...)
}