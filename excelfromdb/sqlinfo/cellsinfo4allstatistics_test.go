package sqlinfo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAllnumcasesql(t *testing.T) {

	Convey("BuildSqlParams", t, func() {
		// shouldFunc := func(actual interface{}, expected ...interface{}) string {
		// 	if !reflect.DeepEqual(actual, expected[0]) { // 因为slice不能比较直接，借助反射包中的方法比较
		// 		return fmt.Sprintf("excepted:%v, got:%v", expected, actual) // 测试失败输出错误提示
		// 	}
		// 	return ""
		// }
		Convey("allnumcasesql：构建的参数与期望得到的参数一致，且同为[]interface{}", func() {
			s := allnumcasesql{}
			s.BuildSqlParams("114400000069401674")
			So(s.sqlparams, ShouldResemble, []interface{}{"114400000069401674",
				"114400000069401674", "114400000069401674", "114400000069401674", "114400000069401674", "114400000069401674", "114400000069401674"})
		})
	})
}
