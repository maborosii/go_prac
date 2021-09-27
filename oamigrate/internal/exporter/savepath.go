package exporter

import (
	"flag"
	"os"
	"path/filepath"
)

func mkpath() (string, error) {

	prepath, err := getpath()
	if err != nil {
		return "", err
	}
	path := filepath.Join(prepath, "sql/oamigrate")
	isDirExists := func(path string) bool {
		_, err = os.Stat(path)
		return err == nil || os.IsExist(err)
	}

	if isDirExists(path) {
		return path, nil
	}
	err = os.MkdirAll(path, 0777)
	if err != nil {
		return "", err
	}
	return path, nil

}

func getpath() (string, error) {
	//默认路径为执行命令的位置

	allargs := os.Args
	if len(allargs) == 1 {
		default_path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return "", err
		}
		return default_path, nil
	}
	customize_path := flag.String("path", os.Args[0], "sqlfile's location")
	flag.Parse()
	return *customize_path, nil
}
