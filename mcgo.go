package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1 // 生成1到100之间的随机数
	var guess int
	attempts := 0

	fmt.Println("欢迎来到猜数字游戏！")
	fmt.Println("我已经想好了一个1到100之间的数字，请猜猜看是多少。")

	for {
		fmt.Print("请输入你的猜测: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("输入无效，请输入一个整数。")
			continue
		}
		attempts++

		if guess < target {
			fmt.Println("太小了，再试一次！")
		} else if guess > target {
			fmt.Println("太大了，再试一次！")
		} else {
			fmt.Printf("恭喜你，猜对了！你一共猜了 %d 次。\n", attempts)
			break
		}
	}
}
