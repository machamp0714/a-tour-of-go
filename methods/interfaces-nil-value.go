package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("nil!")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	i.M()

	i = &T{"Hello World"}
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}