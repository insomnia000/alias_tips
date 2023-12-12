package logic

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

var retryTimes = 2
var retryCount = 0

// 监听键盘输入
func ListenKeyboard() int {

	// 打开键盘输入
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("请输入想要执行的命令编号，按Enter确定 (按ESC键或q键退出) ...")

	var input string
	var isExit bool = false
	count := 0
	for count < 1000 {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		// 检查按键类型
		if key == keyboard.KeyEsc || string(char) == "q" || string(char) == "Q" {
			isExit = true
			return 0
		}
		// 检查是否是字符输入
		if key == keyboard.KeySpace {
			// 空格键不算字符，跳过
			continue
		}
		fmt.Printf("%s", string(char))

		// 检查回车退出循环
		if key == keyboard.KeyEnter {
			break
		}

		input += string(char)
		count++
	}

	if isExit {
		fmt.Println("Quit...")
		return 0
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		retryCount += 1
		if retryCount > retryTimes {
			return 0
		}
		fmt.Println("输入错误，请重新输入数字编号...")
		return ListenKeyboard()
	}
	fmt.Println(" ")
	return num
}
