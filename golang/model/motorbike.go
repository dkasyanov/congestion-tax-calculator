package model

type Motorbike struct {
}

func (m Motorbike) GetVehicleType() string {
	return "Motorbike"
}

func (m Motorbike) IsTollFree() bool {
	return true
}
