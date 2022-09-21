package entity

type Tractor struct{}

func (c Tractor) GetVehicleType() string {
	return "Tractor"
}

func (c Tractor) IsTollFree() bool {
	return true
}
