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
	importPath := c.ImportPath

	// 导入数据前置处理，rename表，删除多余备份表
	enginePre, err := xorm.NewEngine("mysql", dbConfig.BuildConnectString())
	if err != nil {
		return err
	}
	if err = preTableHook(enginePre); err != nil {
		Log.Error("preTableHook occur err")
		return err
	}

	// 导入数据
	engineImp, err := xorm.NewEngine("mysql", dbConfig.BuildConnectString())
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var countErr = 0
	files, err := getSqlFile(importPath)

	if err != nil {
		Log.Error(err)
		return err
	}
	if len(files) < 1 {
		Log.Error("sql file is not exists")
		return errors.New("sql file is not exists")
	}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			mutex.Lock()
			Log.Info(file, " is processing")
			_, err := engineImp.ImportFile(file)
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

	// 导出数据后置处理，修改各项目passid
	enginePost, err := xorm.NewEngine("mysql", dbConfig.BuildConnectString())
	if err != nil {
		return err
	}
	if err = postTableHook(enginePost); err != nil {
		Log.Error(err)
		return err
	}
	return nil
}

func getSqlFile(path string) ([]string, error) {
	var files []string

	//检查sql文件路径是否存在
	isDirExists := func(path string) bool {
		_, err := os.Stat(path)
		return err == nil || os.IsExist(err)
	}

	if !isDirExists(path) {
		Log.Error(path, " is not exists")
		return nil, errors.New("path is not exists")
	}

	err := filepath.Walk(path, func(pathRoot string, info os.FileInfo, err error) error {
		files = append(files, pathRoot)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files[1:], nil
}
