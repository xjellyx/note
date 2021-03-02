package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "动感")
	f.SetCellValue("Sheet1", "A2", 231)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
