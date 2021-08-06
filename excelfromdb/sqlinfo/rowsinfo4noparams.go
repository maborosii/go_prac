package sqlinfo

type noparamssql struct {
	sentence  string
	sqlparams []interface{}
}

func (ds *noparamssql) BuildSqlParams(arg ...interface{}) {
	/* 构造dailycase的sql参数
	 */

	demo := make([]interface{}, 0)
	ds.sqlparams = demo
}

func (ds *noparamssql) BuildSql(sql string) {
	ds.sentence = sql
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
