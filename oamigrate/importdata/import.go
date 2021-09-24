package importdata

import (
	"embed"
	"errors"
	. "excelfromdb/locallog"
	"os"
	"path/filepath"
	"sync"

	db "oamigrate/dbconfig"

	_ "github.com/go-sql-driver/mysql"
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

	var wg sync.WaitGroup
	var mutex sync.Mutex
	var countErr = 0
	files, err := getSqlFile("/home/bonbon/golang/github.com/maborosii/oamigrate/sql/oamigrate")
	if err != nil {
		Log.Error(err)
		return err
	}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			mutex.Lock()
			Log.Info(file, " is processing")
			_, err = engine.ImportFile(file)
			if err != nil {
				Log.Error(file, " import failed")
				countErr += 1
			} else {
				Log.Info(file, " import success")
			}
			mutex.Unlock()
			wg.Done()
		}(file)
	}
	wg.Wait()
	if countErr != 0 {
		errImportAll := errors.New("some tables import failed")
		Log.Error(errImportAll)
		return errImportAll
	}
	return nil

	// group := new(errgroup.Group)
	// files, err := getSqlFile("/home/bonbon/golang/github.com/maborosii/oamigrate/sql/oamigrate")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, file := range files {
	// 	// 避免协程只引用最后一个变量，创建一个闭包函数的上下文变量
	// 	file := file
	// 	group.Go(func() error {
	// 		//ImportFile 线程不安全
	// 		_, err = engine.ImportFile(file)
	// 		if err != nil {
	// 			Log.Error(file, " import failed")
	// 			return err
	// 		}
	// 		Log.Info(file, " import success")
	// 		return nil
	// 	})
	// }

	// if err := group.Wait(); err != nil {
	// 	Log.Error(err)
	// } else {
	// 	Log.Info("all table import success")
	// }
	// return nil
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
