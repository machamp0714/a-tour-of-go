package main

import "fmt"

func main() {
	var i interface{}

	i = 46
	describe(i)

	i = "Hello World"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
