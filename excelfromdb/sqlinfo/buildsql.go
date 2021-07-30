package sqlinfo

// type DynamicSql struct {
// 	Sentence  string
// 	Sqlparams []interface{}
// }

type DynamicSql interface {
	BuildSqlParams()
	BuildSql(path string)
	GetParams() []interface{}
	GetSql() string
}

// func NewDynamicSql(sqlpath string, args ...interface{}) *DynamicSql {

// 	dynamicsql, err := ioutil.ReadFile(sqlpath)

// 	if err != nil {
// 		panic(err)
// 	}
// 	sqlcontent := string(dynamicsql)

// 	ds := &DynamicSql{Sentence: sqlcontent, Sqlparams: args}
// 	fmt.Println(ds)
// 	return ds
// }
