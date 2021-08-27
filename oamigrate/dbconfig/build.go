package dbconfig

import (
	"embed"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	_ "xorm.io/xorm/schemas"
)

// var engine *xorm.Engine

//go:embed db.conf
var testconf embed.FS

func build() error {
	configfile := Newconfigfile()(testconf, "db.conf")
	local_config := ImportConfig(configfile, "local")

	engine, err := xorm.NewEngine("mysql", local_config.BuildConnectString())
	if err != nil {
		return err
	}
	fmt.Println(engine)
	return nil

}

// func DumpTable(filename string) error {
// 	configfile := Newconfigfile()(testconf, "db.conf")
// 	local_config := ImportConfig(configfile, "local_read")

// 	engine, err := xorm.NewEngine("mysql", local_config.BuildConnectString())
// 	if err != nil {
// 		return err
// 	}
// 	engine.ShowSQL(true)
// 	var tables []*schemas.Table
// 	if tables, err = engine.DBMetas(); err != nil {
// 		return err
// 	}
// 	if err = engine.DumpTablesToFile(tables, filename, schemas.MYSQL); err != nil {
// 		return err
// 	}
// 	return nil
// }

func ImportTable(filename string) error {
	configfile := Newconfigfile()(testconf, "db.conf")
	local_config := ImportConfig(configfile, "local_test")

	engine, err := xorm.NewEngine("mysql", local_config.BuildConnectString())
	if err != nil {
		return err
	}
	engine.ShowSQL(true)
	_, err = engine.ImportFile(filename)
	if err != nil {
		return err
	}
	return nil
}
