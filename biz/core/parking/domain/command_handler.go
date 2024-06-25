package parking_domain

import "github.com/google/wire"

var ProvideCommandHandlers = wire.NewSet(
	wire.Struct(new(CheckInCommandHandler), "*"),
	wire.Struct(new(CheckOutCommandHandler), "*"),
	wire.Struct(new(CalculateFeeCommandHandler), "*"),
	wire.Struct(new(NotifyPaidCommandHandler), "*"),
)
