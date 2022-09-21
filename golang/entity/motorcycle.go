package entity

type Motorcycle struct {
}

func (m Motorcycle) GetVehicleType() string {
	return "Motorcycle"
}

func (m Motorcycle) IsTollFree() bool {
	return true
}
