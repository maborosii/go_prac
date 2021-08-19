package sqlinfo

import "time"

type deltanumcasesql struct {
	noparamssql
}

func (ds *deltanumcasesql) BuildSqlParams(arg ...interface{}) {

	a := make([]interface{}, 14)
	for i := 1; i < 14; i += 2 {
		a[i] = arg[0]
		a[i-1] = time.Now().AddDate(0, 0, 0).Format("2006-01-02")

	}

	ds.sqlparams = a
}

func Newdeltanumcasesql() DynamicSql {
	return &deltanumcasesql{}
}
