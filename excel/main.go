package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"net/http"
)

var sheet2 = "sheet2"
var sheet1 = "sheet1"

func main() {
	setExce()
	openExce()
	addChart()
	h, _ := http.Get("https://five2.shanchayou028.com/demo/data433.xlsx")
	f, _ := excelize.OpenReader(h.Body)
	a := f.GetCellValue("sheet1", "A2")
	fmt.Println(a)
}

func setExce() {
	file := excelize.NewFile()
	// create new sheet
	index := file.NewSheet(sheet2)

	// set value of a cell
	file.SetCellValue(sheet2, "A2", "hello world 哈哈哈")
	file.SetCellValue(sheet1, "B2", 1656)

	// set active sheet of the workbook
	file.SetActiveSheet(index)
	err := file.SaveAs("test.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func openExce() {
	f, err := excelize.OpenFile("./test.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	cell := f.GetCellValue(sheet2, "A2")
	fmt.Println(cell, "aaa")

	rows := f.GetRows(sheet2)
	for _, v := range rows {
		for _, v1 := range v {
			fmt.Println(v1, "\t")
		}
		fmt.Println()
	}
}

type serie struct {
	Name       string `json:"name"`
	Categories string `json:"categories"`
	Values     string `json:"values"`
}
type title struct {
	Name string `json:"name"`
}
type format struct {
	Type   string  `json:"type"`
	Series []serie `json:"series"`
	Title  title   `json:"title"`
}

func addChart() {
	categories := map[string]string{"A2": "small", "A3": "Normal", "A4": "Large", "A5": "TODO",
		"B1": "Apple", "C1": "Orange", "D1": "Pear",
	}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3,
		"B3": 5, "C3": 2, "D3": 4,
		"B4": 6, "C4": 7, "D4": 8,
		"B5": 10, "C5": 20, "D5": 16,
	}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue(sheet1, k, v)
	}
	for k, v := range values {
		f.SetCellValue(sheet1, k, v)
	}
	var fo *format
	fo = new(format)
	fo.Type = "col3DClustered"
	fo.Series = []serie{{Name: "A2", Categories: "B1:D1", Values: "B2:D2"},
		{Name: "A3", Categories: "B1:D1", Values: "B3:D3"},
		{Name: "A4", Categories: "B1:D1", Values: "B4:D4"},
		{Name: "A5", Categories: "B1:D1", Values: "B5:D5"},
	}
	fo.Title = title{
		Name: "哈哈哈哈哈",
	}
	s, _e := json.Marshal(fo)
	fmt.Println(_e)
	err := f.AddChart(sheet1, "E1", string(s))
	if err != nil {
		fmt.Println(err)
	}
	err = f.SaveAs("./test1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
