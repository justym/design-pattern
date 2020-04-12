package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance(), no nil")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling counter1.AddOne() for the first time, it must be 1")
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Errorf("counter2 must be a instance as same as expectedCount")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling counter2.AddOne(), it must be 2. Because counter2 shared instance with counter1")
	}
}
