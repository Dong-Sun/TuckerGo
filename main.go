package main

import "fmt"

const (
	C1 = iota
	C2
	C3
)

const (
	B1 = iota + 1
	B2
	B3
)

func main() {
	fmt.Println(C3, B3)
}
