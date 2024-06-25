package parking_domain

import (
	"context"

	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/logx"
	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	parking_infra "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/infra"
)

type NotifyPaidCommandHandler struct {
	ParkingRepository parking_infra.ParkingRepository
	EventPublisher    domain.EventPublisher
}

func (h NotifyPaidCommandHandler) HandleNotifyPayCommand(ctx context.Context, command parking_entity.NotifyPaidCommand) (success bool) {
	logx.InfoInput(ctx, map[string]any{"command": command})
	defer logx.InfoOutput(ctx, map[string]any{"success": success})
	parking := h.ParkingRepository.FindByIdOrDefault(ctx, command.CarPlate)
	parking.HandleNotifyPayCommand(ctx, h.EventPublisher, command)
	h.ParkingRepository.Save(ctx, parking)
	return true
}
