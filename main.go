package main

import "fmt"

func F(n int) int {
	// 여기에 탈출 조건을 채우세요.
	if n < 2 {
		return n
	}

	return F(n-2) + F(n-1)
}

func main() {
	// 피보나치 수열 9번째 값을 출력합니다.
	fmt.Println(F(9))
}
