// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tcscheurer/rentals/db/sqlc (interfaces: Querier)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlc "github.com/tcscheurer/rentals/db/sqlc"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// GetRentalByID mocks base method.
func (m *MockQuerier) GetRentalByID(arg0 context.Context, arg1 int32) (sqlc.GetRentalByIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRentalByID", arg0, arg1)
	ret0, _ := ret[0].(sqlc.GetRentalByIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRentalByID indicates an expected call of GetRentalByID.
func (mr *MockQuerierMockRecorder) GetRentalByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRentalByID", reflect.TypeOf((*MockQuerier)(nil).GetRentalByID), arg0, arg1)
}

// GetRentals mocks base method.
func (m *MockQuerier) GetRentals(arg0 context.Context, arg1 sqlc.GetRentalsParams) ([]sqlc.GetRentalsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRentals", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.GetRentalsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRentals indicates an expected call of GetRentals.
func (mr *MockQuerierMockRecorder) GetRentals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRentals", reflect.TypeOf((*MockQuerier)(nil).GetRentals), arg0, arg1)
}
