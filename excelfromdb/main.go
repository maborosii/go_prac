package main

import (
	"excelfromdb/dbconfig"
	"excelfromdb/excelops"
	"excelfromdb/sqlinfo"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// init

	// -----------------  local_config
	node := "apollo_mysql"
	configfile := &dbconfig.ConfigFile{FileName: "dbconfig/db.conf"}

	mm := sqlinfo.Newdemosql()

	mm.BuildSqlParams()
	mm.BuildSql("sqlinfo/demo.sql")

	// -----------------  server_config
	// node := "law_case_review"
	// configfile := &dbconfig.ConfigFile{FileName: "/root/golang/dbconfig/db.conf"}

	// mm := sqlinfo.Newdailycasesql()

	// mm.BuildSqlParams()
	// mm.BuildSql("/root/golang/sqlinfo/dailycaseinfo.sql")

	local_config := dbconfig.ImportConfig(configfile, node)

	db := local_config.InitConnector()

	allrecords := local_config.QuerySql(db, mm.GetSql(), mm.GetParams()...)

	excelops.GenerateExcelByRows(allrecords, "books.xlsx")
}
