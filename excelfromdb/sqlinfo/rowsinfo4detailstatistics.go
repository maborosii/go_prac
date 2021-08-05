package sqlinfo

import (
	"fmt"
	"time"
)

type dailycasesql struct {
	noparamssql
}

func (ds *dailycasesql) BuildSqlParams(arg ...interface{}) {
	/* 构造dailycase的sql参数
	 */
	var pre_date, suf_date string

	dates := make([]string, 0)

	t := time.Now()
	if int(t.Weekday()) != 1 {
		pre_date = time.Now().AddDate(0, 0, -1).Format("2006-01-02") + " 00:00:00"
		suf_date = time.Now().AddDate(0, 0, 0).Format("2006-01-02") + " 00:00:00"
	} else {
		pre_date = time.Now().AddDate(0, 0, -3).Format("2006-01-02") + " 00:00:00"
		suf_date = time.Now().AddDate(0, 0, 0).Format("2006-01-02") + " 00:00:00"
	}
	dates = append(dates, pre_date, suf_date)

	fmt.Println(dates)

	paramsArgs := make([]interface{}, len(dates))
	for i := range dates {
		paramsArgs[i] = &dates[i]
	}
	ds.sqlparams = paramsArgs
}

func Newdailycasesql() DynamicSql {
	return &dailycasesql{}
}
