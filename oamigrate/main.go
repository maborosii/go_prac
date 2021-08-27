package main

import (
	// "xorm.io/xorm/schemas"
	// . "oamigrate/models"
	"fmt"
	"oamigrate/mysqldump"
)

// var tables = [...]*schemas.Table{&CFormFormset{},}

func main() {

	err, paths := mysqldump.FullExport()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(paths)
	// if err := dbconfig.DumpTable("a.sql"); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := dbconfig.ImportTable("backup.sql"); err != nil {
	// 	fmt.Println(err)
	// }

}
