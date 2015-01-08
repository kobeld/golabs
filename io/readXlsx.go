package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "/Users/kobeld/Desktop/222.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {
			if index == 0 {
				continue
			}

			println(index)
			if len(row.Cells) > 0 {
				fmt.Printf("%+v*", row.Cells[0])
			}
		}
	}
}
