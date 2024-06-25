package app

var (
	parkingApp ParkingApp
)

func Init() {
	parkingApp = Setup()
}

func GetParkingApp() ParkingApp {
	return parkingApp
}
