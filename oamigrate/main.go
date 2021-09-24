package main

import (
	"fmt"
	// . "oamigrate/importdata"
	"regexp"
)

func main() {

	// err := mysqldump.FullExport()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if err := dbconfig.ImportTable("backup.sql"); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := ImportTable(); err != nil {
	// 	fmt.Println(err)
	// }
	demo := "cc_sdfsdfsdf_20212093_bygo"
	pattern := regexp.MustCompile(`20\d+`)
	fmt.Println(pattern.FindAllString(demo, -1))

}
