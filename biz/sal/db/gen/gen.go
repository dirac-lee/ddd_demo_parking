package main

import (
	"code.byted.org/gopkg/env"
	"code.byted.org/gorm/bytedgen"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db"
	"gorm.io/gen"
)

const (
	OutPath = "biz/sal/db/dao"
)

func main() {
	env.SetPSM("oec.governance.guard_center")

	g := bytedgen.NewGenerator(gen.Config{
		OutPath:       OutPath,
		Mode:          gen.WithQueryInterface,
		FieldNullable: true,
	})
	g.UseDB(db.NewDB())

	g.ApplyBasic(
		g.GenerateModelAs(
			"parking",
			"ParkingPo",
		),
		g.GenerateModelAs(
			"alarm",
			"AlarmPo",
		),
		g.GenerateModelAs(
			"daily_revenue",
			"DailyRevenuePo",
		),
		g.GenerateModelAs(
			"parking_view",
			"ParkingViewPo",
		),
		g.GenerateModelAs(
			"summary",
			"SummaryPo",
		),
	)

	g.Execute()
}
