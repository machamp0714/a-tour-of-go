package main

import "fmt"

func main() {
	// sliceにはlengthとcapの両方がある
	// sliceのlengthはsliceに含まれる要素の数。
	// sliceのcapは基になる配列の要素の数。
	s := []int{2, 3, 5, 7, 9, 11, 13}
	printSlice(s)

	s = s[:0] // Slice to zero length
	printSlice(s)

	s = s[:4] // Extend length
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
