package excelops

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GenerateExcelByRows(f *excelize.File, sheetname string, allrows [][]string) {

	for i, row := range allrows {
		// for j, filed := range row {
		// 	f.SetCellValue("Sheet1", ChangIndexToAxis(j, i), filed)
		// }
		f.SetSheetRow(sheetname, "A"+strconv.Itoa(i+1), &row)
	}

}

// func ChangIndexToAxis(intIndexX int, intIndexY int) string {
// 	/* 根据二维数组的索引值构建excelize的索引,形如"a1","a2","ab1"等
// 	 */
// 	var columns_arrs = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
// 	var strIndexX, strIndexY string
// 	if intIndexX > 25 {
// 		if (intIndexX+1)%26 != 0 {
// 			strIndexX = columns_arrs[(intIndexX+1)/26-1] + columns_arrs[(intIndexX+1)%26-1]
// 		} else {
// 			strIndexX = columns_arrs[(intIndexX+1)/26-1] + columns_arrs[len(columns_arrs)-1]
// 		}

// 	} else {
// 		strIndexX = columns_arrs[intIndexX]
// 	}
// 	strIndexY = strconv.Itoa((intIndexY + 1))
// 	return strIndexX + strIndexY

// }
