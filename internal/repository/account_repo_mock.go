// Code generated by mockery v2.43.2. DO NOT EDIT.

package repository

import (
	context "context"

	models "github.com/gopay/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// MockAccountRepo is an autogenerated mock type for the AccountRepo type
type MockAccountRepo struct {
	mock.Mock
}

type MockAccountRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountRepo) EXPECT() *MockAccountRepo_Expecter {
	return &MockAccountRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, name, lastname
func (_m *MockAccountRepo) Create(ctx context.Context, name string, lastname string) (string, error) {
	ret := _m.Called(ctx, name, lastname)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, name, lastname)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, name, lastname)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, lastname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - lastname string
func (_e *MockAccountRepo_Expecter) Create(ctx interface{}, name interface{}, lastname interface{}) *MockAccountRepo_Create_Call {
	return &MockAccountRepo_Create_Call{Call: _e.mock.On("Create", ctx, name, lastname)}
}

func (_c *MockAccountRepo_Create_Call) Run(run func(ctx context.Context, name string, lastname string)) *MockAccountRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockAccountRepo_Create_Call) Return(_a0 string, _a1 error) *MockAccountRepo_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepo_Create_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *MockAccountRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields: ctx
func (_m *MockAccountRepo) FindAll(ctx context.Context) ([]models.Account, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []models.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Account, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Account); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepo_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockAccountRepo_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccountRepo_Expecter) FindAll(ctx interface{}) *MockAccountRepo_FindAll_Call {
	return &MockAccountRepo_FindAll_Call{Call: _e.mock.On("FindAll", ctx)}
}

func (_c *MockAccountRepo_FindAll_Call) Run(run func(ctx context.Context)) *MockAccountRepo_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccountRepo_FindAll_Call) Return(_a0 []models.Account, _a1 error) *MockAccountRepo_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepo_FindAll_Call) RunAndReturn(run func(context.Context) ([]models.Account, error)) *MockAccountRepo_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindOne provides a mock function with given fields: ctx, id
func (_m *MockAccountRepo) FindOne(ctx context.Context, id string) (models.Account, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindOne")
	}

	var r0 models.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepo_FindOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type MockAccountRepo_FindOne_Call struct {
	*mock.Call
}

// FindOne is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockAccountRepo_Expecter) FindOne(ctx interface{}, id interface{}) *MockAccountRepo_FindOne_Call {
	return &MockAccountRepo_FindOne_Call{Call: _e.mock.On("FindOne", ctx, id)}
}

func (_c *MockAccountRepo_FindOne_Call) Run(run func(ctx context.Context, id string)) *MockAccountRepo_FindOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountRepo_FindOne_Call) Return(_a0 models.Account, _a1 error) *MockAccountRepo_FindOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepo_FindOne_Call) RunAndReturn(run func(context.Context, string) (models.Account, error)) *MockAccountRepo_FindOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountRepo creates a new instance of MockAccountRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountRepo {
	mock := &MockAccountRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
