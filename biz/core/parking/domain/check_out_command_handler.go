package parking_domain

import (
	"context"

	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/logx"
	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	parking_infra "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/infra"
)

type CheckOutCommandHandler struct {
	ParkingRepository parking_infra.ParkingRepository
	EventPublisher    domain.EventPublisher
}

func (h CheckOutCommandHandler) Handle(ctx context.Context, command parking_entity.CheckOutCommand) (success bool) {
	logx.InfoInput(ctx, map[string]any{"command": command})
	defer func() { logx.InfoOutput(ctx, map[string]any{"success": success}) }()
	parking := h.ParkingRepository.FindByIdOrDefault(ctx, command.CarPlate)
	success = parking.HandleCheckOutCommand(ctx, h.EventPublisher, command)
	h.ParkingRepository.Save(ctx, parking)
	return success
}
