package main

import (
	"fmt"
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

func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- 1

}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	fmt.Println(len(chs))
	for _, ch := range chs {
		fmt.Println("*")
		<-ch
	}
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
