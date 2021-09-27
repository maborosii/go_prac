package sqlinfo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDeltanumcasesql(t *testing.T) {

	Convey("BuildSqlParams", t, func() {
		Convey("deltanumcasesql：构建的参数与期望得到的参数一致，且同为[]interface{}", func() {
			s := deltanumcasesql{}
			s.BuildSqlParams("a")
			So(s.sqlparams, ShouldResemble, []interface{}{"2021-08-04", "a", "2021-08-04",
				"a", "2021-08-04", "a", "2021-08-04", "a", "2021-08-04", "a", "2021-08-04", "a", "2021-08-04", "a"})
		})
	})
}
