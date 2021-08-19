package excelwriting

import (
	"excelfromdb/excelops"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func InputAllNums(citycells []excelops.Cell, sheetname string, f *excelize.File, datas [][]string) {
	for i, data := range datas[1] {
		f.SetCellValue(sheetname, (citycells[i].Xzone[0] + strconv.Itoa(citycells[i].Yzone[0])), data)
	}

}

func InputDeltaNums(citycells []excelops.Cell, sheetname string, f *excelize.File, datas [][]string) {

	for i, data := range datas[1:] {
		f.SetCellValue(sheetname, (citycells[i].Xzone[0] + strconv.Itoa(citycells[i].Yzone[0])), data[0])
	}

}
