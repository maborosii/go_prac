package dbconfig

import (
	"embed"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
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
