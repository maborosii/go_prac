package importer

import (
	"fmt"

	c "oamigrate/config"
	. "oamigrate/pkg/log"

	"xorm.io/xorm"
)

//修改各地市环境passid
func postTableHook(engine *xorm.Engine) error {
	currnetPassId := func() string {
		if _, ok := c.CityPassId[c.Node]; ok {
			return c.CityPassId[c.Node]
		}
		return c.CityPassId["default"]

	}()
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		Log.Error("starting update passid transaction occur error")
		return err
	}

	updateSql := fmt.Sprintf("UPDATE c_form_maininfo SET FORM_STRU = REPLACE(FORM_STRU, 'gdxzzfpc', '%s') where FORM_STRU like '%%gdxzzfpc%%';", currnetPassId)

	// engine.ShowSQL(true)
	_, err = session.Exec(updateSql)
	if err != nil {
		Log.Error("upate passid occur error")
		session.Rollback()
		return err
	}
	Log.Info("update passid success")
	return session.Commit()
}
