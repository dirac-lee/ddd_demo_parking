package parking_policy

import (
	"context"

	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	parking_infra "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/infra"
)

type AlarmPolicy struct {
	AlarmService parking_infra.AlarmService
}

func (p AlarmPolicy) OnEvent(ctx context.Context, event domain.Event) {
	switch e := event.(type) {
	case parking_entity.CheckOutFailedEvent:
		p.AlarmService.Alarm(ctx, e.CarPlate, e.Message, e.CheckOutTime)
	case parking_entity.CheckInFailedEvent:
		p.AlarmService.Alarm(ctx, e.CarPlate, e.Message, e.CheckInTime)
	}
}
