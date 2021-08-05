package main

import (
	"fmt"
	"reflect"
	"time"
)

type MsgStru struct {
	msg []int
}

var msgChan chan MsgStru

func sendMsg() {
	buf := make([]int, 10)
	for i := 0; i < 10; i++ {
		buf[0] = i
		tmp := MsgStru{msg: buf[:1]}
		fmt.Printf("send:%+v\n", tmp)
		fmt.Println(reflect.TypeOf(tmp))
		msgChan <- tmp
	}
	fmt.Println("finish send")
	close(msgChan)
}

func recvMsg() {
	for {
		tmp, ok := <-msgChan
		fmt.Printf("recv:%+v\n", tmp)
		// fmt.Printf("receive address in mem:%p\n", &tmp)
		if !ok {
			break
		}
	}
}

func main() {

	msgChan = make(chan MsgStru)
	go sendMsg()
	time.Sleep(1 * time.Second)
	recvMsg()

}
