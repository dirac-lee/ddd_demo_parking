package parking_infra_impl

import (
	"context"

	"code.byted.org/oec/status_code/Status"
	E "github.com/dirac-lee/ddd_demo_parking/biz/common/exception"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/logx"
	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	parking_infra "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/infra"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/dao"
	"github.com/google/wire"
)

var ProvideRepo = wire.NewSet(
	wire.Struct(new(ParkingBuilder), "*"),
	wire.Struct(new(ParkingRepository), "*"), wire.Bind(new(parking_infra.ParkingRepository), new(*ParkingRepository)),
)

type ParkingRepository struct {
	Dao            *dao.Query
	ParkingBuilder ParkingBuilder
}

func (p ParkingRepository) Transaction(ctx context.Context, f parking_infra.ParkingTxFunc) {
	if err := p.Dao.Transaction(func(tx *dao.Query) error {
		var err error
		repoTx := ParkingRepository{
			Dao:            tx,
			ParkingBuilder: p.ParkingBuilder,
		}
		E.Try(func() {
			f(ctx, repoTx)
		}).Catch(func(e *E.Exception) {
			err = e
		})
		return err
	}); err != nil {
		E.Throw(E.New(Status.DBTransactionException, E.WithCause(err)))
	}
}

func (p ParkingRepository) FindByIdOrDefault(ctx context.Context, id parking_entity.CarPlate) (parking parking_entity.Parking) {
	logx.InfoInput(ctx, map[string]any{"id": id})
	defer logx.InfoOutput(ctx, map[string]any{"parking": parking})
	po, err := p.Dao.ParkingPo.WithContext(ctx).FirstOrInit()
	if err != nil {
		E.Throw(E.New(Status.DBQueryException, E.WithCause(err)))
	}
	if po.ID == "" {
		po.ID = string(id)
	}
	return p.ParkingBuilder.ToEntity(po)
}

func (p ParkingRepository) Save(ctx context.Context, parking parking_entity.Parking) {
	logx.InfoInput(ctx, map[string]any{"parking": parking})
	po := p.ParkingBuilder.FromEntity(parking)
	if err := p.Dao.ParkingPo.WithContext(ctx).Save(po); err != nil {
		E.Throw(E.New(Status.DBCreateException, E.WithCause(err)))
	}
}