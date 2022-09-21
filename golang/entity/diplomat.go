package entity

type Diplomat struct{}

func (c Diplomat) GetVehicleType() string {
	return "Diplomat"
}

func (c Diplomat) IsTollFree() bool {
	return true
}
