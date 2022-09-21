package entity

type Vehicle interface {
	GetVehicleType() string
	IsTollFree() bool
}

func NewVehicle(name string) Vehicle {
	switch name {
	case "Car":
		return Car{}
	case "Motorcycle":
		return Motorcycle{}
	case "Tractor":
		return Tractor{}
	case "Emergency":
		return Emergency{}
	case "Diplomat":
		return Diplomat{}
	case "Foreign":
		return Foreign{}
	case "Military":
		return Military{}
	default:
		return nil
	}
}
