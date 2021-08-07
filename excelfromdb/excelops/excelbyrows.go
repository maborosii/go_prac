package excelops

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GenerateExcelByRows(f *excelize.File, sheetname string, allrows [][]string) {

	for i, row := range allrows {
		f.SetSheetRow(sheetname, "A"+strconv.Itoa(i+1), &row)
	}

}

func GenerateExcelByRowsAnyPlace(f *excelize.File, x, sheetname string, y int, allrows [][]string) {

	for i, row := range allrows {
		f.SetSheetRow(sheetname, x+strconv.Itoa(y+i), &row)
	}

}
