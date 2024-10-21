package dao

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var defaultCtx = context.Background()

var ProvideDao = wire.NewSet(
	New,
	wire.FieldsOf(new(Dao), "Q", "AlarmDao", "DailyRevenueDao", "ParkingDao", "ParkingViewDao", "SummaryDao"),
)

func New(db *gorm.DB) Dao {
	q := Use(db)
	return Dao{
		Q:               q,
		AlarmDao:        q.AlarmPo.WithContext(defaultCtx),
		DailyRevenueDao: q.DailyRevenuePo.WithContext(defaultCtx),
		ParkingDao:      q.ParkingPo.WithContext(defaultCtx),
		ParkingViewDao:  q.ParkingViewPo.WithContext(defaultCtx),
		SummaryDao:      q.SummaryPo.WithContext(defaultCtx),
	}
}

type Dao struct {
	Q               *Query
	AlarmDao        IAlarmPoDo
	DailyRevenueDao IDailyRevenuePoDo
	ParkingDao      IParkingPoDo
	ParkingViewDao  IParkingViewPoDo
	SummaryDao      ISummaryPoDo
}
