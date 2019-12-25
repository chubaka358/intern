package singleton

import (
	"sync"
)

// founder representation
type Founder struct {
	name string
	age  uint
}

var (
	instance *Founder
	once     sync.Once
)

// SetName set founder name
func (f *Founder) SetName(name string) {
	f.name = name
}

// GetName return founder name
func (f *Founder) GetName() string {
	return f.name
}

// SetAge set founder age
func (f *Founder) SetAge(age uint) {
	f.age = age
}

// GetAge return founder age
func (f *Founder) GetAge() uint {
	return f.age
}

// GetFounder return founder instance
func GetFounder() *Founder {
	once.Do(func() {
		instance = &Founder{}
	})
	return instance
}
