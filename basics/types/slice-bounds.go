package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr[1:4])
	fmt.Println(arr[:3])
	fmt.Println(arr[3:])
	fmt.Println(arr[:])
}
