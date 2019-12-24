package singleton

import (
	"sync"
)

type founder struct {
	name string
	age  uint
}

var (
	instance *founder
	once     sync.Once
)

func (f *founder) SetName(name string) {
	f.name = name
}

func (f *founder) GetName() string {
	return f.name
}

func (f *founder) SetAge(age uint) {
	f.age = age
}

func (f *founder) GetAge() uint {
	return f.age
}

func GetFounder() *founder {
	once.Do(func() {
		instance = &founder{}
	})
	return instance
}
