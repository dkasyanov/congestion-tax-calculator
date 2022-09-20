package model

type Car struct{}

func (c Car) GetVehicleType() string {
	return "Car"
}

func (c Car) IsTollFree() bool {
	return false
}
