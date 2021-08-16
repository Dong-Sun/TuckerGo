package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	// 초기 잔액
	Balance = 1000
	// 맞췄을 때 버는 양
	EarnPoint = 500
	// 틀렸을 때 잃는 양
	LosePoint = 100
	// 게임 승리 포인트
	VictoryPoint = 5000
	// 게임 오버 포인트
	GameoverPoint = 0
)

var stdin = bufio.NewReader(os.Stdin)

func InputNumber() (int, error) {
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil {
		stdin.ReadString('\n')
	}
	return number, err
}

func main() {
	money := Balance
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Money: %d \n", money)
	for {
		fmt.Print("Input 1 ~ 5:")
		r := rand.Intn(5) + 1
		n, err := InputNumber()
		if err != nil {
			fmt.Println("Please Input Type Int")
		} else if n < 1 || n > 5 {
			fmt.Println("Please Input Value 1 ~ 5")
		} else {
			if n == r {
				fmt.Printf("Matching Number! +%d\n", EarnPoint)
				money += EarnPoint
			} else {
				fmt.Printf("Miss Number...   -%d\n", LosePoint)
				money -= LosePoint
			}
		}
		fmt.Printf("Money: %d \n", money)
		if money >= VictoryPoint || money <= GameoverPoint {
			break
		}
	}
	fmt.Println("Game End")
}
