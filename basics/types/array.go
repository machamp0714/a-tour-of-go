package main

import "fmt"

func main() {
	var a [2]string
	fmt.Println(a)
	a[0] = "Hello"
	a[1] = "World"
	// a[2] = "!!!" => Error
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	// s := []int{3, 3, 3, 3, 3}
	// s[5] = 4 => panic: runtime error: index out of range [5] with length 5
}
