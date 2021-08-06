package dbconfig

import (
	"embed"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//go:embed db.conf
var democonf embed.FS

func TestDBconfig(t *testing.T) {

	Convey("ImportConfig", t, func() {
		Convey("embedFS导入ini", func() {
			s := &configfile{democonf, "db.conf"}
			dbconfigdemo := ImportConfig(s, "apollo_mysql") ///
			So(dbconfigdemo, ShouldResemble, &DBConfig{Host: "172.30.65.163", User: "root", Password: "123456", Database: "ApolloConfigDB", CharSet: "utf8", Port: 50006})
		})
	})
}
