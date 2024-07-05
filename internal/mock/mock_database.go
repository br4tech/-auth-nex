// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/database.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIDatabase is a mock of IDatabase interface.
type MockIDatabase[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIDatabaseMockRecorder[T]
}

// MockIDatabaseMockRecorder is the mock recorder for MockIDatabase.
type MockIDatabaseMockRecorder[T any] struct {
	mock *MockIDatabase[T]
}

// NewMockIDatabase creates a new mock instance.
func NewMockIDatabase[T any](ctrl *gomock.Controller) *MockIDatabase[T] {
	mock := &MockIDatabase[T]{ctrl: ctrl}
	mock.recorder = &MockIDatabaseMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDatabase[T]) EXPECT() *MockIDatabaseMockRecorder[T] {
	return m.recorder
}

// Create mocks base method.
func (m *MockIDatabase[T]) Create(ctx context.Context, value ...interface{}) (*any, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range value {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIDatabaseMockRecorder[T]) Create(ctx interface{}, value ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, value...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIDatabase[T])(nil).Create), varargs...)
}

// FindOne mocks base method.
func (m *MockIDatabase[T]) FindOne(ctx context.Context, conds ...interface{}) (T, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockIDatabaseMockRecorder[T]) FindOne(ctx interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockIDatabase[T])(nil).FindOne), varargs...)
}
