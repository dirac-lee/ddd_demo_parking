package parking_infra_impl

import (
	"code.byted.org/lang/gg/gptr"
	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/model"
)

type ParkingBuilder struct {
}

func (b ParkingBuilder) ToEntity(po *model.ParkingPo) parking_entity.Parking {
	if po == nil {
		return nil
	}
	return &parking_entity.ParkingImpl{
		Id:          parking_entity.CarPlate(po.ID),
		CheckInTime: gptr.Map(po.CheckInTime, parking_entity.IntoTime),
		LastPayTime: gptr.Map(po.LastPayTime, parking_entity.IntoTime),
		TotalPaid:   po.TotalPaid,
	}
}

func (b ParkingBuilder) FromEntity(parking parking_entity.Parking) *model.ParkingPo {
	return &model.ParkingPo{
		ID:          parking.GetId().String(),
		CheckInTime: gptr.Map(parking.GetCheckInTime(), parking_entity.Time.Into),
		LastPayTime: gptr.Map(parking.GetLastPayTime(), parking_entity.Time.Into),
		TotalPaid:   parking.GetTotalPaid(),
	}
}
