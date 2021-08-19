package main

import (
	"database/sql"
	"excelfromdb/dbconfig"
	"excelfromdb/excelops"
	. "excelfromdb/locallog"
	"excelfromdb/shell"
	"excelfromdb/sqlinfo"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func Publictiy() {
	newexcel := "采集，公式统计.xlsx"
	node_info := "law_collection"
	node_six := "law_team"

	// 修改excelize源码，添加了读取embed.FS的方法
	f, err := excelize.OpenFileEmbed(publicityTemplate, string(publictiy_template))

	if err != nil {
		Log.Fatal(err)
	}

	// sheetname: 02、公示六大类统计
	local_config_six := dbconfig.ImportConfig(configfile, node_six)
	db_six := local_config_six.InitConnector()
	defer func() {
		local_config_six.CloseConnector(db_six)
	}()

	err = publicitysix(f, db_six, local_config_six)
	if err != nil {
		Log.Error("sheetname: 02、公示六大类统计,err: ", err)
	}

	// sheetname: 03、公示采集统计
	local_config_info := dbconfig.ImportConfig(configfile, node_info)
	db_info := local_config_info.InitConnector()
	defer func() {
		local_config_info.CloseConnector(db_info)
	}()

	err = publicityinfo(f, db_info, local_config_info)
	if err != nil {
		Log.Error("sheetname: 03、公示采集统计,err: ", err)
	}

	// sheetname: 04、访问量统计
	err = acessStat(f)
	if err != nil {
		Log.Error("sheetname: 04、访问量统计,err: ", err)
	}

	err = f.SaveAs(newexcel)
	if err != nil {
		Log.Fatal("文件保存失败")
	}

}

func strtoslice(output string) [][]string {
	a := [][]string{}
	b := strings.Split(strings.Trim(output, "\n"), "\n")
	for _, data := range b {
		a = append(a, []string{data})

	}
	return a

}

func publicityinfo(f *excelize.File, db *sql.DB, lconfig *dbconfig.DBConfig) error {

	sqllist := []string{publicityallnum_sql, publicitydeltanum_sql}
	insertplace := []struct {
		x string
		y int
	}{{x: "B", y: 4}, {x: "B", y: 16}}

	mm := sqlinfo.Newnoparamssql()
	mm_chan := make(chan [][]string)

	for i, sql := range sqllist {
		buildsql(sql, mm)
		go getdata4nums(db, mm, lconfig, mm_chan)
		excelops.GenerateExcelByRowsAnyPlace(f, insertplace[i].x, publicity_sheet, insertplace[i].y, (<-mm_chan)[1:])
	}

	return nil

}

func acessStat(f *excelize.File) error {
	output, err := shell.AccessStatOutput()
	if err != nil {
		return err
	}
	accesslogslice := strtoslice(output)

	insertplace := []struct {
		x string
		y int
	}{{x: "B", y: 5}}

	for _, place := range insertplace {
		excelops.GenerateExcelByRowsAnyPlace(f, place.x, acceessstat_sheet, place.y, accesslogslice)

	}

	return nil

}

func publicitysix(f *excelize.File, db *sql.DB, lconfig *dbconfig.DBConfig) error {

	sqllist := []string{sixdeltanum_sql, sixdeltanum_pub_sql, sixallnum_sql,
		sixallnum_pub_sql, publicityallnum_dept_sql, publicityallnum_user_sql}
	insertplace := []struct {
		x string
		y int
	}{{x: "C", y: 4}, {x: "C", y: 11}, {x: "C", y: 18},
		{x: "C", y: 25}, {x: "C", y: 32}, {x: "C", y: 37}}

	mm := sqlinfo.Newnoparamssql()
	mm_chan := make(chan [][]string)

	for i, sql := range sqllist {
		buildsql(sql, mm)
		go getdata4nums(db, mm, lconfig, mm_chan)
		excelops.GenerateExcelByRowsAnyPlace(f, insertplace[i].x, six_sheet, insertplace[i].y, (<-mm_chan)[1:])
	}

	return nil

}
