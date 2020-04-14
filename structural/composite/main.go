package main

import (
	"fmt"
	"github.com/justym/design-pattern/structural/composite/model"
)

func embedStruct() {
	swimmer := model.StructSwimmer{
		MySwim: Swim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	fish := model.Fish{
		Swim: Swim,
	}
	fish.Animal.Eat()
	fish.Swim()
}

func embedInterface() {
	swimmer := model.InterfaceSwimmer{
		Swimmer:&model.SwimmerImpl{},
		Trainer:&model.Athlete{},
	}

	swimmer.Swim()
	swimmer.Train()
}

func Swim() {
	fmt.Println("Athlete or Animal: Swimming")
}

func main() {
	embedStruct()
	embedInterface()
}
