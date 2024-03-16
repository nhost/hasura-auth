// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go
//
// Generated by this command:
//
//	mockgen -package mock -destination mock/validator.go --source=validator.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHIBPClient is a mock of HIBPClient interface.
type MockHIBPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHIBPClientMockRecorder
}

// MockHIBPClientMockRecorder is the mock recorder for MockHIBPClient.
type MockHIBPClientMockRecorder struct {
	mock *MockHIBPClient
}

// NewMockHIBPClient creates a new mock instance.
func NewMockHIBPClient(ctrl *gomock.Controller) *MockHIBPClient {
	mock := &MockHIBPClient{ctrl: ctrl}
	mock.recorder = &MockHIBPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHIBPClient) EXPECT() *MockHIBPClientMockRecorder {
	return m.recorder
}

// IsPasswordPwned mocks base method.
func (m *MockHIBPClient) IsPasswordPwned(ctx context.Context, password string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPasswordPwned", ctx, password)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPasswordPwned indicates an expected call of IsPasswordPwned.
func (mr *MockHIBPClientMockRecorder) IsPasswordPwned(ctx, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPasswordPwned", reflect.TypeOf((*MockHIBPClient)(nil).IsPasswordPwned), ctx, password)
}
