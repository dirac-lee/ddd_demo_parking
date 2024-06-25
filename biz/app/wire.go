//go:build wireinject
// +build wireinject

package app

import (
	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/id_gen"
	parking_domain "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain"
	parking_infra_impl "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/infra"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/dao"
	"github.com/google/wire"
)

func Setup() ParkingApp {
	panic(
		wire.Build(
			id_gen.New, wire.Bind(new(domain.IDGen), new(id_gen.IDGen)),
			db.NewDB,
			dao.ProvideDao,
			parking_infra_impl.ProvideRepo,
			parking_infra_impl.ProvideService,
			parking_infra_impl.ProvideListeners,
			parking_infra_impl.ProvideSyncEvent,
			parking_domain.ProvideCommandHandlers,
			ProvideApp,
		),
	)
}
