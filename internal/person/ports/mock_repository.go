// Code generated by mockery v2.20.0. DO NOT EDIT.

package ports

import (
	domain "github.com/danilopimenta/family-tree-app/internal/person/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the InmemPersonRepository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// Find provides a mock function with given fields: _a0
func (_m *MockRepository) FindOne(_a0 string) (*domain.Person, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Person, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Person); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type MockRepository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - _a0 string
func (_e *MockRepository_Expecter) Find(_a0 interface{}) *MockRepository_Find_Call {
	return &MockRepository_Find_Call{Call: _e.mock.On("FindOne", _a0)}
}

func (_c *MockRepository_Find_Call) Run(run func(_a0 string)) *MockRepository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRepository_Find_Call) Return(_a0 *domain.Person, _a1 error) *MockRepository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Find_Call) RunAndReturn(run func(string) (*domain.Person, error)) *MockRepository_Find_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *MockRepository) FindAll() []domain.Person {
	ret := _m.Called()

	var r0 []domain.Person
	if rf, ok := ret.Get(0).(func() []domain.Person); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Person)
		}
	}

	return r0
}

// MockRepository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockRepository_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockRepository_Expecter) FindAll() *MockRepository_FindAll_Call {
	return &MockRepository_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockRepository_FindAll_Call) Run(run func()) *MockRepository_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRepository_FindAll_Call) Return(_a0 []domain.Person) *MockRepository_FindAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_FindAll_Call) RunAndReturn(run func() []domain.Person) *MockRepository_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindAncestors provides a mock function with given fields: _a0
func (_m *MockRepository) FindAncestors(_a0 string) ([]domain.Person, error) {
	ret := _m.Called(_a0)

	var r0 []domain.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.Person, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Person); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindAncestors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAncestors'
type MockRepository_FindAncestors_Call struct {
	*mock.Call
}

// FindAncestors is a helper method to define mock.On call
//   - _a0 string
func (_e *MockRepository_Expecter) FindAncestors(_a0 interface{}) *MockRepository_FindAncestors_Call {
	return &MockRepository_FindAncestors_Call{Call: _e.mock.On("FindAncestors", _a0)}
}

func (_c *MockRepository_FindAncestors_Call) Run(run func(_a0 string)) *MockRepository_FindAncestors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRepository_FindAncestors_Call) Return(_a0 []domain.Person, _a1 error) *MockRepository_FindAncestors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindAncestors_Call) RunAndReturn(run func(string) ([]domain.Person, error)) *MockRepository_FindAncestors_Call {
	_c.Call.Return(run)
	return _c
}

// FindByName provides a mock function with given fields: _a0
func (_m *MockRepository) FindByName(_a0 string) (domain.Person, error) {
	ret := _m.Called(_a0)

	var r0 domain.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.Person, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) domain.Person); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.Person)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByName'
type MockRepository_FindByName_Call struct {
	*mock.Call
}

// FindByName is a helper method to define mock.On call
//   - _a0 string
func (_e *MockRepository_Expecter) FindByName(_a0 interface{}) *MockRepository_FindByName_Call {
	return &MockRepository_FindByName_Call{Call: _e.mock.On("FindByName", _a0)}
}

func (_c *MockRepository_FindByName_Call) Run(run func(_a0 string)) *MockRepository_FindByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRepository_FindByName_Call) Return(_a0 domain.Person, _a1 error) *MockRepository_FindByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindByName_Call) RunAndReturn(run func(string) (domain.Person, error)) *MockRepository_FindByName_Call {
	_c.Call.Return(run)
	return _c
}

// StoreRelationship provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) StoreRelationship(_a0 *domain.Person, _a1 *domain.Person) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Person, *domain.Person) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_StoreRelationship_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreRelationship'
type MockRepository_StoreRelationship_Call struct {
	*mock.Call
}

// StoreRelationship is a helper method to define mock.On call
//   - _a0 *domain.Person
//   - _a1 *domain.Person
func (_e *MockRepository_Expecter) StoreRelationship(_a0 interface{}, _a1 interface{}) *MockRepository_StoreRelationship_Call {
	return &MockRepository_StoreRelationship_Call{Call: _e.mock.On("StoreRelationship", _a0, _a1)}
}

func (_c *MockRepository_StoreRelationship_Call) Run(run func(_a0 *domain.Person, _a1 *domain.Person)) *MockRepository_StoreRelationship_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Person), args[1].(*domain.Person))
	})
	return _c
}

func (_c *MockRepository_StoreRelationship_Call) Return(_a0 error) *MockRepository_StoreRelationship_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_StoreRelationship_Call) RunAndReturn(run func(*domain.Person, *domain.Person) error) *MockRepository_StoreRelationship_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
