package parking_infra_impl

import (
	"context"

	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
	parking_policy "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/policy"
	"github.com/google/wire"
)

var ProvideSyncEvent = wire.NewSet(
	wire.Struct(new(EventListenerRegistration), "*"), RegisterListeners,
	wire.Struct(new(SyncEventBroker), "*"),
	wire.Struct(new(SyncEventPublisher), "*"),
	wire.Bind(new(domain.EventPublisher), new(SyncEventPublisher)),
)

type EventListenerRegistration struct {
	AlarmPolicy          parking_policy.AlarmPolicy
	SummaryListener      SummaryListener
	ParkingViewListener  ParkingViewListener
	DailyRevenueListener DailyRevenueListener
}

func RegisterListeners(registration EventListenerRegistration) []domain.EventListener {
	return []domain.EventListener{
		registration.AlarmPolicy,
		registration.SummaryListener,
		registration.ParkingViewListener,
		registration.DailyRevenueListener,
	}
}

type SyncEventPublisher struct {
	EventBroker SyncEventBroker
}

func (ep SyncEventPublisher) PublishEvent(ctx context.Context, event domain.Event) {
	ep.EventBroker.NotifyAll(ctx, event)
}

type SyncEventBroker struct {
	EventListeners []domain.EventListener
}

func (eb SyncEventBroker) NotifyAll(ctx context.Context, event domain.Event) {
	for _, listener := range eb.EventListeners {
		listener.OnEvent(ctx, event)
	}
}
