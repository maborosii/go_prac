package sqlinfo

import (
	"io/ioutil"
)

type noparamssql struct {
	sentence  string
	sqlparams []interface{}
}

func (ds *noparamssql) BuildSqlParams() {
	/* 构造dailycase的sql参数
	 */

	demo := make([]interface{}, 0)
	ds.sqlparams = demo
}

func (ds *noparamssql) BuildSql(path string) {
	dynamicsql, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	sqlcontent := string(dynamicsql)
	ds.sentence = sqlcontent
}

func (ds *noparamssql) GetSql() string {
	return ds.sentence

}

func (ds *noparamssql) GetParams() []interface{} {
	return ds.sqlparams
}

func Newnoparamssql() DynamicSql {
	return &noparamssql{}
}
