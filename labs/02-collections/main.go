package main

import "fmt"

func main() {
	fmt.Println("Collections")
	// Array
	var x [5]int

	x[0] = 100
	x[1] = 101
	x[3] = 103
	x[4] = 104

	fmt.Println(x)

	// Slices
	var s1 []int = x[1:4]
	var s2 []int = x[:4]
	var s3 []int = x[1:]
	var s4 []int = x[:]
	fmt.Println(s1, s2, s3, s4)
}
