package model

type Vehicle interface {
	GetVehicleType() string
	IsTollFree() bool
}

func NewVehicle(name string) Vehicle {
	switch name {
	case "Car":
		return Car{}
	case "Motorbike":
		return Motorbike{}
	}
	return nil
}
