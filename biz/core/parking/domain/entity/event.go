package parking_entity

import (
	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
)

type CheckedInEvent struct {
	CarPlate    CarPlate
	CheckInTime Time

	domain.EventTag
}

func (event CheckedInEvent) EventID() string {
	return event.CarPlate.String()
}

type CheckInFailedEvent struct {
	CarPlate    CarPlate
	CheckInTime Time
	Message     string

	domain.EventTag
}

func (event CheckInFailedEvent) EventID() string {
	return event.CarPlate.String()
}

type CheckOutFailedEvent struct {
	CarPlate     CarPlate
	CheckOutTime Time
	Message      string

	domain.EventTag
}

func (event CheckOutFailedEvent) EventID() string {
	return event.CarPlate.String()
}

type CheckedOutEvent struct {
	CarPlate     CarPlate
	CheckOutTime Time

	domain.EventTag
}

func (event CheckedOutEvent) EventID() string {
	return event.CarPlate.String()
}

type PaidEvent struct {
	CarPlate CarPlate
	Amount   int64
	PayTime  Time

	domain.EventTag
}

func (event PaidEvent) EventID() string {
	return event.CarPlate.String()
}
