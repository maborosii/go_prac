package sqlinfo

type DynamicSql interface {
	BuildSqlParams(arg ...interface{})
	BuildSql(path string)
	GetParams() []interface{}
	GetSql() string
}
