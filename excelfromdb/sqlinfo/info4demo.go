package sqlinfo

import (
	"io/ioutil"
)

type demosql struct {
	sentence  string
	sqlparams []interface{}
}

func (ds *demosql) BuildSqlParams() {
	/* 构造dailycase的sql参数
	 */

	a := make([]interface{}, 1)
	b := "1"

	a[0] = &b

	ds.sqlparams = a
}

func (ds *demosql) BuildSql(path string) {
	dynamicsql, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	sqlcontent := string(dynamicsql)
	ds.sentence = sqlcontent
}

func (ds *demosql) GetSql() string {
	return ds.sentence

}

func (ds *demosql) GetParams() []interface{} {
	return ds.sqlparams
}

func Newdemosql() DynamicSql {
	return &demosql{}
}
