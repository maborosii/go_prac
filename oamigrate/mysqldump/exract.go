// package mysqldump
package mysqldump

import (
	"embed"
	. "excelfromdb/locallog"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

//go:embed export.yaml
var exportfile embed.FS

func GetConfig() (map[string]string, []string) {

	tables := []string{}
	dbconfig := map[string]string{}
	config := viper.New()
	config.SetConfigType("yaml")

	iofile, err := exportfile.ReadFile("export.yaml")

	if err != nil {
		Log.Fatal(err)
	}
	if err = config.ReadIOInConfig(iofile); err != nil {
		// 这里修改了viper的源码，增加了读取fs.file的选项
		Log.Fatal(err)
	}
	dbconfig["host"] = config.Get("dbconfig.host").(string)
	dbconfig["user"] = config.Get("dbconfig.user").(string)
	dbconfig["database"] = config.Get("dbconfig.database").(string)
	dbconfig["password"] = config.Get("dbconfig.password").(string)
	dbconfig["port"] = strconv.Itoa(config.Get("dbconfig.port").(int))

	c := reflect.ValueOf(config.Get("dbconfig.tables"))
	for i := 0; i < c.Len(); i++ {
		ele := c.Index(i)
		tables = append(tables, ele.Interface().(string))
	}

	return dbconfig, tables
}
