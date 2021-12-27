package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// var v = i.(T)
	// interfaceがTを保持しているか確認するために型アサーションは2つの値を返すことが出来る
	// 型アサーションに失敗した場合、タイプTのゼロ値を返す
}
