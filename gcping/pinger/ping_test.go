package pinger

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSocket(t *testing.T) {
	Convey("Socket class 的测试如下:", t, func() {
		Convey("Socket:建立不存在的连接应该返回error不为nil", func() {
			s := Tcpsocket{"127.0.0.1", 8001, time.Duration(time.Second * 1)}
			_, err := s.Getconn()
			So(err, ShouldNotBeNil)
		})
		Convey("Socket:建立可通的连接应该返回error为nil", func() {
			s := Tcpsocket{"127.0.0.1", 8000, time.Duration(time.Second * 1)}
			_, err := s.Getconn()
			So(err, ShouldBeNil)
		})
		Convey("Socket:关闭连接后，再向连接写入数据时应该返回error不为nil", func() {
			s := Tcpsocket{"127.0.0.1", 8000, time.Duration(time.Second * 1)}
			conn, _ := s.Getconn()
			s.Closeconn(conn)
			words := "Hello Server!"
			_, err := (*conn).Write([]byte(words))
			So(err, ShouldNotBeNil)
		})
	})
}

func TestPing(t *testing.T) {
	Convey("Ping class 的测试如下:", t, func() {
		Convey("Ping:建立不存在的连接应该返回error不为nil", func() {
			p := Ping{
				Tcpsocket{"127.0.0.1", 8001, time.Duration(time.Second * 1)},
				0,
				0,
				make([]float64, 0),
			}
			_, err := p.Createsocket()
			So(err, ShouldNotBeNil)
		})
		Convey("Ping:建立可通的连接应该返回error为nil", func() {
			p := Ping{
				Tcpsocket{"127.0.0.1", 8000, time.Duration(time.Second * 1)},
				0,
				0,
				make([]float64, 0),
			}
			_, err := p.Createsocket()
			So(err, ShouldBeNil)
		})
		Convey("Ping:Successedrate 返回成功率(基础为0)", func() {
			p := Ping{
				Tcpsocket{"127.0.0.1", 8000, time.Duration(time.Second * 1)},
				0,
				0,
				make([]float64, 0),
			}
			c, r := p.Successedrate()
			So(c, ShouldEqual, 0)
			So(r, ShouldEqual, "0.00%")
		})
		Convey("Ping:Successedrate 返回成功率(基础不为0)", func() {
			p := Ping{
				Tcpsocket{"127.0.0.1", 8000, time.Duration(time.Second * 1)},
				5,
				5,
				make([]float64, 0),
			}
			c, r := p.Successedrate()
			So(c, ShouldEqual, 10)
			So(r, ShouldEqual, "50.00%")
		})
		Convey("Ping:Statstics 返回最大最小平均值(连接时间列表为空)", func() {
			p := Ping{
				Tcpsocket{"127.0.0.1", 8000, time.Duration(time.Second * 1)},
				0,
				0,
				[]float64{},
			}
			min, max, avg := p.Statstics()
			So(min, ShouldEqual, 0)
			So(max, ShouldEqual, 0)
			So(avg, ShouldEqual, 0)
		})
		Convey("Ping:Statstics 返回最大最小平均值(连接时间列表不为空)", func() {
			p := Ping{
				Tcpsocket{"127.0.0.1", 8000, time.Duration(time.Second * 1)},
				0,
				0,
				[]float64{1, 2, 3},
			}
			min, max, avg := p.Statstics()
			So(min, ShouldEqual, 1)
			So(max, ShouldEqual, 3)
			So(avg, ShouldEqual, 2)
		})

		Convey("Ping:Multiping 建立无法连接的连接\n", func() {
			p := Ping{
				Tcpsocket{"127.0.0.1", 8001, time.Duration(time.Second * 1)},
				0,
				0,
				[]float64{},
			}
			p.Multiping(5)
			So(p.Successed, ShouldEqual, 0)
			So(p.Failed, ShouldEqual, 5)
			So(len(p.Conn_times), ShouldEqual, 0)
		})
		Convey("Ping:Multiping 建立连接\n", func() {
			p := Ping{
				Tcpsocket{"127.0.0.1", 8000, time.Duration(time.Second * 1)},
				0,
				0,
				[]float64{},
			}
			p.Multiping(5)
			So(p.Successed, ShouldEqual, 5)
			So(p.Failed, ShouldEqual, 0)
			So(len(p.Conn_times), ShouldEqual, 5)
		})
	})
}
