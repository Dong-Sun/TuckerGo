package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	}
	return n, err
}

func CheckNumber(n int, r int) bool {
	fmt.Println(n)
	if n < r {
		fmt.Println("Smaller")
		return false
	} else if n > r {
		fmt.Println("Bigger")
		return false
	} else {
		fmt.Println("Same")
		return true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	cnt := 0

	for {
		cnt++
		fmt.Print("InputIntValue:")
		n, err := InputIntValue()
		if err != nil {
			continue
		}
		if CheckNumber(n, r) {
			break
		}
	}
}
