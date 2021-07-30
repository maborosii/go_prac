package excelops

import (
	"io/ioutil"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Title *[6]Cell
type CityContent *[29]Cell
type AllContent []CityContent

func GetStyle(jsonpath string) *excelize.Style {
	styleconfigbytes, err := ioutil.ReadFile(jsonpath)

	if err != nil {
		panic(err)
	}
	styleconfig := string(styleconfigbytes)
}
