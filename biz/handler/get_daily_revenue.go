// Code generated by hertztool.

package handler

import (
	"context"

	hertz_app "code.byted.org/middleware/hertz/pkg/app"
	"code.byted.org/middleware/hertz_ext/v2/binding"
	"github.com/dirac-lee/ddd_demo_parking/biz/app"
	"github.com/dirac-lee/ddd_demo_parking/biz/model/ddd_demo_parking"
)

// GetDailyRevenue .
// @router /api/ddd/demo/parking/daily_revenue [GET]
func GetDailyRevenue(ctx context.Context, c *hertz_app.RequestContext) {
	var err error
	var req ddd_demo_parking.GetDailyRevenueRequest
	err = binding.BindAndValidate(c, &req)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	resp := app.GetParkingApp().GetDailyRevenue(ctx, &req)
	err = binding.WriteHeader(c, resp)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, resp)
}
