package importdata

import (
	"embed"
	. "excelfromdb/locallog"
	"fmt"
	"os"
	"path/filepath"

	db "oamigrate/dbconfig"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
	"xorm.io/xorm"
	_ "xorm.io/xorm/schemas"
)

//go:embed db.conf
var testconf embed.FS

func ImportTable() error {
	configfile := db.Newconfigfile()(testconf, "db.conf")
	local_config := db.ImportConfig(configfile, "local_test")

	engine, err := xorm.NewEngine("mysql", local_config.BuildConnectString())
	if err != nil {
		return err
	}
	// engine.ShowSQL(true)
	// _, err = engine.ImportFile("/home/bonbon/golang/github.com/maborosii/oamigrate/sql/oamigrate/wf_step_def.sql")
	// if err != nil {
	// 	Log.Error(" import failed")
	// 	return err
	// }
	// Log.Info(" import success")
	// return nil

	group := new(errgroup.Group)
	files, err := getSqlFile("/home/bonbon/golang/github.com/maborosii/oamigrate/sql/oamigrate")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		// 避免协程只引用最后一个变量，创建一个闭包函数的上下文变量
		file := file
		group.Go(func() error {

			//ImportFile 线程不安全
			_, err = engine.ImportFile(file)
			if err != nil {
				Log.Error(file, " import failed")
				return err
			}
			Log.Info(file, " import success")
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		Log.Error(err)
	} else {
		Log.Info("all table import success")
	}
	return nil
}

func getSqlFile(path string) ([]string, error) {
	var files []string

	err := filepath.Walk(path, func(pathRoot string, info os.FileInfo, err error) error {
		files = append(files, pathRoot)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files[1:], nil
}
