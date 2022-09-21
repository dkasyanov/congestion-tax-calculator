package entity

type Military struct{}

func (c Military) GetVehicleType() string {
	return "Military"
}

func (c Military) IsTollFree() bool {
	return true
}
