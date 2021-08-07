package main

import (
	"fmt"
	"unsafe"
)

type Padding struct {
	A int8    // 1
	G int8    // 1
	D uint16  // 2
	F float32 // 4
	B int     // 8
	C float64 // 8
	E int     // 8
}

func main() {
	padding := Padding{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(unsafe.Sizeof(padding))
}
