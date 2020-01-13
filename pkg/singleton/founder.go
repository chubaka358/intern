package singleton

import (
	"sync"
)

var (
	instance *founder
	once     sync.Once
)

type Founderer interface {
	SetName(name string)
	GetName() string
	SetAge(age uint)
}

// founder representation
type founder struct {
	name string
	age  uint
}

// SetName set founder name
func (f *founder) SetName(name string) {
	f.name = name
}

// GetName return founder name
func (f *founder) GetName() string {
	return f.name
}

// SetAge set founder age
func (f *founder) SetAge(age uint) {
	f.age = age
}

// GetAge return founder age
func (f *founder) GetAge() uint {
	return f.age
}

// GetFounder return founder instance
func NewFounder() *founder {
	once.Do(func() {
		instance = &founder{}
	})
	return instance
}
