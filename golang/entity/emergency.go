package entity

type Emergency struct{}

func (c Emergency) GetVehicleType() string {
	return "Emergency"
}

func (c Emergency) IsTollFree() bool {
	return true
}
