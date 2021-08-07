package main

import (
	"excelfromdb/excelops"
	"excelfromdb/shell"
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func strtoslice(output string) [][]string {
	a := [][]string{}
	b := strings.Split(strings.Trim(output, "\n"), "\n")
	for _, data := range b {
		a = append(a, []string{data})

	}
	return a

}

func AcessStat() error {
	output, err := shell.AccessStatOutput()
	if err != nil {
		return err
	}
	accesslogslice := strtoslice(output)

	newexcel := "demo3.xlsx"
	insertplace := []struct {
		x string
		y int
	}{{x: "B", y: 5}}

	// 修改excelize源码，添加了读取embed.FS的方法
	f, err := excelize.OpenFileEmbed(publicityTemplate, string(publictiy_template))

	if err != nil {
		return err
	}

	for _, place := range insertplace {
		excelops.GenerateExcelByRowsAnyPlace(f, place.x, acceessstat_sheet, place.y, accesslogslice)

	}

	err = f.SaveAs(newexcel)
	if err != nil {
		fmt.Println("文件保存失败")
		return err
	}
	return nil

}
