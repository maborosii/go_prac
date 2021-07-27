package sqlinfo

import (
	"io/ioutil"
)

type DynamicSql struct {
	Sentence  string
	Sqlparams []interface{}
}

func NewDynamicSql(sqlpath string, args ...interface{}) *DynamicSql {
	b, err := ioutil.ReadFile(sqlpath)
	if err != nil {
		panic(err)
	}

	sql := string(b)
	ds := &DynamicSql{Sentence: sql, Sqlparams: args}
	return ds
}
