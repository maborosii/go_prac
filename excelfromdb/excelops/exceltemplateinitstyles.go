package excelops

import (
	. "excelfromdb/locallog"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GetStyle(f *excelize.File, styleconfig string) int {
	style, err := f.NewStyle(styleconfig)
	if err != nil {
		Log.Fatal(err)
	}
	return style
}
