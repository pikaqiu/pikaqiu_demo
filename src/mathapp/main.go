package main

import (
	//"bufio"
	//"fmt"
	"mysql"
	//"os"
)

func main() {
	//fmt.Printf("Sqrt(2) = %v\n", mymath.Sqrt(2))

	//当前所选取的机器人logic
	//currentLogic := "pikaqiu"

	//循环读取用户的输入
	/*running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		switch command {
		case "stop":
			running = false
		case "pikaqiu":
			fmt.Println("currentLogic is ", currentLogic)
		}
		fmt.Println(command)
	}*/

	//测试mysql
	mysql.TestMysql()
}
