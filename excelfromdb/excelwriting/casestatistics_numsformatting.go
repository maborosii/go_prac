package project

import (
	"excelfromdb/excelops"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func NumsBuildFormat(f *excelize.File, sheetname string, citieslist []string) *[][][]excelops.Cell {

	style_title1 := excelops.GetStyle(f, "/root/golang/excelops/style_all_case_title1.json")
	style_title2 := excelops.GetStyle(f, "/root/golang/excelops/style_all_case_title2.json")
	style_content1 := excelops.GetStyle(f, "/root/golang/excelops/style_all_case_content1.json")
	style_content2 := excelops.GetStyle(f, "/root/golang/excelops/style_all_case_content2.json")

	// 这里不能直接使用goroutine同时写入excel文件，会造成数据丢失，这里需要采用锁，保证一个时刻只有一个goroutine在写文件
	for _, cell := range *excelops.GetTitle(style_title1, style_title2) {
		excelops.Formatting(&cell, sheetname, f)
		excelops.Writing(&cell, sheetname, f)
	}
	allcitycontent, inputareacell := excelops.GetAllCityContent(style_content1, style_content2, citieslist)

	for _, cell := range *allcitycontent {
		// 需要动态插入值的单元格做空值写入
		excelops.Formatting(&cell, sheetname, f)
		excelops.Writing(&cell, sheetname, f)
	}
	return inputareacell
}
