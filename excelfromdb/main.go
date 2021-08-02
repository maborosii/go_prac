package main

import (
	_ "excelfromdb/dbconfig"
	"excelfromdb/excelops"
	_ "excelfromdb/sqlinfo"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// init

	// -----------------  local_config
	// node := "apollo_mysql"
	// configfile := &dbconfig.ConfigFile{FileName: "dbconfig/db.conf"}

	// mm := sqlinfo.Newdemosql()

	// mm.BuildSqlParams()
	// mm.BuildSql("sqlinfo/demo.sql")

	// -----------------  server_config
	// node := "law_case_review"
	// configfile := &dbconfig.ConfigFile{FileName: "/root/golang/dbconfig/db.conf"}

	// mm := sqlinfo.Newdailycasesql()

	// mm.BuildSqlParams()
	// mm.BuildSql("/root/golang/sqlinfo/dailycaseinfo.sql")
	f := excelize.NewFile()

	index := f.NewSheet("Sheet1")

	style_title1 := excelops.GetStyle(f, "excelops/style_all_case_title1.json")
	style_title2 := excelops.GetStyle(f, "excelops/style_all_case_title2.json")
	style_content1 := excelops.GetStyle(f, "excelops/style_all_case_content1.json")
	style_content2 := excelops.GetStyle(f, "excelops/style_all_case_content2.json")
	for _, cell := range *excelops.GetTitle(style_title1, style_title2) {
		excelops.Formatting(&cell, "Sheet1", f)
		excelops.Writing(&cell, "Sheet1", f)
	}

	for _, cell := range *excelops.GetCityContent(style_content1, style_content2, 3) {
		excelops.Formatting(&cell, "Sheet1", f)
		excelops.Writing(&cell, "Sheet1", f)
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs("test1.xlsx"); err != nil {
		panic(err)
	}

	// local_config := dbconfig.ImportConfig(configfile, node)

	// db := local_config.InitConnector()

	// allrecords := local_config.QuerySql(db, mm.GetSql(), mm.GetParams()...)

	// excelops.GenerateExcelByRows(allrecords, "books.xlsx")
}
