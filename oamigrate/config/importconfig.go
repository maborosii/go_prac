package config

import (
	"embed"
	"flag"
	db "oamigrate/pkg/dbconfig"
	"os"
	"path/filepath"
	// . "oamigrate/pkg/log"
)

//go:embed db.conf
var testConf embed.FS

var Node, ImportPath = func() (string, string) {
	rootPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", ""
	}
	defaultPath := filepath.Join(rootPath, "sql/oamigrate")
	customizePath := flag.String("path", defaultPath, "sqlfile's location")
	node := flag.String("node", "local_test", `city name: sheng, jiangmen, zhongshan, zhaoqing, foshan, yunfu, huizhou, shantou, yangjiang, maoming, zhuhai, dongguan`)
	flag.Parse()
	return *node, *customizePath
}()

var ImportConfig = func(node string) *db.DBConfig {
	return db.ImportConfig(db.Newconfigfile()(testConf, "db.conf"), node)

}(Node)

var CityPassId = map[string]string{
	"default":     "gdxzzfpc",
	"prepro":      "gdxzzfpc",
	"sheng":       "gdxzzfpc",
	"jiangmen":    "rz_gdxzzfptjmptpc",
	"zhongshan":   "rz_gdxzpc",
	"zhaoqing":    "rz_gdzfxxjd",
	"foshan":      "rz_fszfxx",
	"yunfu":       "rz_xzzfyfs",
	"huizhou":     "rz_hzxzzf",
	"shantou":     "rz_gdsxzz",
	"yangjiang":   "rz_xzyj",
	"maoming":     "rz_mmsxzzf",
	"zhuhai":      "rz_xzzfpc",
	"dongguan":    "r_rz_xzzf",
	"dongguanpre": "r_rz_xzzf",
}
