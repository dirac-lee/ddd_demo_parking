package parking_entity

type CheckInCommand struct {
	CarPlate    CarPlate
	CheckInTime Time
}

type CheckOutCommand struct {
	CarPlate     CarPlate
	CheckOutTime Time
}

type CalculateFeeCommand struct {
	CarPlate         CarPlate
	CalculateFeeTime Time
}

type NotifyPaidCommand struct {
	CarPlate CarPlate
	Amount   int64
	PayTime  Time
}
