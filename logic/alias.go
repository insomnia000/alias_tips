package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Project         string  `json:"project"`
	ShellConfigPath string  `json:"shell_config_path"`
	Shell           string  `json:"shell"`
	Alias           []Alias `json:"alias"`
}

type Alias struct {
	Name string   `json:"name"`
	Cmd  string   `json:"cmd"`
	Desc string   `json:"desc"`
	Tags []string `json:"tags"`
}

func ReadConfig(configPath string) Config {
	if configPath == "" {
		configPath = GetBasePath() + "/config"
	}
	var config Config

	// 遍历目录下的所有配置文件
	var confList []string = make([]string, 0)
	filepath.Walk(configPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问路径时发生错误: %v\n", err)
			return err
		}

		if !info.IsDir() {
			// 如果不是目录，可以在这里执行相应的操作
			confList = append(confList, path)
			itemConfig := getConfig(path)
			if itemConfig.Project != "" {
				config.Project = itemConfig.Project
			}
			if itemConfig.ShellConfigPath != "" {
				config.ShellConfigPath = itemConfig.ShellConfigPath
			}
			if itemConfig.Shell != "" {
				config.Shell = itemConfig.Shell
			}
			if len(itemConfig.Alias) > 0 {
				config.Alias = append(config.Alias, itemConfig.Alias...)
			}
		}

		return nil
	})

	return config
}

// 从单个配置文件读取配置：
func getConfig(filePath string) Config {
	var config Config
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件时出错:", err)
		return config
	}

	// 将JSON内容解析到Person结构体中
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		fmt.Println("解析JSON时出错:", err, "文件路径:", filePath)
		return config
	}
	return config
}

// 获取执行文件的路径
func GetBasePath() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(ex)
}

// 初始化别名列表
func AliasInit(config Config) {

	fmt.Println("初始化alias中...")
	outputText := ""
	//输出alias列表
	for _, alias := range config.Alias {
		outputText += fmt.Sprintf("# %s \n", alias.Desc)
		outputText += fmt.Sprintf("alias %s='%s'\n\n", alias.Name, alias.Cmd)
	}
	outputFile := GetBasePath() + "/output_alias.sh"
	err := ioutil.WriteFile(outputFile, []byte(outputText), 0644)
	if err != nil {
		fmt.Println("写入文件时发生错误:", err)
		return
	}
	fmt.Println("alias文件已生成，文件位置：", outputFile)
	fmt.Println("初始化完成...")
}
