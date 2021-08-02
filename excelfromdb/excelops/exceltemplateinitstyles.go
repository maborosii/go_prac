package excelops

import (
	"io/ioutil"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GetStyle(f *excelize.File, jsonpath string) int {
	styleconfigbytes, err := ioutil.ReadFile(jsonpath)

	if err != nil {
		panic(err)
	}
	styleconfig := string(styleconfigbytes)
	style, err := f.NewStyle(styleconfig)
	if err != nil {
		panic(err)
	}
	return style
}
