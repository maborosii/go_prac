package main

import "embed"

type path string

var (
	//go:embed dbconfig/db.conf
	dbconf embed.FS

	//go:embed sqlinfo/casestatistics_detail.sql
	casedetail_sql string

	//go:embed sqlinfo/casestatistics_allnum.sql
	caseallnum_sql string

	//go:embed sqlinfo/casestatistics_deltanum.sql
	casedeltanum_sql string

	//go:embed sqlinfo/publicity/publicity_allnum.sql
	publicityallnum_sql string

	//go:embed sqlinfo/publicity/publicity_deltanum.sql
	publicitydeltanum_sql string

	//go:embed sqlinfo/publicity/publicity_allnum_dept.sql
	publicityallnum_dept_sql string

	//go:embed sqlinfo/publicity/publicity_allnum_user.sql
	publicityallnum_user_sql string

	//go:embed sqlinfo/publicity/six_allnum.sql
	sixallnum_sql string

	//go:embed sqlinfo/publicity/six_allnum_pub.sql
	sixallnum_pub_sql string

	//go:embed sqlinfo/publicity/six_deltanum.sql
	sixdeltanum_sql string

	//go:embed sqlinfo/publicity/six_deltanum_pub.sql
	sixdeltanum_pub_sql string

	//go:embed excelops/templates/publicty_template.xlsx
	publicityTemplate embed.FS
)

var (
	publictiy_template path   = "excelops/templates/publicty_template.xlsx"
	six_sheet          string = "02、公示六大类统计"
	publicity_sheet    string = "03、公示采集统计"
	acceessstat_sheet  string = "04、访问量统计"
)
