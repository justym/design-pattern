package model

import "fmt"

//Athlete Section
type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Athlete: Training")
}

type StructSwimmer struct {
	MyAthlete Athlete
	MySwim    func()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

type InterfaceSwimmer struct {
	Swimmer
	Trainer
}

//Animal Section
type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Animal: Eating")
}

type Fish struct {
	Animal
	Swim func()
}
