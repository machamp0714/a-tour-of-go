package main

import "fmt"

// 2つ以上の関数パラメータが型を共有するとき、型を省略できる。
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
