package excelops

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Cell struct {
	IsMerge bool
	Xzone   []string
	Yzone   []int
	Xwidth  []float64
	Yheight []float64
	Format  *excelize.Style
}

func (cl *Cell) SetWidth(sheetname string, f *excelize.File) {
	for i, strX := range cl.Xzone {
		if err := f.SetColWidth(sheetname, strX, strX, cl.Xwidth[i]); err != nil {
			panic(err)
		}
	}
}

func (cl *Cell) SetHeight(sheetname string, f *excelize.File) {
	for i, intY := range cl.Yzone {
		if err := f.SetRowHeight(sheetname, intY, cl.Yheight[i]); err != nil {
			panic(err)
		}
	}
}

func (cl *Cell) MergeCell(sheetname string, f *excelize.File) {
	if cl.IsMerge {
		if err := f.MergeCell(sheetname, cl.Xzone[0]+strconv.Itoa(cl.Yzone[0]), cl.Xzone[len(cl.Xzone)]+strconv.Itoa(cl.Yzone[len(cl.Yzone)])); err != nil {
			panic(err)
		}
	}
}

func (cl *Cell) SetFormat(sheetname string, f *excelize.File) {
	style, err := f.NewStyle(cl.Format)
	if err != nil {
		panic(err)
	}
	err = f.SetCellStyle(sheetname, cl.Xzone[0]+strconv.Itoa(cl.Yzone[0]), cl.Xzone[len(cl.Xzone)]+strconv.Itoa(cl.Yzone[len(cl.Yzone)]), style)
	if err != nil {
		panic(err)
	}
}
