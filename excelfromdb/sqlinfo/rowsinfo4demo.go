package sqlinfo

type demosql struct {
	noparamssql
}

func (ds *demosql) BuildSqlParams(arg ...interface{}) {
	/* 构造dailycase的sql参数
	 */

	a := make([]interface{}, 1)
	b := "1"

	a[0] = &b

	ds.sqlparams = a
}

func Newdemosql() DynamicSql {
	return &demosql{}
}
