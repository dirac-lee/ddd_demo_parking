package app

import (
	"context"

	"code.byted.org/gopkg/logs"
	E "github.com/dirac-lee/ddd_demo_parking/biz/common/exception"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/logx"
	parking_domain "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain"
	"github.com/dirac-lee/ddd_demo_parking/biz/model/ddd_demo_parking"
	"github.com/google/wire"
)

var ProvideApp = wire.NewSet(
	wire.Struct(new(DtoAssembler), "*"),
	wire.Struct(new(ParkingQuery), "*"),
	wire.Struct(new(ParkingApp), "*"),
)

type ParkingApp struct {
	DtoAssembler               DtoAssembler
	CheckInCommandHandler      parking_domain.CheckInCommandHandler
	CheckOutCommandHandler     parking_domain.CheckOutCommandHandler
	CalculateFeeCommandHandler parking_domain.CalculateFeeCommandHandler
	NotifyPaidCommandHandler   parking_domain.NotifyPaidCommandHandler

	ParkingQuery
}

func (app ParkingApp) CheckIn(ctx context.Context, req *ddd_demo_parking.CheckInRequest) (resp *ddd_demo_parking.CheckInResponse) {
	logx.InfoInput(ctx, map[string]any{"req": req})
	defer logx.InfoOutput(ctx, map[string]any{"resp": resp})
	var (
		success bool
		ex      *E.Exception
	)
	E.Try(func() {
		success = app.CheckInCommandHandler.Handle(ctx, app.DtoAssembler.ToCheckInCommand(req))
	}).Catch(func(e *E.Exception) {
		logs.CtxError(ctx, "[CheckIn] check in failed, err: %s", e)
		ex = e
	})
	return app.DtoAssembler.ToCheckInResponse(success, ex)
}

func (app ParkingApp) CheckOut(ctx context.Context, req *ddd_demo_parking.CheckOutRequest) (resp *ddd_demo_parking.CheckOutResponse) {
	logx.InfoInput(ctx, map[string]any{"req": req})
	defer logx.InfoOutput(ctx, map[string]any{"resp": resp})
	var (
		success bool
		ex      *E.Exception
	)
	E.Try(func() {
		success = app.CheckOutCommandHandler.Handle(ctx, app.DtoAssembler.ToCheckOutCommand(req))
	}).Catch(func(e *E.Exception) {
		logs.CtxError(ctx, "[CheckOut] check out failed, err: %s", e)
		ex = e
	})
	return app.DtoAssembler.ToCheckOutResponse(success, ex)
}

func (app ParkingApp) CalculateFee(ctx context.Context, req *ddd_demo_parking.CalculateFeeRequest) (resp *ddd_demo_parking.CalculateFeeResponse) {
	logx.InfoInput(ctx, map[string]any{"req": req})
	defer logx.InfoOutput(ctx, map[string]any{"resp": resp})
	var (
		fee int64
		ex  *E.Exception
	)
	E.Try(func() {
		fee = app.CalculateFeeCommandHandler.HandleCalculateFeeCommand(ctx, app.DtoAssembler.ToCalculateFeeCommand(req))
	}).Catch(func(e *E.Exception) {
		logs.CtxError(ctx, "[CheckOut] check out failed, err: %s", e)
		ex = e
	})
	return app.DtoAssembler.ToCalculateFeeResponse(fee, ex)
}

func (app ParkingApp) NotifyPaid(ctx context.Context, req *ddd_demo_parking.NotifyPaidRequest) (resp *ddd_demo_parking.NotifyPaidResponse) {
	logx.InfoInput(ctx, map[string]any{"req": req})
	defer logx.InfoOutput(ctx, map[string]any{"resp": resp})
	var (
		success bool
		ex      *E.Exception
	)
	E.Try(func() {
		success = app.NotifyPaidCommandHandler.HandleNotifyPayCommand(ctx, app.DtoAssembler.ToNotifyPayCommand(req))
	}).Catch(func(e *E.Exception) {
		logs.CtxError(ctx, "[CheckOut] check out failed, err: %s", e)
		ex = e
	})
	return app.DtoAssembler.ToNotifyPaidResponse(success, ex)
}
