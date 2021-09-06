package importdata

import (
	"embed"

	db "oamigrate/dbconfig"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	_ "xorm.io/xorm/schemas"
)

//go:embed db.conf
var testconf embed.FS

func ImportTable(filename string) error {
	configfile := db.Newconfigfile()(testconf, "db.conf")
	local_config := db.ImportConfig(configfile, "local_test")

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
