module oamigrate

go 1.16

replace excelfromdb => ../excelfromdb

require (
	excelfromdb v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
	github.com/smartystreets/goconvey v1.6.4
	gopkg.in/ini.v1 v1.62.0
	xorm.io/xorm v1.2.3
)
