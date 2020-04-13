package abstract

import "errors"

const (
	LuxuryCarType = iota
	FamilyCarType
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New("That Vehicle type is not found\n")
	}
}
