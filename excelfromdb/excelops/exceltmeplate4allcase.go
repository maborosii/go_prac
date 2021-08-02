/*
	size： excelize/excel

	列 		 A		 	 B		 	 C	     	  D	 	 	 E
	width 	8.38/10	    22.88/27    42.25/48	 8.63/9.37	14/14.7
	...

	行	height
	1	31
	2	15
	3   16.5
	.	.
	.	.
	.	.
	37	16.5
*/

package excelops

import "strconv"

func GetTitle(style1, style2 int) *[]Cell {
	title := []Cell{
		{
			IsMerge: true,
			Xzone:   []string{"a"},
			Yzone:   []int{1, 2},
			Xwidth:  []float64{10},
			Yheight: []float64{31, 15},
			Format:  style1,
			Content: "序号"},
		{
			IsMerge: true,
			Xzone:   []string{"b"},
			Yzone:   []int{1, 2},
			Xwidth:  []float64{27},
			Yheight: []float64{31, 15},
			Format:  style1,
			Content: "单位名称"},
		{
			IsMerge: true,
			Xzone:   []string{"c"},
			Yzone:   []int{1, 2},
			Xwidth:  []float64{48},
			Yheight: []float64{31, 15},
			Format:  style2,
			Content: "统计内容（注：统计数据为执法主体下发的任务数）"},
		{
			IsMerge: true,
			Xzone:   []string{"d"},
			Yzone:   []int{1, 2},
			Xwidth:  []float64{11},
			Yheight: []float64{31, 15},
			Format:  style1,
			Content: "累计数据"},
		{
			IsMerge: false,
			Xzone:   []string{"e"},
			Yzone:   []int{1},
			Xwidth:  []float64{16},
			Yheight: []float64{31},
			Format:  style1,
			Content: "7月19日"},
		{
			IsMerge: false,
			Xzone:   []string{"e"},
			Yzone:   []int{2},
			Xwidth:  []float64{16},
			Yheight: []float64{15},
			Format:  style1,
			Content: "当日新增数量",
		},
	}
	return &title
}

func GetCityContent(style1, style2, start_index_y int) *[]Cell {
	c_columns_content := []string{"下发评查任务数量", "接收评查任务数量", "上报案件数量",
		"上报案卷资料数量", "评查案卷数量", "评查单确认已完成数量", "已生成案卷评查结果数量"}
	c_columns_style := []int{style2, style1, style2, style1, style1, style1, style2}

	a_columns := []Cell{}
	b_columns := []Cell{
		{
			IsMerge: true,
			Xzone:   []string{"b"},
			Yzone:   []int{3, 4, 5, 6, 7, 8, 9},
			Xwidth:  []float64{10},
			Yheight: []float64{16.5, 16.5, 16.5, 16.5, 16.5, 16.5, 16.5},
			Format:  style1,
			Content: "省司法厅",
		},
	}
	c_columns := []Cell{}
	d_columns := []Cell{}
	e_columns := []Cell{}
	citycontent := []Cell{}
	for i := start_index_y; i < start_index_y+7; i++ {
		a_columns = append(a_columns, Cell{
			IsMerge: false,
			Xzone:   []string{"a"},
			Yzone:   []int{i},
			Xwidth:  []float64{10},
			Yheight: []float64{16.5},
			Format:  style1,
			Content: strconv.Itoa(i - start_index_y + 1),
		})
		c_columns = append(c_columns, Cell{
			IsMerge: false,
			Xzone:   []string{"c"},
			Yzone:   []int{i},
			Xwidth:  []float64{48},
			Yheight: []float64{16.5},
			Format:  c_columns_style[i-start_index_y],
			Content: c_columns_content[i-start_index_y],
		})
		d_columns = append(d_columns, Cell{
			IsMerge: false,
			Xzone:   []string{"d"},
			Yzone:   []int{i},
			Xwidth:  []float64{11},
			Yheight: []float64{16.5},
			Format:  style1,
		})
		e_columns = append(e_columns, Cell{
			IsMerge: false,
			Xzone:   []string{"e"},
			Yzone:   []int{i},
			Xwidth:  []float64{16},
			Yheight: []float64{16.5},
			Format:  style1,
		})
	}
	citycontent = append(citycontent, a_columns...)
	citycontent = append(citycontent, b_columns...)
	citycontent = append(citycontent, c_columns...)
	citycontent = append(citycontent, d_columns...)
	citycontent = append(citycontent, e_columns...)
	return &citycontent
}
