package abstract

import (
	"errors"
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType = iota
	MotorbikeFactoryType
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorBikeFactory), nil
	default:
		return nil, errors.New("This Factory type is not recognized")
	}
}
