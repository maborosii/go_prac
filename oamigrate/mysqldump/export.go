package mysqldump

import (
	. "excelfromdb/locallog"
	"io/ioutil"
	"os/exec"
	"path"

	"golang.org/x/sync/errgroup"
)

func export(tablename string, prepath string) error {

	dbconfig, _ := GetConfig()

	var cmd *exec.Cmd
	argv := []string{"--ssl-mode=disable", "--single-transaction", "-h" + dbconfig["host"],
		"-u" + dbconfig["user"], "-p" + dbconfig["password"],
		"-P" + dbconfig["port"], "--databases", dbconfig["database"], "--tables", tablename}
	cmd = exec.Command("mysqldump", argv...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Log.Fatal(err)
		return err
	}

	if err := cmd.Start(); err != nil {
		Log.Fatal(err)
		return err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		Log.Fatal(err)
		return err
	}
	// now := time.Now().Format("20060102150405")

	// backupPath := tablename + "_" + now + ".sql"
	backupPath := path.Join(prepath, tablename+".sql")
	Log.Info("backup path is ", prepath)
	err = ioutil.WriteFile(backupPath, bytes, 0644)

	if err != nil {
		Log.Fatal(err)
		return err
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

	_, tables := GetConfig()
	for _, table := range tables {

		// 避免协程只引用最后一个变量，创建一个闭包函数的上下文变量
		table := table
		group.Go(func() error {
			err := export(table, prepath)
			if err != nil {
				Log.Fatal(table, " export failed")
				return err
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
