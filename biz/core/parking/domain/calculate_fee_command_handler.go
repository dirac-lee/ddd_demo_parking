package parking_domain

import (
	"context"

	"github.com/dirac-lee/ddd_demo_parking/biz/common/logx"
	parking_entity2 "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	parking_infra "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/infra"
)

type CalculateFeeCommandHandler struct {
	ParkingRepository parking_infra.ParkingRepository
}

func (h CalculateFeeCommandHandler) HandleCalculateFeeCommand(ctx context.Context, command parking_entity2.CalculateFeeCommand) (fee int64) {
	logx.InfoInput(ctx, map[string]any{"command": command})
	defer func() { logx.InfoOutput(ctx, map[string]any{"fee": fee}) }()
	h.ParkingRepository.Transaction(ctx, func(ctx context.Context, repo parking_infra.ParkingRepository) {
		var parking parking_entity2.Parking
		parking = repo.FindByIdOrDefault(ctx, command.CarPlate)
		fee = parking.CalculateFee(ctx, command.CalculateFeeTime)
	})
	return fee
}
