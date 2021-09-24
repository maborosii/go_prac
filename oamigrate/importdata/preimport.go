package importdata

import (
	. "excelfromdb/locallog"
	"regexp"
	"sort"
	"strconv"
	"strings"

	db "oamigrate/dbconfig"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

const suffix = "bygo"

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

func preTableHook(local_config *db.DBConfig) error {
	engine, err := xorm.NewEngine("mysql", local_config.BuildConnectString())
	if err != nil {
		return err
	}
	var tables []*schemas.Table
	var names tableNames
	if tables, err = engine.DBMetas(); err != nil {
		return err
	}
	for _, table := range tables {
		names = append(names, table.Name)
	}
	Log.Info(names)
	return nil
}

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
			// 组内元素个数小于等于1时，直接返回
			if len(tn) <= 1 {
				return tn
			}
			// 组内元素个数大于1时，进行从大到小排序，并剔除最大的元素
			sort.Slice(tn, func(i, j int) bool {
				//这里不用做匹配选择，全量查找时已经将符合规则的表名添加到切片中
				iSerial, _ := strconv.Atoi(pattern.FindAllString(tn[i], -1)[0])
				jSerial, _ := strconv.Atoi(pattern.FindAllString(tn[j], -1)[0])
				return iSerial > jSerial
			})
			return tn[1:]
		},
		)
		delTables = append(delTables, delGroup)
	}
	return delTables
}
