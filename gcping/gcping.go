package main

import (
	"fmt"
	"time"

	"github.com/maborosii/gcping/formatter"
	"github.com/maborosii/gcping/pinger"
)

func main() {
	demo_ping := pinger.Ping{
		Socket: pinger.Tcpsocket{
			Host:    "127.0.0.1",
			Port:    8001,
			Timeout: time.Duration(1 * time.Second),
		},
		Successed:  0,
		Failed:     0,
		Conn_times: make([]float64, 0),
	}

	demo_ping.Multiping(5)

	ping_formatter := formatter.NewFormatter()
	ping_formatter = ping_formatter.Addstat(demo_ping)
	fmt.Println(ping_formatter.Setraw())
	ping_formatter.Settb()
}
