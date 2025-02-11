package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 定义谜语和答案的映射
var riddles = map[string]string{
	"什么东西越洗越脏？":             "水",
	"什么东西不能吃？":             "亏",
	"什么东西越用越有钱？":           "存钱罐",
	"什么东西看不见摸不着但能感觉到？": "风",
	"什么东西越分越多？":            "知识",
}

func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 将谜语的键（问题）存入切片
	var riddleQuestions []string
	for question := range riddles {
		riddleQuestions = append(riddleQuestions, question)
	}

	// 随机选择一个谜语
	selectedRiddle := riddleQuestions[rand.Intn(len(riddleQuestions))]
	correctAnswer := riddles[selectedRiddle]

	// 创建读取用户输入的Scanner
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("欢迎来到猜谜语游戏！")
	fmt.Println("谜语是:", selectedRiddle)
	fmt.Print("请输入你的答案: ")

	// 读取用户输入
	scanner.Scan()
	userAnswer := scanner.Text()

	// 判断用户答案是否正确
	if userAnswer == correctAnswer {
		fmt.Println("恭喜你，答对了！")
	} else {
		fmt.Printf("很遗憾，答错了。正确答案是: %s\n", correctAnswer)
	}
}
