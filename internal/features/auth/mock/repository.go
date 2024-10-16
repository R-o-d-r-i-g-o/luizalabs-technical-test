// Code generated by MockGen. DO NOT EDIT.
// Source: internal/features/auth/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	auth "luizalabs-technical-test/internal/features/auth"
	entity "luizalabs-technical-test/internal/pkg/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryImp is a mock of RepositoryImp interface.
type MockRepositoryImp struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryImpMockRecorder
}

// MockRepositoryImpMockRecorder is the mock recorder for MockRepositoryImp.
type MockRepositoryImpMockRecorder struct {
	mock *MockRepositoryImp
}

// NewMockRepositoryImp creates a new mock instance.
func NewMockRepositoryImp(ctrl *gomock.Controller) *MockRepositoryImp {
	mock := &MockRepositoryImp{ctrl: ctrl}
	mock.recorder = &MockRepositoryImpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryImp) EXPECT() *MockRepositoryImpMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockRepositoryImp) GetUser(filter auth.GetUserFilter) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", filter)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryImpMockRecorder) GetUser(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepositoryImp)(nil).GetUser), filter)
}

// RegisterUser mocks base method.
func (m *MockRepositoryImp) RegisterUser(user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockRepositoryImpMockRecorder) RegisterUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockRepositoryImp)(nil).RegisterUser), user)
}
