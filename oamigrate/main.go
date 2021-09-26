package main

import (
	"fmt"
	"oamigrate/importdata"
)

func main() {
	// err := mysqldump.FullExport()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if err := dbconfig.ImportTable("backup.sql"); err != nil {
	// 	fmt.Println(err)
	// }
	if err := importdata.ImportTable(); err != nil {
		fmt.Println(err)
	}
}
