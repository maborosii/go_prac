package pinger

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"time"

	"github.com/maborosii/gcping/sum"
	"github.com/maborosii/gcping/timerecorder"
	"github.com/shopspring/decimal"
)

type Tcpsocket struct {
	Host string
	Port int

	// 这里的参数类型不能为int型，会造成参数类型错误，无法使用net dialtimeout
	Timeout time.Duration
}

type Conn interface {
	Getconn()
	Closeconn()
}

func (socket *Tcpsocket) Getconn() (*net.Conn, error) {
	addr := socket.Host + ":" + strconv.Itoa(socket.Port)
	conn, err := net.DialTimeout("tcp", addr, socket.Timeout)
	return &conn, err
}
func (socket *Tcpsocket) Closeconn(conn *net.Conn) {
	(*conn).Close()
}

type Ping struct {
	Socket     Tcpsocket
	Successed  int
	Failed     int
	Conn_times []float64
}

type Costtime struct {
	rec timerecorder.TimeRecorder
	f   func() error
}

func (ping *Ping) Createsocket() (*net.Conn, error) {
	return ping.Socket.Getconn()
}
func (ping *Ping) Successedrate() (count int, rate_percent string) {
	var rate float64

	count = ping.Successed + ping.Failed
	if count == 0 {
		rate = 0.00
	} else {
		rate = float64(ping.Successed) / float64(count) * 100
	}
	rate_percent = fmt.Sprintf("%.2f%%", rate)
	return
}

func (ping *Ping) Statstics() (mininum float64, maxinum float64, avg float64) {
	if len(ping.Conn_times) == 0 {
		ping.Conn_times = append(ping.Conn_times, 0)
	}
	sort.Float64s(ping.Conn_times)
	mininum = ping.Conn_times[0]
	maxinum = ping.Conn_times[len(ping.Conn_times)-1]
	avg = sum.Avg(ping.Conn_times)
	return
}

func (ping *Ping) singleping() error {

	// package from connection build to connection close
	s, err := ping.Createsocket()
	if err != nil {
		return err
	} else {
		ping.Socket.Closeconn(s)
	}
	return nil
}

func (ping *Ping) Multiping(count int) {
	args := Costtime{
		timerecorder.NewTimeRecorder(),
		ping.singleping,
	}
	time_cost := timerecorder.TimeCostDecorator(args.rec, args.f)
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		err := time_cost()

		cost_time, _ := decimal.NewFromFloat(args.rec.Cost().Seconds() * 1000).Round(2).Float64()

		if err != nil {
			fmt.Printf("Connectting to %s[:%d]: FAILED，error info: %s\n", ping.Socket.Host, ping.Socket.Port, err)
			ping.Failed += 1
		} else {
			fmt.Printf("Connected to %s[:%d]: seq=%d time=%.2f ms\n", ping.Socket.Host, ping.Socket.Port, i+1, cost_time)
			ping.Successed += 1
			ping.Conn_times = append(ping.Conn_times, cost_time)

		}

	}
}
