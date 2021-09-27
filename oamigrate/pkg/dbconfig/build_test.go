package dbconfig

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestXormdbconfig(t *testing.T) {

	Convey("ImportConfig", t, func() {
		Convey("查看xorm engine", func() {
			// dbconfigdemo := ImportConfig(s, "apollo_mysql") ///
			a := build()
			So(a, ShouldBeNil)
		})
	})
}
