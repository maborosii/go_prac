package setting

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCity(t *testing.T) {

	Convey("GetCity", t, func() {
		Convey("确认解析得到的配置值:name", func() {
			name, _ := GetCity()
			So(name, ShouldResemble, []string{"省", "梅州市", "惠州市", "中山市"})
			// So(codes, ShouldResemble, []string{"114400000069401674", "11441400007208381Y", "11441300007188069T", "114420000073327379"})
		})
		Convey("确认解析得到的配置值:code", func() {
			_, codes := GetCity()
			// So(name, ShouldResemble, []string{"省", "梅州市", "惠州市", "中山市"})
			So(codes, ShouldResemble, []string{"114400000069401674", "11441400007208381Y", "11441300007188069T", "114420000073327379"})
		})
	})
}

// func TestGetCityFromEmbed(t *testing.T) {

// 	Convey("GetCityFromEmbed", t, func() {
// 		Convey("确认解析得到的配置值:name", func() {
// 			name, _ := GetCityFromEmbed()
// 			So(name, ShouldResemble, []string{"省", "梅州市", "惠州市", "中山市"})
// 		})
// 		Convey("确认解析得到的配置值:code", func() {
// 			_, codes := GetCityFromEmbed()
// 			// So(name, ShouldResemble, []string{"省", "梅州市", "惠州市", "中山市"})
// 			So(codes, ShouldResemble, []string{"114400000069401674", "11441400007208381Y", "11441300007188069T", "114420000073327379"})
// 		})
// 	})
// }
