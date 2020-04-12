//This example creates A Unique Counter Implementation

package creational

//Singleton is interface have AddOne() method
type Singleton interface {
	AddOne() int
}

//singleton is implementation of Singleton
type singleton struct {
	count int
}

//instance cannot export from this pacakge
var instance *singleton

//GetInstance create Singleton Instance.
//If instance variable is nil(At first, instance variable is nil) return instance
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

//AddOne add 1 to singleton.count and return it.
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
