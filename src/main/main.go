package main

import (
	"bufio"
	"fmt"
	"mysql"
	"os"
)

func main() {
	//fmt.Printf("Sqrt(2) = %v\n", mymath.Sqrt(2))

	//当前所选取的机器人logic
	currentLogic := "pikaqiu"

	//循环读取用户的输入
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		question := string(data)
		switch question {
		case "stop":
			running = false
		case "logic":
			fmt.Println("currentLogic is", currentLogic)
		default:
			fmt.Println(mysql.GetAns(question))
		}
	}

	//测试mysql
	//mysql.TestMysql()
}
