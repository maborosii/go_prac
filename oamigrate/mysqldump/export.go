package mysqldump

import (
	. "excelfromdb/locallog"
	"io/ioutil"
	"os/exec"
	"time"

	"golang.org/x/sync/errgroup"
)

func export(tablename string) (string, error) {

	dbconfig, _ := GetConfig()

	var cmd *exec.Cmd
	argv := []string{"--ssl-mode=disable", "--single-transaction", "-h" + dbconfig["host"],
		"-u" + dbconfig["user"], "-p" + dbconfig["password"],
		"-P" + dbconfig["port"], "--databases", dbconfig["database"], "--tables", tablename}
	cmd = exec.Command("mysqldump", argv...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Log.Fatal(err)
		return "", err
	}

	if err := cmd.Start(); err != nil {
		Log.Fatal(err)
		return "", err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		Log.Fatal(err)
		return "", err
	}
	now := time.Now().Format("20060102150405")

	backupPath := tablename + "_" + now + ".sql"
	err = ioutil.WriteFile(backupPath, bytes, 0644)

	if err != nil {
		Log.Fatal(err)
		return "", err
	}
	return backupPath, nil
}

func FullExport() ([]string, error) {
	// 使用 sync/errgroup 进行goroutine协程错误控制

	group := new(errgroup.Group)
	path := make(chan *string)
	tablepaths := []string{}

	_, tables := GetConfig()
	for _, table := range tables {

		// 避免协程只引用最后一个变量，创建一个闭包函数的上下文变量
		table := table
		group.Go(func() error {
			filepath, err := export(table)
			path <- &filepath

			if err != nil {
				Log.Fatal(table, " export failed")
				return err
			}
			Log.Info(table, " export success")
			return err
		})

	}

	// tablepaths = append(tablepaths, *(<-path))
	// Log.Info(tablepaths)
	if err := group.Wait(); err != nil {
		Log.Fatal(err)
	} else {
		Log.Info("all table export success")
		tablepaths = append(tablepaths, *(<-path))
		Log.Info(tablepaths)
	}

	close(path)
	return tablepaths, nil
}
