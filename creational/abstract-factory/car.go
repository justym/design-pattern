package abstract

type Car interface {
	NumDoors() int
}

//LuxuryCar section
type LuxuryCar struct{}

func (l *LuxuryCar) NumDoors() int {
	return 4
}

func (l *LuxuryCar) NumWheels() int {
	return 4
}

func (l *LuxuryCar) NumSeats() int {
	return 5
}

//FamilyCar section
type FamilyCar struct{}

func (f *FamilyCar) NumDoors() int {
	return 5
}

func (f *FamilyCar) NumWheels() int {
	return 4
}

func (f *FamilyCar) NumSeats() int {
	return 5
}
