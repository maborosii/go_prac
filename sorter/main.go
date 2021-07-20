// package main

// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// )

// var infile *string = flag.String("i", "infile", "File contains values for sorting")
// var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
// var algorithm *string = flag.String("a", "qsort", "sort algorithm")

// func readValues(infile string) (values []int, err error) {
// 	file, err := os.Open(infile)
// 	if err != nil {
// 		fmt.Println("Failed to open the input file", infile)
// 		return
// 	}
// 	defer file.Close()

// 	br := bufio.NewReader(file)

// 	// values := make([]int, 0)
// 	values := make([]int, 0)

// 	for {
// 		line, isPrefix, errline := br.ReadLine()
// 		if errline != nil {
// 			if errline != io.EOF {
// 				err = errline
// 			}
// 			break
// 		}

// 		if isPrefix {
// 			fmt.Println("A too long line, seems unexpected.")
// 			return
// 		}

// 		str := string(line)
// 		value, errline := strconv.Atoi(str)

// 		if errline != nil {
// 			err = errline
// 			return
// 		}

// 		values = append(values, value)

// 	}
// 	return
// }

// func main() {
// 	flag.Parse()

// 	if infile != nil {
// 		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
// 	}

// 	values, err := readValues(*infile)
// 	if err != nil {
// 		fmt.Println(err)

// 	} else {
// 		fmt.Println("Read values:", values)
// 	}

// }
package main

import "fmt"

type Interger int

func (a *Interger) Add(b Interger) {
	*a += b
}
func main() {
	var a Interger = 1
	a.Add(2)
	fmt.Println(a)
}
