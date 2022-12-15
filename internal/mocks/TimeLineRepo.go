// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/internal/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"

	query "github.com/bangumi/server/internal/dal/query"
)

// TimeLineRepo is an autogenerated mock type for the TimeLineRepo type
type TimeLineRepo struct {
	mock.Mock
}

type TimeLineRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *TimeLineRepo) EXPECT() *TimeLineRepo_Expecter {
	return &TimeLineRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, tl
func (_m *TimeLineRepo) Create(ctx context.Context, tl *model.TimeLine) error {
	ret := _m.Called(ctx, tl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.TimeLine) error); ok {
		r0 = rf(ctx, tl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TimeLineRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type TimeLineRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - tl *model.TimeLine
func (_e *TimeLineRepo_Expecter) Create(ctx interface{}, tl interface{}) *TimeLineRepo_Create_Call {
	return &TimeLineRepo_Create_Call{Call: _e.mock.On("Create", ctx, tl)}
}

func (_c *TimeLineRepo_Create_Call) Run(run func(ctx context.Context, tl *model.TimeLine)) *TimeLineRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.TimeLine))
	})
	return _c
}

func (_c *TimeLineRepo_Create_Call) Return(_a0 error) *TimeLineRepo_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TimeLineRepo) GetByID(ctx context.Context, id model.TimeLineID) (*model.TimeLine, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.TimeLine
	if rf, ok := ret.Get(0).(func(context.Context, model.TimeLineID) *model.TimeLine); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TimeLine)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.TimeLineID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeLineRepo_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type TimeLineRepo_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id model.TimeLineID
func (_e *TimeLineRepo_Expecter) GetByID(ctx interface{}, id interface{}) *TimeLineRepo_GetByID_Call {
	return &TimeLineRepo_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *TimeLineRepo_GetByID_Call) Run(run func(ctx context.Context, id model.TimeLineID)) *TimeLineRepo_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.TimeLineID))
	})
	return _c
}

func (_c *TimeLineRepo_GetByID_Call) Return(_a0 *model.TimeLine, _a1 error) *TimeLineRepo_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListByUID provides a mock function with given fields: ctx, uid, limit, since
func (_m *TimeLineRepo) ListByUID(ctx context.Context, uid model.UserID, limit int, since model.TimeLineID) ([]*model.TimeLine, error) {
	ret := _m.Called(ctx, uid, limit, since)

	var r0 []*model.TimeLine
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, int, model.TimeLineID) []*model.TimeLine); ok {
		r0 = rf(ctx, uid, limit, since)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TimeLine)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, int, model.TimeLineID) error); ok {
		r1 = rf(ctx, uid, limit, since)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeLineRepo_ListByUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByUID'
type TimeLineRepo_ListByUID_Call struct {
	*mock.Call
}

// ListByUID is a helper method to define mock.On call
//   - ctx context.Context
//   - uid model.UserID
//   - limit int
//   - since model.TimeLineID
func (_e *TimeLineRepo_Expecter) ListByUID(ctx interface{}, uid interface{}, limit interface{}, since interface{}) *TimeLineRepo_ListByUID_Call {
	return &TimeLineRepo_ListByUID_Call{Call: _e.mock.On("ListByUID", ctx, uid, limit, since)}
}

func (_c *TimeLineRepo_ListByUID_Call) Run(run func(ctx context.Context, uid model.UserID, limit int, since model.TimeLineID)) *TimeLineRepo_ListByUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(int), args[3].(model.TimeLineID))
	})
	return _c
}

func (_c *TimeLineRepo_ListByUID_Call) Return(_a0 []*model.TimeLine, _a1 error) *TimeLineRepo_ListByUID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// WithQuery provides a mock function with given fields: _a0
func (_m *TimeLineRepo) WithQuery(_a0 *query.Query) domain.TimeLineRepo {
	ret := _m.Called(_a0)

	var r0 domain.TimeLineRepo
	if rf, ok := ret.Get(0).(func(*query.Query) domain.TimeLineRepo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.TimeLineRepo)
		}
	}

	return r0
}

// TimeLineRepo_WithQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithQuery'
type TimeLineRepo_WithQuery_Call struct {
	*mock.Call
}

// WithQuery is a helper method to define mock.On call
//   - _a0 *query.Query
func (_e *TimeLineRepo_Expecter) WithQuery(_a0 interface{}) *TimeLineRepo_WithQuery_Call {
	return &TimeLineRepo_WithQuery_Call{Call: _e.mock.On("WithQuery", _a0)}
}

func (_c *TimeLineRepo_WithQuery_Call) Run(run func(_a0 *query.Query)) *TimeLineRepo_WithQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*query.Query))
	})
	return _c
}

func (_c *TimeLineRepo_WithQuery_Call) Return(_a0 domain.TimeLineRepo) *TimeLineRepo_WithQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewTimeLineRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewTimeLineRepo creates a new instance of TimeLineRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTimeLineRepo(t mockConstructorTestingTNewTimeLineRepo) *TimeLineRepo {
	mock := &TimeLineRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
