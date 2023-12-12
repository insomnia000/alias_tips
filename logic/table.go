package logic

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
)

// 组装表格
func BuildTable(aliasList []Alias) [][]string {

	//表格过长：
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Num", "Alias", "Tags", "Description", "Command"})

	// 示例数据，包括一个过长的描述
	data := [][]string{}

	for k, alias := range aliasList {
		data = append(data, []string{
			strconv.Itoa(k + 1),
			alias.Name,
			strings.Join(alias.Tags, ","),
			alias.Desc,
			alias.Cmd,
		})
	}

	// 设置表格样式
	//table.SetBorder(false)
	//table.SetHeaderLine(false)
	table.SetAutoWrapText(false) // 启用自动换行
	//table.SetColMinWidth(3, 80)

	// 遍历数据并在需要时手动进行换行
	for _, row := range data {
		table.Append([]string{row[0], row[1], row[2], row[3], row[4]})

		//alias := row[0]
		//tags := row[1]
		//command := row[2]
		//description := row[3]

		// 指定换行的最大字符数!
		//maxLineLength := 180
		//
		//// 将描述文本拆分成多行
		//descriptionLines := splitDescription(description, maxLineLength)
		//fmt.Println(descriptionLines)
		//
		//// 添加行到表格
		//for i, line := range descriptionLines {
		//	if i == 0 {
		//		table.Append([]string{alias, tags, command, line})
		//	} else {
		//		table.Append([]string{"", "", "", line}) // 空字符串用于保持表格对齐
		//	}
		//}

	}

	table.Render()
	return data
}

// 根据最大字符数拆分描述文本
func splitDescription(description string, maxLineLength int) []string {
	var lines []string
	words := strings.Fields(description)
	currentLine := ""

	for _, word := range words {
		if len(currentLine)+len(word)+1 <= maxLineLength {
			// 当前行可以容纳该词
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		} else {
			// 当前行无法容纳，将当前行添加到行列表中
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	// 添加最后一行
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}
