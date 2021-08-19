/* 在数据库中字段类型为bit类型的查询会有问题，需要进行转码
 */

package dbconfig

import (
	"database/sql"
	"embed"
	"fmt"

	. "excelfromdb/locallog"

	"gopkg.in/ini.v1"
)

type configfile struct {
	fileinmem embed.FS
	name      string
}

type CreateConfigfile func(realfile embed.FS, filename string) *configfile

func Newconfigfile() CreateConfigfile {
	return func(realfile embed.FS, filename string) *configfile {
		return &configfile{fileinmem: realfile, name: filename}
	}

}

type DBConfig struct {
	Host     string `ini:"DB_HOST"`
	User     string `ini:"DB_USER"`
	Password string `ini:"DB_PWD"`
	Database string `ini:"DB_NAME"`
	CharSet  string `ini:"CHARSET"`
	Port     int32  `ini:"DB_PORT"`
}

func ImportConfig(file *configfile, node string) *DBConfig {

	readfile, err := file.fileinmem.Open(file.name)
	if err != nil {
		Log.Fatal(err)
	}

	cfg, err := ini.Load(readfile)

	if err != nil {
		Log.Fatal("Fail to read file: %v", err)

	}
	dbconfig := &DBConfig{}
	err = cfg.Section(node).MapTo(dbconfig)

	if err != nil {
		Log.Fatal("Fail to find section %s: %v", node, err)

	}
	return dbconfig

}

func (dbconfig *DBConfig) BuildConnectString() string {
	/*
		build connector string for mysql
	*/
	connetor_string := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbconfig.User, dbconfig.Password, dbconfig.Host, dbconfig.Port, dbconfig.Database, dbconfig.CharSet)
	return connetor_string
}

func (dbconfig *DBConfig) InitConnector() *sql.DB {

	db, err := sql.Open("mysql", dbconfig.BuildConnectString())

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err != nil {
		Log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		Log.Fatal(err)
	}
	Log.Info("连接已建立")

	return db
}

func (dbconfig *DBConfig) QuerySql(db *sql.DB, dynamicsql string, args ...interface{}) [][]string {
	/* 查询
	 */
	rows, err := db.Query(dynamicsql, args...)

	if err != nil {
		Log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		Log.Fatal(err)
	}

	values := make([]sql.RawBytes, len(columns))
	// values := make([][]byte, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		// 将*sqlRawBytes赋值给空接口，方便传入rows.Scan(dest ...interface{})

		scanArgs[i] = &values[i] //(interface{},*sqlRawBytes)
	}

	totalValues := make([][]string, 0)

	totalValues = append(totalValues, columns)

	defer rows.Close()
	for rows.Next() {

		var s []string

		// 将扫描后的内容赋值给values
		err = rows.Scan(scanArgs...)

		if err != nil {
			Log.Fatal(err)
		}

		for _, v := range values {
			s = append(s, string(v))
		}
		totalValues = append(totalValues, s)
	}
	Log.Info(totalValues)

	if err = rows.Err(); err != nil {
		Log.Fatal(err)
	}
	return totalValues
}

func (dbconfig *DBConfig) CloseConnector(db *sql.DB) {
	err := db.Close()
	if err != nil {
		Log.Fatal(err)
	}
	Log.Info("连接已关闭")
}
