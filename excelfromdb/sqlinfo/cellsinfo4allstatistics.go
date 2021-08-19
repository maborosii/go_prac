package sqlinfo

type allnumcasesql struct {
	noparamssql
}

func (ds *allnumcasesql) BuildSqlParams(arg ...interface{}) {

	a := make([]interface{}, 7)
	for i := range a {
		a[i] = arg[0]
	}

	ds.sqlparams = a
}

func Newallnumcasesql() DynamicSql {
	return &allnumcasesql{}
}
