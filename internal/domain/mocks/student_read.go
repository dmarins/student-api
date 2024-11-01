// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/usecases/student_read.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/usecases/student_read.go -destination=internal/domain/mocks/student_read.go -typed=true -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	dtos "github.com/dmarins/student-api/internal/domain/dtos"
	gomock "go.uber.org/mock/gomock"
)

// MockIStudentReadUseCase is a mock of IStudentReadUseCase interface.
type MockIStudentReadUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIStudentReadUseCaseMockRecorder
	isgomock struct{}
}

// MockIStudentReadUseCaseMockRecorder is the mock recorder for MockIStudentReadUseCase.
type MockIStudentReadUseCaseMockRecorder struct {
	mock *MockIStudentReadUseCase
}

// NewMockIStudentReadUseCase creates a new mock instance.
func NewMockIStudentReadUseCase(ctrl *gomock.Controller) *MockIStudentReadUseCase {
	mock := &MockIStudentReadUseCase{ctrl: ctrl}
	mock.recorder = &MockIStudentReadUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStudentReadUseCase) EXPECT() *MockIStudentReadUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIStudentReadUseCase) Execute(ctx context.Context, studentId string) *dtos.Result {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, studentId)
	ret0, _ := ret[0].(*dtos.Result)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockIStudentReadUseCaseMockRecorder) Execute(ctx, studentId any) *MockIStudentReadUseCaseExecuteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIStudentReadUseCase)(nil).Execute), ctx, studentId)
	return &MockIStudentReadUseCaseExecuteCall{Call: call}
}

// MockIStudentReadUseCaseExecuteCall wrap *gomock.Call
type MockIStudentReadUseCaseExecuteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStudentReadUseCaseExecuteCall) Return(arg0 *dtos.Result) *MockIStudentReadUseCaseExecuteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStudentReadUseCaseExecuteCall) Do(f func(context.Context, string) *dtos.Result) *MockIStudentReadUseCaseExecuteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStudentReadUseCaseExecuteCall) DoAndReturn(f func(context.Context, string) *dtos.Result) *MockIStudentReadUseCaseExecuteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
