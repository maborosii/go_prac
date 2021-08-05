package main

import (
	"database/sql"
	"excelfromdb/dbconfig"
	ew "excelfromdb/excelwriting"
	"excelfromdb/setting"
	"excelfromdb/sqlinfo"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/go-sql-driver/mysql"
)

func buildsql(sqlfile string, mm *sqlinfo.DynamicSql, sqlparams ...interface{}) {
	(*mm).BuildSqlParams(sqlparams...)
	(*mm).BuildSql(sqlfile)
}

func getdata4nums(db *sql.DB, mm *sqlinfo.DynamicSql, config *dbconfig.DBConfig, mm_chan chan *[][]string) {
	records := config.QuerySql(db, (*mm).GetSql(), (*mm).GetParams()...)
	mm_chan <- &records

}

func CaseNums() {
	dbconfigpath := "/root/golang/dbconfig/db.conf"
	settingpath := "/root/golang/setting"
	sheetname := "Sheet1"
	allnum_sqlpath := "/root/golang/sqlinfo/casestatistics_allnum.sql"
	deltanum_sqlpath := "/root/golang/sqlinfo/casestatistics_deltanum.sql"
	savexlsx := "案卷评查数据统计表.xlsx"
	node := "law_case_review"

	f := excelize.NewFile()
	index := f.NewSheet(sheetname)

	citieslist, codelist := setting.GetCity(settingpath)
	inputareacell := ew.NumsBuildFormat(f, sheetname, citieslist)

	configfile := &dbconfig.ConfigFile{FileName: dbconfigpath}
	local_config := dbconfig.ImportConfig(configfile, node)
	db := local_config.InitConnector()
	defer func() {
		local_config.CloseConnector(db)
	}()

	mm := sqlinfo.Newallnumcasesql()
	aa := sqlinfo.Newdeltanumcasesql()

	mm_chan := make(chan *[][]string)
	aa_chan := make(chan *[][]string)

	for i, code := range codelist {
		buildsql(allnum_sqlpath, &mm, code)
		buildsql(deltanum_sqlpath, &aa, code)
		go getdata4nums(db, &mm, local_config, mm_chan)
		go getdata4nums(db, &aa, local_config, aa_chan)
		ew.InputAllNums(&(*inputareacell)[i][0], sheetname, f, <-mm_chan)
		ew.InputDeltaNums(&(*inputareacell)[i][1], sheetname, f, <-aa_chan)

	}

	f.SetActiveSheet(index)
	if err := f.SaveAs(savexlsx); err != nil {
		panic(err)
	}
}
