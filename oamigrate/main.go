package main

import (
	// "xorm.io/xorm/schemas"
	// . "oamigrate/models"
	"fmt"
	. "oamigrate/importdata"
)

func main() {

	// err := mysqldump.FullExport()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if err := dbconfig.ImportTable("backup.sql"); err != nil {
	// 	fmt.Println(err)
	// }
	if err := ImportTable(); err != nil {
		fmt.Println(err)
	}

}
