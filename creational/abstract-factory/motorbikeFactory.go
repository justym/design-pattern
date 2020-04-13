package abstract

import "errors"

const (
	SportMotorBikeType = iota
	CruiseMotorBikeType
)

type MotorBikeFactory struct{}

func (b *MotorBikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorBikeType:
		return new(SportMotorBike), nil
	case CruiseMotorBikeType:
		return new(CruiseMotorBike), nil
	default:
		return nil, errors.New("That Vehicle type is not found\n")
	}
}
