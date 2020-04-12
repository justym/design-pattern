package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4, but it's %v", &car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5, but it's %v", &car.Seats)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure must be Car, but it's %v", &car.Structure)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.Build()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a car must be 2, but it's %v", &bike.Wheels)
	}

	if bike.Seats != 1 {
		t.Errorf("Seats on a bike must be 1, but it's %v", &bike.Seats)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure must be Bike, but it's %v", &bike.Structure)
	}

}
