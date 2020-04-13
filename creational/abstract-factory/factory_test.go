package abstract

import (
	"testing"
)

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car Vehicle has %d seats\n", carVehicle.NumSeats())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct Assertion has failed\n")
	}

	t.Logf("Luxury Car has %d dorrs.\n", luxuryCar.NumDoors())
}

func TestMotorBikeFactory(t *testing.T) {
	bikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	mb, err := bikeF.NewVehicle(SportMotorBikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motor Bike Vehicle has %d seats\n", mb.NumSeats())

	sportMB, ok := mb.(MotorBike)
	if !ok {
		t.Fatal("Struct Assertion has failed\n")
	}

	t.Logf("Sport Motor Bike has type %d\n", sportMB.GetMotorbikeType())
}
