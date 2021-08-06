package main

import (
	"embed"
	"fmt"
)

//go:embed test.yaml
var cityfile embed.FS

func main() {
	filefromyml, err := cityfile.ReadFile("test.yaml")
	// io.Copy(os.Stdout, filefromyml)

	if err != nil {
		panic(err)
	}
	fmt.Println(filefromyml)
	// b := []byte{}

	// _, err = filefromyml.Read(b)
	// // fmt.Println(n)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(b)

}
