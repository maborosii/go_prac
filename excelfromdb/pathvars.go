package main

import "embed"

var (
	//go:embed dbconfig/db.conf
	dbconf embed.FS

	//go:embed sqlinfo/casestatistics_detail.sql
	casedetail_sql string

	//go:embed sqlinfo/casestatistics_allnum.sql
	caseallnum_sql string

	//go:embed sqlinfo/casestatistics_deltanum.sql
	casedeltanum_sql string
)
