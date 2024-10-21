package app

import (
	"context"

	"code.byted.org/oec/status_code/Status"
	E "github.com/dirac-lee/ddd_demo_parking/biz/common/exception"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/logx"
	"github.com/dirac-lee/ddd_demo_parking/biz/model/ddd_demo_parking"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/dao"
)

type ParkingQuery struct {
	ParkingViewDao  dao.IParkingViewPoDo
	SummaryDao      dao.ISummaryPoDo
	DailyRevenueDao dao.IDailyRevenuePoDo
	DtoAssembler    DtoAssembler
}

func (q ParkingQuery) GetHistory(ctx context.Context, req *ddd_demo_parking.GetHistoryRequest) (resp *ddd_demo_parking.GetHistoryResponse) {
	logx.InfoInput(ctx, map[string]any{"req": req})
	defer func() { logx.InfoOutput(ctx, map[string]any{"resp": resp}) }()
	var (
		ex *E.Exception
	)
	pos, err := q.ParkingViewDao.WithContext(ctx).Find()
	if err != nil {
		ex = E.New(Status.DBQueryException, E.WithCause(err))
	}
	return q.DtoAssembler.ToGetHistoryResponse(pos, ex)
}

func (q ParkingQuery) GetTotalInPark(ctx context.Context, req *ddd_demo_parking.GetTotalInParkRequest) (resp *ddd_demo_parking.GetTotalInParkResponse) {
	logx.InfoInput(ctx, map[string]any{"req": req})
	defer func() { logx.InfoOutput(ctx, map[string]any{"resp": resp}) }()
	var (
		ex *E.Exception
	)
	po, err := q.SummaryDao.WithContext(ctx).FirstOrInit()
	if err != nil {
		ex = E.New(Status.DBQueryException, E.WithCause(err))
	}
	return q.DtoAssembler.ToGetTotalInParkResponse(po, ex)
}

func (q ParkingQuery) GetDailyRevenue(ctx context.Context, req *ddd_demo_parking.GetDailyRevenueRequest) (resp *ddd_demo_parking.GetDailyRevenueResponse) {
	logx.InfoInput(ctx, map[string]any{"req": req})
	defer func() { logx.InfoOutput(ctx, map[string]any{"resp": resp}) }()
	var (
		ex *E.Exception
	)
	pos, err := q.DailyRevenueDao.WithContext(ctx).Find()
	if err != nil {
		ex = E.New(Status.DBQueryException, E.WithCause(err))
	}

	return q.DtoAssembler.ToGetDailyRevenueResponse(pos, ex)
}
