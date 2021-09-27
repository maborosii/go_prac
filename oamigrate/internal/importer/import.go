package importer

import (
	"errors"
	. "oamigrate/pkg/log"
	"os"
	"path/filepath"
	"sync"

	c "oamigrate/config"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	_ "xorm.io/xorm/schemas"
)

func ImportTable() error {
	dbConfig := c.ImportConfig
	enginePre, err := xorm.NewEngine("mysql", dbConfig.BuildConnectString())
	if err != nil {
		return err
	}
	if err = preTableHook(enginePre); err != nil {
		Log.Error("preTableHook occur err")
		return err
	}

	engineImp, err := xorm.NewEngine("mysql", dbConfig.BuildConnectString())
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
			_, err = engineImp.ImportFile(file)
			if err != nil {
				Log.Error(file, " import failed. ", "err info：", err)
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
