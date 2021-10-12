package importer

import (
	"errors"
	"fmt"
	. "oamigrate/pkg/log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	c "oamigrate/config"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

const suffix = "bygo"

var wantedTables = c.Tables

// 导入表的前置操作
func preTableHook(engine *xorm.Engine) error {
	var tables []*schemas.Table
	var allTableNames tableNames
	tables, err := engine.DBMetas()
	if err != nil {
		Log.Error("get all tables information occur error")
		return err
	}

	err = renameTable(engine, wantedTables)
	if err != nil {
		return err
	}

	for _, table := range tables {
		allTableNames = append(allTableNames, table.Name)
	}
	delTableGroup := getDelTable(wantedTables, allTableNames)

	spreadTables := func(tg []*tableGroup) tableNames {
		delTables := tableNames{}
		for _, g := range tg {
			delTables = append(delTables, g.names...)
		}
		return delTables
	}(delTableGroup)

	if err = dropTables(engine, spreadTables); err != nil {
		return err
	}

	return nil
}

// 删除的表分组
type tableGroup struct {
	prefix string
	names  tableNames
}

type tableNames []string

func NewTableGroup(prefix string) *tableGroup {
	return &tableGroup{
		prefix: prefix,
		names:  []string{},
	}
}

func (tg *tableGroup) addTable(table string) {
	tg.names = append(tg.names, table)
}
func (tg *tableGroup) cleanTable(f func(tableNames) tableNames) {
	tg.names = f(tg.names)
}

// 获取删除表的集合
func getDelTable(wanted tableNames, all tableNames) []*tableGroup {
	var delTables []*tableGroup
	pattern := regexp.MustCompile(`20\d+`)

	for _, prefix := range wanted {
		delGroup := NewTableGroup(prefix)
		for _, table := range all {
			if strings.HasPrefix(table, prefix) && strings.HasSuffix(table, suffix) {
				delGroup.addTable(table)
			}
		}

		delGroup.cleanTable(func(tn tableNames) tableNames {
			// 组内元素个数小于等于1时，直接空切片
			if len(tn) <= 2 {
				emptySlice := make([]string, 0)
				return emptySlice
			}
			// 组内元素个数大于1时，进行从大到小排序，并剔除最大的元素
			sort.Slice(tn, func(i, j int) bool {
				//这里不用做匹配选择，全量查找时已经将符合规则的表名添加到切片中
				iSerial, _ := strconv.Atoi(pattern.FindAllString(tn[i], -1)[0])
				jSerial, _ := strconv.Atoi(pattern.FindAllString(tn[j], -1)[0])
				return iSerial > jSerial
			})
			return tn[2:]
		},
		)
		delTables = append(delTables, delGroup)
	}
	return delTables
}

func dropTables(engine *xorm.Engine, tables tableNames) error {
	// []string类型的实参 无法作为参数传入以[]interface{}类型为形参的函数,需要做转换

	// 转换[]string 为[]interface{}
	tableInterfaceList := make([]interface{}, len(tables))
	for i, t := range tables {
		//string convert to interface{}
		tableInterfaceList[i] = t
	}

	// 	这个方法本身就做了事务处理
	err := engine.DropTables(tableInterfaceList...)
	if err != nil {
		Log.Error("drop backup tables occur error")
		return err
	}
	Log.Info("drop backup tables successs")
	return nil
}

//重命名主表
func renameTable(engine *xorm.Engine, tables tableNames) error {
	for _, t := range tables {
		//方法本身做了事务处理
		isExist, _ := engine.IsTableExist(t)
		if !isExist {
			Log.Error(t, " is not exist! please check it")
			return errors.New("table is not exists")
		}
	}

	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		Log.Error("starting reneme tables transaction occur error")
		return err
	}

	for _, t := range tables {
		// todo: 当次重复导入
		newTableName := t + "_" + time.Now().Format("20060102") + "_" + suffix
		renameSql := fmt.Sprintf("RENAME TABLE %s TO %s;", t, newTableName)

		backupTableIsExist, _ := session.IsTableExist(newTableName)

		if !backupTableIsExist {
			_, err = session.Exec(renameSql)
			if err != nil {
				Log.Error("rename table ", t, " occur error")
				session.Rollback()
				return err
			}
			Log.Info("rename table ", t, " success")
		}
	}
	Log.Info("rename all table transaction committed")
	return session.Commit()
}
