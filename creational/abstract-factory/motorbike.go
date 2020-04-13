package abstract

type MotorBike interface {
	GetMotorbikeType() int
}

type SportMotorBike struct{}

func (s *SportMotorBike) NumWheels() int {
	return 2
}

func (s *SportMotorBike) NumSeats() int {
	return 1
}

func (s *SportMotorBike) GetMotorbikeType() int {
	return SportMotorBikeType
}

type CruiseMotorBike struct{}

func (c *CruiseMotorBike) NumWheels() int {
	return 2
}

func (c *CruiseMotorBike) NumSeats() int {
	return 2
}

func (c *CruiseMotorBike) GetMotorbikeType() int {
	return CruiseMotorBikeType
}
