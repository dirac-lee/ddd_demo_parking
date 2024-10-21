package parking_infra_impl

import (
	"context"
	"time"

	"code.byted.org/lang/gg/gptr"
	"code.byted.org/oec/status_code/Status"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
	E "github.com/dirac-lee/ddd_demo_parking/biz/common/exception"
	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	parking_policy "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/policy"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/dao"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/model"
	"github.com/google/wire"
)

var ProvideListeners = wire.NewSet(
	wire.Struct(new(SummaryListener), "*"),
	wire.Struct(new(ParkingViewListener), "*"),
	wire.Struct(new(DailyRevenueListener), "*"),
	wire.Struct(new(parking_policy.AlarmPolicy), "*"),
)

type SummaryListener struct {
	Dao   *dao.Query
	IDGen domain.IDGen
}

func (l SummaryListener) OnEvent(ctx context.Context, event domain.Event) {
	switch event.(type) {
	case parking_entity.CheckedInEvent:
		l.increaseCar(ctx, 1)
	case parking_entity.CheckedOutEvent:
		l.increaseCar(ctx, -1)
	}
}

const parkID = 7384102776687670387

func (l SummaryListener) increaseCar(ctx context.Context, increment int64) {
	summaryDao := l.Dao.SummaryPo
	po, err := summaryDao.WithContext(ctx).Where(l.Dao.SummaryPo.ID.Eq(parkID)).FirstOrInit()
	if err != nil {
		E.Throw(E.New(Status.DBQueryException, E.WithCause(err)))
	}
	po.TotalInParking += increment
	err = summaryDao.WithContext(ctx).Save(po)
	if err != nil {
		E.Throw(E.New(Status.DBSaveException, E.WithCause(err)))
	}
}

type ParkingViewListener struct {
	Dao   *dao.Query
	IDGen domain.IDGen
}

func (l ParkingViewListener) OnEvent(ctx context.Context, event domain.Event) {
	switch e := event.(type) {
	case parking_entity.CheckedInEvent:
		l.insertHistory(ctx, e)
	case parking_entity.CheckedOutEvent:
		l.updateOnCheckOut(ctx, e)
	case parking_entity.PaidEvent:
		l.updateOnPaid(ctx, e)
	}
}

func (l ParkingViewListener) insertHistory(ctx context.Context, event parking_entity.CheckedInEvent) {
	viewDao := l.Dao.ParkingViewPo
	if err := viewDao.WithContext(ctx).Save(&model.ParkingViewPo{
		ID:           l.IDGen.GetID(ctx),
		Plate:        event.CarPlate.String(),
		CheckInTime:  event.CheckInTime.Into(),
		CheckOutTime: nil,
		PayAmount:    0,
	}); err != nil {
		E.Throw(E.New(Status.DBSaveException, E.WithCause(err)))
	}
}

func (l ParkingViewListener) updateOnCheckOut(ctx context.Context, event parking_entity.CheckedOutEvent) {
	viewDao := l.Dao.ParkingViewPo
	po, err := viewDao.WithContext(ctx).Where(viewDao.Plate.Eq(event.CarPlate.String())).Order(viewDao.ID.Desc()).FirstOrInit()
	if err != nil {
		E.Throw(E.New(Status.DBQueryException, E.WithCause(err)))
	}
	po.CheckOutTime = gptr.Of(event.CheckOutTime.Into())
	if err = viewDao.WithContext(ctx).Save(po); err != nil {
		E.Throw(E.New(Status.DBSaveException, E.WithCause(err)))
	}
}

func (l ParkingViewListener) updateOnPaid(ctx context.Context, event parking_entity.PaidEvent) {
	viewDao := l.Dao.ParkingViewPo
	po, err := viewDao.WithContext(ctx).Where(viewDao.Plate.Eq(event.CarPlate.String())).Order(viewDao.ID.Desc()).FirstOrInit()
	if err != nil {
		E.Throw(E.New(Status.DBQueryException, E.WithCause(err)))
	}
	po.PayAmount += event.Amount
	if err = viewDao.WithContext(ctx).Save(po); err != nil {
		E.Throw(E.New(Status.DBSaveException, E.WithCause(err)))
	}
}

type DailyRevenueListener struct {
	Dao   *dao.Query
	IDGen domain.IDGen
}

func (l DailyRevenueListener) OnEvent(ctx context.Context, event domain.Event) {
	switch e := event.(type) {
	case parking_entity.PaidEvent:
		date := time.Now().Format(time.DateOnly)
		l.increaseRevenueByDate(ctx, date, e.Amount)
	}
}

func (l DailyRevenueListener) increaseRevenueByDate(ctx context.Context, date string, amount int64) {
	dailyRevenueDao := l.Dao.DailyRevenuePo
	po, err := dailyRevenueDao.WithContext(ctx).Where(dailyRevenueDao.Date.Eq(date)).FirstOrInit()
	if err != nil {
		E.Throw(E.New(Status.DBQueryException, E.WithCause(err)))
	}
	po.Revenue += amount
	err = dailyRevenueDao.WithContext(ctx).Save(po)
	if err != nil {
		E.Throw(E.New(Status.DBSaveException, E.WithCause(err)))
	}
}
