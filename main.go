package main

import (
	"alias_tips/logic"
	"fmt"
	"os"
)

func main() {

	var aliasList []logic.Alias = make([]logic.Alias, 0)
	//获取配置文件
	config := logic.ReadConfig("")

	logic.CleanRunShell(config.Shell)

	//接受用户输入的参数，1,根据tag 筛选展示的列表
	param := logic.GetFirstParam()
	if param == "" {
		//如果用户没有输入参数，则展示全部的命令
		for _, alias := range config.Alias {
			aliasList = append(aliasList, alias)
		}
	} else if param == "init" {
		//如果用户输入init，则初始配置alias
		logic.AliasInit(config)
		return
	} else {
		//如果用户输入参数，根据tag筛选命令
		for _, alias := range config.Alias {
			if logic.InArray(alias.Tags, param) {
				aliasList = append(aliasList, alias)
			}
		}
	}

	//fmt.Println(aliasList)
	//构建表格：
	searchList := logic.BuildTable(aliasList)
	if len(searchList) == 0 {
		fmt.Println("暂无相关命令...")
		return
	}

	//等待用户键入：Q或者回车表示终止
	num := logic.ListenKeyboard()
	if num == 0 {
		fmt.Println("暂无选中命令...")
		return
	}

	//执行用户所选择的命令
	command := ""
	if num-1 < len(searchList) && len(searchList[num-1]) >= 4 {
		command = searchList[num-1][4]
	} else {
		fmt.Println("暂无选中命令...")
		return
	}

	fmt.Println("即将执行命令：", command)
	fmt.Println("---------------------------------------------------------\n")

	//生成shell脚本
	shellHead := ""
	if config.Shell == "zsh" {
		shellHead = "#!/bin/zsh\n"
	} else {
		shellHead = "#!/bin/bash\n"
	}

	shellContent := shellHead + command
	// 创建或打开文件（如果文件不存在则创建，如果存在则截断）
	file, err := os.Create(logic.GetBasePath() + "/run.sh")
	if err != nil {
		fmt.Println("创建shell文件时出错:", err)
		return
	}
	defer file.Close()

	// 将字符串写入文件
	_, err = file.WriteString(shellContent)
	if err != nil {
		fmt.Println("写入shell命令时出错:", err)
		return
	}

}
