package sum

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSum(t *testing.T) {
	Convey("浮点型数组求和Sum", t, func() {
		Convey("should return 7.27", func() {
			a := []float64{1.2, 3.4, 2.67}
			So(Sum(a), ShouldEqual, 7.27)
		})

		Convey("Sum should return 0.00", func() {
			a := []float64{}
			So(Sum(a), ShouldEqual, 0.00)
		})

		Convey("Sum should return 0", func() {
			a := []float64{0}
			So(Sum(a), ShouldEqual, 0)
		})
	})
}
func TestAvg(t *testing.T) {
	Convey("浮点型数组求均值Avg", t, func() {
		Convey("should return 7.27", func() {
			a := []float64{1.2, 3.4, 2.67}
			So(Avg(a), ShouldEqual, 2.4233333333333333)
		})

		Convey("Avg should return 0.00", func() {
			a := []float64{}
			So(Avg(a), ShouldEqual, 0.00)
		})

		Convey("Avg should return 0", func() {
			a := []float64{0}
			So(Avg(a), ShouldEqual, 0)
		})
	})
}
