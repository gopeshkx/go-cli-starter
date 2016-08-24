package mocks

import "github.com/stretchr/testify/mock"

// Countable is an autogenerated mock type for the Countable type
type Countable struct {
	mock.Mock
}

// Inc provides a mock function with given fields:
func (_m *Countable) Inc() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}