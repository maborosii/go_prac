package dbconfig

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GenerateExcelByRows(allrows [][]string) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// // Set value of a cell.
	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// // Set active sheet of the workbook.
	// f.SetActiveSheet(index)
	// // Save spreadsheet by the given path.

	err := f.InsertRow("Sheet1", 3)
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		panic(err.Error())
	}
}

func ChangIndexToAxis(intIndexX int, intIndexY int) string {
	var columns_arrs = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	if intIndexX > 26 {
		strIndexX := columns_arrs[intIndexX/26] + columns_arrs[intIndexX%26]
		fmt.Println(strIndexX)
	}

}
