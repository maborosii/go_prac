package config

import (
	"embed"
	db "oamigrate/pkg/dbconfig"
	// . "oamigrate/pkg/log"
)

//go:embed db.conf
var testConf embed.FS

var node = "local_test"
var ImportConfig = func(node string) *db.DBConfig {
	return db.ImportConfig(db.Newconfigfile()(testConf, "db.conf"), node)

}(node)
