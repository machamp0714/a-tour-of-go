package main

import "fmt"

func Test(a *int) { // aはint型のポインタ変数になる
	*a += 1 // aポインタが指している値を+1する
	fmt.Println("aの値(関数内):", *a)
}

func main() {
	a := 10

	fmt.Println("aの値:", a)
	Test(&a)
	fmt.Println("aの値:", a)
}
