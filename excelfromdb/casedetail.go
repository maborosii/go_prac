package main

import (
	"excelfromdb/dbconfig"
	"excelfromdb/excelops"
	"excelfromdb/sqlinfo"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/go-sql-driver/mysql"
)

func CaseDetail() {
	dbconfigpath := "/root/golang/dbconfig/db.conf"
	sheetname := "Sheet1"
	detail_sqlpath := "/root/golang/sqlinfo/casestatistics_detail.sql"
	savexlsx := "案卷每日评查情况.xlsx"
	node := "law_case_review"

	f := excelize.NewFile()
	index := f.NewSheet(sheetname)

	configfile := &dbconfig.ConfigFile{FileName: dbconfigpath}

	mm := sqlinfo.Newdailycasesql()
	buildsql(detail_sqlpath, &mm)

	local_config := dbconfig.ImportConfig(configfile, node)
	db := local_config.InitConnector()
	defer func() {
		local_config.CloseConnector(db)
	}()

	allrecords := local_config.QuerySql(db, mm.GetSql(), mm.GetParams()...)
	excelops.GenerateExcelByRows(f, sheetname, allrecords)

	f.SetActiveSheet(index)
	if err := f.SaveAs(savexlsx); err != nil {
		panic(err)
	}
}
