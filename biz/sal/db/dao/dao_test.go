package dao

import (
	"context"
	"testing"

	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db"
	"github.com/luci/go-render/render"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFirstOrInit(t *testing.T) {
	Convey(t.Name(), t, func() {
		dao := New(db.NewDB())
		po, err := dao.ParkingDao.Where(dao.Q.ParkingPo.ID.Eq("non_existing")).FirstOrInit()
		So(err, ShouldBeNil)
		So(po.ID, ShouldNotBeEmpty)
		t.Log(render.Render(po))

		ctx := context.Background()
		po, err = dao.ParkingDao.WithContext(ctx).Where(dao.Q.ParkingPo.ID.Eq("New1234")).FirstOrInit()
		So(err, ShouldBeNil)
		So(po.ID, ShouldNotBeEmpty)
		t.Log(po.CheckInTime)
	})
}
