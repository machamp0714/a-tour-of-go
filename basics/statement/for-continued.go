package main

import "fmt"

// for文で条件を省略すると無限ループになる

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
