package logic

import (
	"fmt"
	"os"
)

// 获取第一个参数
func GetFirstParam() string {

	// 获取命令行参数
	args := os.Args

	// 其他参数位于之后的参数切片
	arguments := args[1:]

	if len(arguments) == 0 {
		return ""
	} else {
		return arguments[0]
	}
}

// 清空待执行命令
func CleanRunShell(shell string) {

	//生成shell脚本
	shellHead := ""
	if shell == "zsh" {
		shellHead = "#!/bin/zsh\n"
	} else {
		shellHead = "#!/bin/bash\n"
	}

	// 创建或打开文件（如果文件不存在则创建，如果存在则截断）
	file, err := os.Create(GetBasePath() + "/run.sh")
	//fmt.Println(logic.GetBasePath() + "/run.sh")
	if err != nil {
		fmt.Println("创建shell文件时出错:", err)
		return
	}
	defer file.Close()
	// 将字符串写入文件
	_, err = file.WriteString(shellHead)
	if err != nil {
		fmt.Println("写入shell命令时出错:", err)
		return
	}

}
