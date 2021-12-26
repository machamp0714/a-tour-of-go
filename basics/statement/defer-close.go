package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close() // Createに失敗した時、Fileがcloseされない問題に対応するため

	return io.Copy(dst, src)
}

func main() {
	written, err := CopyFile("copy.go", "./basics/statement/defer.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(written)
}
