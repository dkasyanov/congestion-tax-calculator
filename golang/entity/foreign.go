package entity

type Foreign struct{}

func (c Foreign) GetVehicleType() string {
	return "Foreign"
}

func (c Foreign) IsTollFree() bool {
	return true
}
