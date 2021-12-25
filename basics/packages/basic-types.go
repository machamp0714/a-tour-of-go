package main

import (
	"fmt"
	"math/cmplx"
)

// 明示的な初期値なしで宣言された変数にはゼロ値が与えられる。

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 2
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v\n", z, z)
}
