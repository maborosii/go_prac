package exporter

import (
	"io/ioutil"
	. "oamigrate/pkg/log"
	"os/exec"
	"path"

	c "oamigrate/config"

	"golang.org/x/sync/errgroup"
)

func export(tablename string, prepath string) error {

	dbConfig := c.ExportDbConfig
	Log.Info(dbConfig)

	var cmd *exec.Cmd
	// local setting
	argv := []string{"--ssl-mode=disable", "--single-transaction", "-h" + dbConfig["host"],
		"-u" + dbConfig["user"], "-p" + dbConfig["password"],
		"-P" + dbConfig["port"], "--databases", dbConfig["database"], "--tables", tablename}
	Log.Info(argv)

	// prod setting
	// argv := []string{"--single-transaction", "-h" + dbConfig["host"],
	// 	"-u" + dbConfig["user"], "-p" + dbConfig["password"],
	// 	"-P" + dbConfig["port"], "--databases", dbConfig["database"], "--tables", tablename}
	cmd = exec.Command("mysqldump", argv...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		Log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		Log.Fatal(err)
	}

	backupPath := path.Join(prepath, tablename+".sql")
	Log.Info("backup path is ", prepath)
	err = ioutil.WriteFile(backupPath, bytes, 0644)

	if err != nil {
		Log.Fatal(err)
	}
	return nil
}

func FullExport() error {
	// 使用 sync/errgroup 进行goroutine协程错误控制
	prepath, err := mkpath()
	if err != nil {
		Log.Fatal(err)
	}

	group := new(errgroup.Group)

	tables := c.Tables
	for _, table := range tables {

		// 避免协程只引用最后一个变量，创建一个闭包函数的上下文变量
		table := table
		group.Go(func() error {
			err := export(table, prepath)
			if err != nil {
				Log.Fatal(table, " export failed")
			}
			Log.Info(table, " export success")
			return err
		})

	}

	if err := group.Wait(); err != nil {
		Log.Fatal(err)
	} else {
		Log.Info("all table export success")
	}

	return nil
}
