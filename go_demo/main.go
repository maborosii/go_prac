package main

import "fmt"

// func get_name() (string, string, string) {
// 	return "a", "b", "c"
// }
// func main() {
// fmt.Println("hello world!")
// _, _, nickname := get_name()
// fmt.Println(nickname)
// 	var arr = [4]int{1, 2, 3, 4}
// 	slice := arr[:3]
// 	slice2 := append(slice, 1, 3)
// 	for value := range slice {
// 		fmt.Printf("**%d**\n", value)
// 		fmt.Println("**", value, "**")
// 	}
// 	fmt.Println(slice2)
// 	fmt.Println(arr)
// }

/*
! 函数的不定长参数
*/
// func MyPrintf(args ...interface{}) {
// 	for i, arg := range args {
// 		fmt.Println(i, arg)
// 		switch arg.(type) {
// 		case int:
// 			fmt.Println(arg, "is an int value")
// 		case string:
// 			fmt.Println(arg, "is an string value")
// 		case int64:
// 			fmt.Println(arg, "is an int64 value")
// 		default:
// fmt.Println(arg, "is an unknown type")

// 		}
// 	}
// }
// func main() {
// 	var v1 int = 1
// 	var v2 int64 = 2
// 	var v3 string = "hello"
// 	var v4 float32 = 1.234

// 	MyPrintf(v1, v2, v3, v4)
// }

/*
! defer与return的优先级，函数的命名返回值和匿名返回值
*/
// func x(i *int) int {
// 	// * 匿名函数花括号直接跟参数列表表示函数调用
// 	defer func() {
// 		*i = 19
// 	}()
// 	return *i
// }

// // ? 命名返回值，就好比函数参数一样，函数体内对命名返回值的任何修改，都会影响它。而非命名返回值即匿名返回值，
// // ? 取决于 return 时候的值。
// func main() {
// 	i := 10
// 	j := x(&i)
// 	fmt.Println(i, j)
// }

/*
! 闭包
*/
// func main() {
// 	j := 5
// 	a := func() func() {
// 		i := 10
// 		return func() {
// 			fmt.Println("i, j:", i, j)
// 		}
// 	}()
// 	a()
// 	j *= 2
// 	a()
// }

/*
! 给新类型添加方法
! 接口的非侵入式实现
*/
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAddr interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	// var a Integer = 1
	// if a.Less2(2) {
	// 	fmt.Println("success")
	// } else {
	// 	fmt.Println("fail")
	// }

	var a Integer = 1
	var b LessAddr = &a
	fmt.Println(b.Less(3))

}

/*
! 匿名组合
*/

// type Job struct {
// 	Command string
// 	*log.Logger
// }

// func (job *Job) Start() {
// 	job.Log("staring now...")

// 	//...

// 	job.Log("started")

// }
