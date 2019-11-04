package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

// RowRule
type RowRule struct {
	RuleName string   // 名称,问题
	Word     []string // 关键词
	Content  []string // 回复语
}

func main() {
	var (
		data  *RowRule
		datas []*RowRule
	)
	file, err := xlsx.OpenFile("/data/gocode/src/github.com/srlemon/note/other/关键词回复demo.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 解析数据
	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows[1:] {
			// TODO 总共有三个列,目前先在这里写死
			if len(row.Cells) >= 3 {
				data = new(RowRule)
				data.RuleName = row.Cells[0].String()
				// 通过#号来分割关键词和content
				data.Word = strings.Split(row.Cells[1].String(), "#")
				data.Content = strings.Split(row.Cells[2].String(), "#")
				datas = append(datas, data)
			}
		}
	}

	fmt.Println(datas)

}
