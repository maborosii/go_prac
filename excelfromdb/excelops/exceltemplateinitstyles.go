package excelops

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GetStyle(f *excelize.File, styleconfig string) int {
	style, err := f.NewStyle(styleconfig)
	if err != nil {
		panic(err)
	}
	return style
}
