package main

import "fmt"

type OpFn func(int, int) int

func getOperator(op string) OpFn {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	var operator OpFn
	operator = getOperator("*")

	var result = operator(3, 5)
	fmt.Println(result)
}
