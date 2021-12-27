package main

import (
	"fmt"
	"math"
)

// Goにクラスという概念はないがtype上にメソッドを定義出来る

type Vertex struct {
	X, Y float64
}

// メソッドは特別なレシーバー引数を持つ関数
// この例では(v Vertex)がレシーバーとなる
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
