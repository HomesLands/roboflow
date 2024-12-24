// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FakePublisher is an autogenerated mock type for the Publisher type
type FakePublisher struct {
	mock.Mock
}

type FakePublisher_Expecter struct {
	mock *mock.Mock
}

func (_m *FakePublisher) EXPECT() *FakePublisher_Expecter {
	return &FakePublisher_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: topic, payload
func (_m *FakePublisher) Publish(topic string, payload []byte) error {
	ret := _m.Called(topic, payload)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(topic, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FakePublisher_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type FakePublisher_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - topic string
//   - payload []byte
func (_e *FakePublisher_Expecter) Publish(topic interface{}, payload interface{}) *FakePublisher_Publish_Call {
	return &FakePublisher_Publish_Call{Call: _e.mock.On("Publish", topic, payload)}
}

func (_c *FakePublisher_Publish_Call) Run(run func(topic string, payload []byte)) *FakePublisher_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]byte))
	})
	return _c
}

func (_c *FakePublisher_Publish_Call) Return(_a0 error) *FakePublisher_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FakePublisher_Publish_Call) RunAndReturn(run func(string, []byte) error) *FakePublisher_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// NewFakePublisher creates a new instance of FakePublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFakePublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *FakePublisher {
	mock := &FakePublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}