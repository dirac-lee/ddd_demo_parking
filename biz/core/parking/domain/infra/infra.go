package parking_infra

import (
	"context"

	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
)

type ParkingRepository interface {
	Transaction(ctx context.Context, f ParkingTxFunc)
	FindByIdOrDefault(ctx context.Context, id parking_entity.CarPlate) parking_entity.Parking
	Save(ctx context.Context, parking parking_entity.Parking)
}

type ParkingTxFunc func(ctx context.Context, repo ParkingRepository)

type AlarmService interface {
	Alarm(ctx context.Context, plate parking_entity.CarPlate, message string, alarmTime parking_entity.Time)
}
