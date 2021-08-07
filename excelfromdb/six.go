package main

import (
	"excelfromdb/dbconfig"
	"excelfromdb/excelops"
	"excelfromdb/sqlinfo"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func PublicitySix() error {

	newexcel := "demo.xlsx"
	dbconfigname := "dbconfig/db.conf"
	node := "law_team"
	sqllist := []string{sixdeltanum_sql, sixdeltanum_pub_sql, sixallnum_sql,
		sixallnum_pub_sql, publicityallnum_dept_sql, publicityallnum_user_sql}
	insertplace := []struct {
		x string
		y int
	}{{x: "C", y: 4}, {x: "C", y: 11}, {x: "C", y: 18},
		{x: "C", y: 25}, {x: "C", y: 32}, {x: "C", y: 37}}

	// 修改excelize源码，添加了读取embed.FS的方法
	f, err := excelize.OpenFileEmbed(publicityTemplate, string(publictiy_template))

	if err != nil {
		return err
	}

	configfile := dbconfig.Newconfigfile(dbconf, dbconfigname)
	local_config := dbconfig.ImportConfig(configfile, node)
	db := local_config.InitConnector()
	defer func() {
		local_config.CloseConnector(db)
	}()

	mm := sqlinfo.Newnoparamssql()
	mm_chan := make(chan [][]string)
	for i, sql := range sqllist {
		buildsql(sql, mm)
		go getdata4nums(db, mm, local_config, mm_chan)
		excelops.GenerateExcelByRowsAnyPlace(f, insertplace[i].x, six_sheet, insertplace[i].y, (<-mm_chan)[1:])
	}

	err = f.SaveAs(newexcel)
	if err != nil {
		fmt.Println("文件保存失败")
		return err
	}
	return nil

}
