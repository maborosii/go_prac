package importer

import (
	c "oamigrate/config"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"xorm.io/xorm"
)

func TestPreImport(t *testing.T) {
	Convey("preTableHook", t, func() {
		Convey("PreTableHook", func() {
			engine, _ := xorm.NewEngine("mysql", c.ImportConfig.BuildConnectString())
			a := preTableHook(engine)
			So(a, ShouldBeNil)
		})
		// Convey("getDelTable", func() {
		// 	a1 := wantedTables
		// 	a2 := []string{"c_form_formset", "c_form_formset_20210815_bygo", "c_form_formset_20210816_bygo", "c_form_formset_20210817_bygo", "c_form_maininfo", "c_form_maininfo_20210815_bygo", "c_form_maininfo_20210816_bygo", "c_form_maininfo_20210817_bygo", "c_form_tree_def", "c_form_tree_def_20210815_bygo", "c_form_tree_def_20210816_bygo", "c_form_tree_def_20210817_bygo", "wf_action", "wf_action_20210815_bygo", "wf_action_20210816_bygo", "wf_action_20210817_bygo", "wf_flow_def", "wf_flow_def_20210815_bygo", "wf_flow_def_20210816_bygo", "wf_flow_def_20210817_bygo", "wf_right_def", "wf_right_def_20210815_bygo", "wf_right_def_20210816_bygo", "wf_right_def_20210817_bygo", "wf_step_def", "wf_step_def_20210815_bygo", "wf_step_def_20210816_bygo", "wf_step_def_20210817_bygo", "wf_tree_def", "wf_tree_def_20210815_bygo", "wf_tree_def_20210816_bygo", "wf_tree_def_20210817_bygo"}
		// 	a := getDelTable(a1, a2)
		// 	for _, d := range a {
		// 		fmt.Println(*d)
		// 	}
		// 	So(a, ShouldBeNil)
		// },
		// Convey("renameTable", func() {
		// engine, _ := xorm.NewEngine("mysql", c.ImportConfig.BuildConnectString())
		// 	a := renameTable(engine, wantedTables)
		// 	So(a, ShouldBeNil)
		// }),
		// Convey("dropTable", func() {
		// engine, _ := xorm.NewEngine("mysql", c.ImportConfig.BuildConnectString())
		// 	a := dropTables(engine, wantedTables)
		// 	So(a, ShouldBeNil)
		// })
	})
}
