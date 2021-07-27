package main

import (
	"excelfromdb/dbconfig"
	"excelfromdb/excelops"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SqlParmsForDailyCase() []string {
	/* 构造dailycase的sql参数
	 */
	var pre_date, suf_date string

	dates := make([]string, 0)

	t := time.Now()
	if int(t.Weekday()) != 1 {
		pre_date = time.Now().AddDate(0, 0, -1).Format("2006-01-02") + " 00:00:00"
		suf_date = time.Now().AddDate(0, 0, 0).Format("2006-01-02") + " 00:00:00"
	} else {
		pre_date = time.Now().AddDate(0, 0, -3).Format("2006-01-02") + " 00:00:00"
		suf_date = time.Now().AddDate(0, 0, 0).Format("2006-01-02") + " 00:00:00"
	}
	dates = append(dates, pre_date, suf_date)

	fmt.Println(dates)

	return dates

}

func main() {

	// init
	node := "apollo_mysql"
	configfile := &dbconfig.ConfigFile{FileName: "dbconfig/db.conf"}
	// sql := "SELECT * from ServerConfig"

	sqlinfo := NewDynamicSql()

	lcoal_config := dbconfig.ImportConfig(configfile, node)

	db := lcoal_config.InitConnector()

	allrecords := lcoal_config.QuerySql(db, sql)

	excelops.GenerateExcelByRows(allrecords)
	SqlParmsForDailyCase()
}
