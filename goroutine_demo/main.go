package main

import (
	"fmt"
	"time"
)

// import (
// 	"fmt"
// 	"time"
// )

// func Add(x, y int) {
// 	z := x + y
// 	fmt.Println(z)
// }
// func main() {
// 	go Add(1, 1)
// 	fmt.Println("***")
// 	time.Sleep(time.Second)
// }

// func Count(ch chan int) {
// 	fmt.Println("Counting")
// 	ch <- 1

// }

var strChan = make(chan string, 3)

func main() {
	// chs := make([]chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	chs[i] = make(chan int)
	// 	go Count(chs[i])
	// }
	// fmt.Println(len(chs))
	// for _, ch := range chs {
	// 	fmt.Println("*")
	// 	<-ch
	// }
	// name := []string{"aa", "cc", "dd"}
	// for _, n := range name {

	// 	go func(who string) {
	// 		fmt.Println("呵呵", who)
	// 	}(n)
	// }
	syncChan1 := make(chan struct{}, 2) // 存储空结构体通道
	syncChan2 := make(chan struct{}, 2)
	go func() { //reveiver
		<-syncChan1
		fmt.Println("waiting a second... [reveiver]")
		time.Sleep(1 * time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() { //sender
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Send a sync signal. [sender]")
			}

		}
		fmt.Println("waiting two seconds... [sender]")
		time.Sleep(2 * time.Second)

		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2

	// // name = "bbb"
	// time.Sleep(100 * time.Millisecond)

}

// func main() {
// 	ch := make(chan int, 2)
// 	go func() {
// 		fmt.Println("SS:  sub goroutine started")
// 		for i := 1; i < 11; i++ {
// 			fmt.Println("SS:  enter the loop number : ", i)
// 			ch <- i
// 			fmt.Println("SS:  write ! --- current channel length is", len(ch))
// 		}

// 		// fmt.Println("sub goroutine is done")
// 	}()
// 	fmt.Println("MM:  main goroutine started")
// 	fmt.Println("MM:  before the channel writed, current channel length is", len(ch))
// 	data := <-ch
// 	fmt.Println("MM:  read ! --- current channel length is", len(ch))

// 	data1 := <-ch
// 	fmt.Println("MM:  read ! --- current channel length is", len(ch))

// 	data2 := <-ch
// 	fmt.Println("MM:  read ! --- current channel length is", len(ch))

// 	data3 := <-ch
// 	fmt.Println("MM:  read ! --- current channel length is", len(ch))

// 	fmt.Println(data)
// 	fmt.Println(data1)
// 	fmt.Println(data2)
// 	fmt.Println(data3)
// 	fmt.Println("MM:  channel done")

// 	// select {}

// }
