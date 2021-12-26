package main

import "fmt"

var pow = []int{1, 2, 3, 4, 5, 6, 7}

func main() {
	for i, v := range pow { // for文はsliceまたはmap上で繰り返される
		fmt.Printf("2**%d = %d\n", i, v) // 反復ごとにインデックスとインデックスの要素のコピー
	}
}
