package parking_infra_impl

import (
	"context"

	"code.byted.org/oec/status_code/Status"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/domain"
	E "github.com/dirac-lee/ddd_demo_parking/biz/common/exception"
	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	parking_infra "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/infra"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/dao"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/model"
	"github.com/google/wire"
)

var ProvideService = wire.NewSet(
	wire.Struct(new(AlarmService), "*"), wire.Bind(new(parking_infra.AlarmService), new(AlarmService)),
)

type AlarmService struct {
	Dao   *dao.Query
	IDGen domain.IDGen
}

func (s AlarmService) Alarm(ctx context.Context, plate parking_entity.CarPlate, message string, alarmTime parking_entity.Time) {
	alarmDao := s.Dao.AlarmPo
	po := model.AlarmPo{
		ID:      s.IDGen.GetID(ctx),
		Plate:   plate.String(),
		Message: message,
		Time:    alarmTime.Into(),
	}
	err := alarmDao.WithContext(ctx).Create(&po)
	if err != nil {
		E.Throw(E.New(Status.DBCreateException, E.WithCause(err)))
	}
}
