package app

import (
	"time"

	"code.byted.org/lang/gg/gptr"
	"code.byted.org/lang/gg/gslice"
	E "github.com/dirac-lee/ddd_demo_parking/biz/common/exception"
	parking_entity "github.com/dirac-lee/ddd_demo_parking/biz/core/parking/domain/entity"
	"github.com/dirac-lee/ddd_demo_parking/biz/model/ddd_demo_parking"
	"github.com/dirac-lee/ddd_demo_parking/biz/sal/db/model"
)

type DtoAssembler struct {
}

func (da DtoAssembler) ToCheckInCommand(req *ddd_demo_parking.CheckInRequest) parking_entity.CheckInCommand {
	return parking_entity.CheckInCommand{
		CarPlate:    parking_entity.CarPlate(req.GetCarPlate()),
		CheckInTime: parking_entity.IntoTime(time.Now()),
	}
}

func (da DtoAssembler) ToCheckInResponse(success bool, ex *E.Exception) *ddd_demo_parking.CheckInResponse {
	return &ddd_demo_parking.CheckInResponse{
		Data: &ddd_demo_parking.CheckInResponseData{
			Success: gptr.Of(success),
		},
		Code:    ex.CodePtr(),
		Message: ex.MessagePtr(),
	}
}

func (da DtoAssembler) ToCheckOutCommand(req *ddd_demo_parking.CheckOutRequest) parking_entity.CheckOutCommand {
	return parking_entity.CheckOutCommand{
		CarPlate:     parking_entity.CarPlate(req.GetCarPlate()),
		CheckOutTime: parking_entity.IntoTime(time.Now()),
	}
}

func (da DtoAssembler) ToCheckOutResponse(success bool, ex *E.Exception) *ddd_demo_parking.CheckOutResponse {
	return &ddd_demo_parking.CheckOutResponse{
		Data: &ddd_demo_parking.CheckOutResponseData{
			Success: gptr.Of(success),
		},
		Code:    ex.CodePtr(),
		Message: ex.MessagePtr(),
	}

}

func (da DtoAssembler) ToCalculateFeeCommand(req *ddd_demo_parking.CalculateFeeRequest) parking_entity.CalculateFeeCommand {
	return parking_entity.CalculateFeeCommand{
		CarPlate:         parking_entity.CarPlate(req.GetCarPlate()),
		CalculateFeeTime: parking_entity.IntoTime(time.Now()),
	}

}

func (da DtoAssembler) ToCalculateFeeResponse(fee int64, ex *E.Exception) *ddd_demo_parking.CalculateFeeResponse {
	return &ddd_demo_parking.CalculateFeeResponse{
		Data: &ddd_demo_parking.CalculateFeeResponseData{
			Fee: gptr.Of(fee),
		},
		Code:    ex.CodePtr(),
		Message: ex.MessagePtr(),
	}

}

func (da DtoAssembler) ToNotifyPayCommand(req *ddd_demo_parking.NotifyPaidRequest) parking_entity.NotifyPaidCommand {
	return parking_entity.NotifyPaidCommand{
		CarPlate: parking_entity.CarPlate(req.GetCarPlate()),
		Amount:   req.GetPayAmount(),
		PayTime:  parking_entity.IntoTime(time.Now()),
	}
}

func (da DtoAssembler) ToNotifyPaidResponse(success bool, ex *E.Exception) *ddd_demo_parking.NotifyPaidResponse {
	return &ddd_demo_parking.NotifyPaidResponse{
		Data: &ddd_demo_parking.NotifyPaidResponseData{
			Success: gptr.Of(success),
		},
		Code:    ex.CodePtr(),
		Message: ex.MessagePtr(),
	}
}

func (da DtoAssembler) ToGetHistoryResponse(pos []*model.ParkingViewPo, ex *E.Exception) *ddd_demo_parking.GetHistoryResponse {
	return &ddd_demo_parking.GetHistoryResponse{
		Data: &ddd_demo_parking.GetHistoryResponseData{
			ParkingRecords: gslice.Map(pos, da.viewPoToVo),
		},
		Code:    ex.CodePtr(),
		Message: ex.MessagePtr(),
	}
}

func (da DtoAssembler) viewPoToVo(po *model.ParkingViewPo) *ddd_demo_parking.ParkingRecord {
	return &ddd_demo_parking.ParkingRecord{
		CarPlate:          gptr.Of(po.Plate),
		CheckInTimestamp:  gptr.Of(po.CheckInTime.Unix()),
		CheckOutTimestamp: gptr.Of(po.CheckOutTime.Unix()),
		PayAmount:         gptr.Of(po.PayAmount),
	}
}

func (da DtoAssembler) ToGetTotalInParkResponse(po *model.SummaryPo, ex *E.Exception) *ddd_demo_parking.GetTotalInParkResponse {
	return &ddd_demo_parking.GetTotalInParkResponse{
		Data: &ddd_demo_parking.GetTotalInParkResponseData{
			TotalInPark: gptr.Of(po.TotalInParking),
		},
		Code:    ex.CodePtr(),
		Message: ex.MessagePtr(),
	}
}

func (da DtoAssembler) ToGetDailyRevenueResponse(pos []*model.DailyRevenuePo, ex *E.Exception) *ddd_demo_parking.GetDailyRevenueResponse {
	return &ddd_demo_parking.GetDailyRevenueResponse{
		Data: &ddd_demo_parking.GetDailyRevenueResponseData{
			DailyRevenues: gslice.Map(pos, da.dailyRevenuePoToVo),
		},
		Code:    ex.CodePtr(),
		Message: ex.MessagePtr(),
	}
}

func (da DtoAssembler) dailyRevenuePoToVo(po *model.DailyRevenuePo) *ddd_demo_parking.DailyRevenue {
	return &ddd_demo_parking.DailyRevenue{
		Date:    gptr.Of(po.Date),
		Revenue: gptr.Of(po.Revenue),
	}
}
