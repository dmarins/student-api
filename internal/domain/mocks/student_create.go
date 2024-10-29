// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/usecases/student_create.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/usecases/student_create.go -destination=internal/domain/mocks/student_create.go -typed=true -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	dtos "github.com/dmarins/student-api/internal/domain/dtos"
	gomock "go.uber.org/mock/gomock"
)

// MockIStudentCreateUseCase is a mock of IStudentCreateUseCase interface.
type MockIStudentCreateUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIStudentCreateUseCaseMockRecorder
	isgomock struct{}
}

// MockIStudentCreateUseCaseMockRecorder is the mock recorder for MockIStudentCreateUseCase.
type MockIStudentCreateUseCaseMockRecorder struct {
	mock *MockIStudentCreateUseCase
}

// NewMockIStudentCreateUseCase creates a new mock instance.
func NewMockIStudentCreateUseCase(ctrl *gomock.Controller) *MockIStudentCreateUseCase {
	mock := &MockIStudentCreateUseCase{ctrl: ctrl}
	mock.recorder = &MockIStudentCreateUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStudentCreateUseCase) EXPECT() *MockIStudentCreateUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIStudentCreateUseCase) Execute(ctx context.Context, studentInput dtos.StudentInput) *dtos.Result {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, studentInput)
	ret0, _ := ret[0].(*dtos.Result)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockIStudentCreateUseCaseMockRecorder) Execute(ctx, studentInput any) *MockIStudentCreateUseCaseExecuteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIStudentCreateUseCase)(nil).Execute), ctx, studentInput)
	return &MockIStudentCreateUseCaseExecuteCall{Call: call}
}

// MockIStudentCreateUseCaseExecuteCall wrap *gomock.Call
type MockIStudentCreateUseCaseExecuteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStudentCreateUseCaseExecuteCall) Return(arg0 *dtos.Result) *MockIStudentCreateUseCaseExecuteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStudentCreateUseCaseExecuteCall) Do(f func(context.Context, dtos.StudentInput) *dtos.Result) *MockIStudentCreateUseCaseExecuteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStudentCreateUseCaseExecuteCall) DoAndReturn(f func(context.Context, dtos.StudentInput) *dtos.Result) *MockIStudentCreateUseCaseExecuteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}