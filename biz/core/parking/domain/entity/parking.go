package parking_entity

import (
	"context"
	"time"

	"code.byted.org/oec/status_code/Status"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
	E "github.com/dirac-lee/ddd_demo_parking/biz/common/exception"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/logx"
)

const FEE_PER_HOUR = 1

type Parking interface {
	GetId() CarPlate
	GetCheckInTime() *Time
	GetLastPayTime() *Time
	GetTotalPaid() int64

	HandleCheckInCommand(ctx context.Context, publisher domain.EventPublisher, command CheckInCommand) bool
	HandleNotifyPayCommand(ctx context.Context, publisher domain.EventPublisher, command NotifyPaidCommand)
	HandleCheckOutCommand(ctx context.Context, publisher domain.EventPublisher, command CheckOutCommand) bool
	CalculateFee(ctx context.Context, time Time) int64
}

type ParkingImpl struct {
	Id          CarPlate
	CheckInTime *Time
	LastPayTime *Time
	TotalPaid   int64
}

func (p *ParkingImpl) GetId() CarPlate {
	return p.Id
}

func (p *ParkingImpl) GetCheckInTime() *Time {
	return p.CheckInTime
}

func (p *ParkingImpl) GetLastPayTime() *Time {
	return p.LastPayTime
}

func (p *ParkingImpl) GetTotalPaid() int64 {
	return p.TotalPaid
}

func (p *ParkingImpl) HandleCheckInCommand(ctx context.Context, publisher domain.EventPublisher, command CheckInCommand) (success bool) {
	logx.InfoInput(ctx, map[string]any{"command": command})
	defer func() { logx.InfoOutput(ctx, map[string]any{"success": success}) }()
	if p.inPark() {
		logx.Info(ctx, "car is in park, can not check in")
		publisher.PublishEvent(ctx, CheckInFailedEvent{CarPlate: p.Id, CheckInTime: command.CheckInTime, Message: "the car is already in park, cannot check in"})
		return false
	}

	publisher.PublishEvent(ctx, CheckedInEvent{CarPlate: p.Id, CheckInTime: command.CheckInTime})
	p.CheckInTime = &command.CheckInTime
	return true
}

func (p *ParkingImpl) HandleCheckOutCommand(ctx context.Context, publisher domain.EventPublisher, command CheckOutCommand) (success bool) {
	logx.InfoInput(ctx, map[string]any{"command": command})
	defer func() { logx.InfoOutput(ctx, map[string]any{"success": success}) }()
	if !p.inPark() {
		logx.Info(ctx, "car is not in park, can not check out")
		publisher.PublishEvent(ctx, CheckOutFailedEvent{CarPlate: p.Id, CheckOutTime: command.CheckOutTime, Message: "car is not in park"})
		return false
	}

	fee := p.CalculateFee(ctx, command.CheckOutTime)
	if fee > 0 {
		logx.Info(ctx, "fee(%v) is not clear, can not check out", fee)
		return false
	}

	p.CheckInTime = nil
	p.TotalPaid = 0
	p.LastPayTime = nil

	publisher.PublishEvent(ctx, CheckedOutEvent{CarPlate: p.Id, CheckOutTime: command.CheckOutTime})
	return true
}

func (p *ParkingImpl) HandleNotifyPayCommand(ctx context.Context, publisher domain.EventPublisher, command NotifyPaidCommand) {
	logx.InfoInput(ctx, map[string]any{"command": command})
	if !p.inPark() {
		E.Throw(E.New(Status.InvalidParam, E.WithDetail("the car is not in park, can not pay")))
	}
	p.LastPayTime = &command.PayTime
	p.TotalPaid += command.Amount

	publisher.PublishEvent(ctx, PaidEvent{CarPlate: p.Id, PayTime: command.PayTime, Amount: command.Amount})
}

func (p *ParkingImpl) CalculateFee(ctx context.Context, now Time) (fee int64) {
	defer func() { logx.InfoOutput(ctx, map[string]any{"fee": fee}) }()

	if !p.inPark() {
		E.Throw(E.New(Status.InvalidParam, E.WithDetail("the car is not in park")))
	}
	currentCheckInTime := *p.CheckInTime
	if p.LastPayTime == nil {
		return p.feeBetween(currentCheckInTime, now)
	}
	lastPayTime0 := *p.LastPayTime
	if p.TotalPaid < p.feeBetween(currentCheckInTime, lastPayTime0) {
		return p.feeBetween(currentCheckInTime, NewTime(time.Now())) - p.TotalPaid
	}
	if lastPayTime0.Add(15 * time.Minute).After(now.Into()) {
		return 0
	}

	return p.feeBetween(currentCheckInTime, now) - p.TotalPaid
}

func (p *ParkingImpl) feeBetween(start, end Time) (fee int64) {
	return p.hoursBetween(start, end) * FEE_PER_HOUR
}

func (p *ParkingImpl) hoursBetween(start, end Time) (hours int64) {
	minutes := int64(end.Sub(start.Into()).Minutes())
	hours = minutes / 60
	if hours*60 == minutes {
		return hours
	} else {
		return hours + 1
	}
}

func (p *ParkingImpl) inPark() bool {
	return p.CheckInTime != nil
}
