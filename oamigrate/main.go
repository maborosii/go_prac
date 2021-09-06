package main

import (
	// "xorm.io/xorm/schemas"
	// . "oamigrate/models"

	"fmt"
	"oamigrate/mysqldump"
)

func main() {

	err := mysqldump.FullExport()
	if err != nil {
		fmt.Println(err)
	}
	// if err := dbconfig.DumpTable("a.sql"); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := dbconfig.ImportTable("backup.sql"); err != nil {
	// 	fmt.Println(err)
	// }

}
