package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// init
	// node := "apollo_mysql"
	// configfile := &config.ConfigFile{FileName: "config/db.conf"}
	// sql := "SELECT * from ServerConfig"

	// dbconfig := config.ImportConfig(configfile, node)

	// db := dbconfig.InitConnector()

	// fmt.Println(dbconfig.QuerySql(db, sql))
	intIndexX := 52
	var columns_arrs = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var strIndexX string
	if intIndexX > 26 {
		if intIndexX%26 == 0 {

		}
		strIndexX := columns_arrs[intIndexX/26-1] + columns_arrs[intIndexX%26-1]
		fmt.Println(strIndexX)
	}

}
