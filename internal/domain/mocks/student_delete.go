// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/usecases/student_delete.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/usecases/student_delete.go -destination=internal/domain/mocks/student_delete.go -typed=true -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	dtos "github.com/dmarins/student-api/internal/domain/dtos"
	gomock "go.uber.org/mock/gomock"
)

// MockIStudentDeleteUseCase is a mock of IStudentDeleteUseCase interface.
type MockIStudentDeleteUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIStudentDeleteUseCaseMockRecorder
	isgomock struct{}
}

// MockIStudentDeleteUseCaseMockRecorder is the mock recorder for MockIStudentDeleteUseCase.
type MockIStudentDeleteUseCaseMockRecorder struct {
	mock *MockIStudentDeleteUseCase
}

// NewMockIStudentDeleteUseCase creates a new mock instance.
func NewMockIStudentDeleteUseCase(ctrl *gomock.Controller) *MockIStudentDeleteUseCase {
	mock := &MockIStudentDeleteUseCase{ctrl: ctrl}
	mock.recorder = &MockIStudentDeleteUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStudentDeleteUseCase) EXPECT() *MockIStudentDeleteUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIStudentDeleteUseCase) Execute(ctx context.Context, studentId string) *dtos.Result {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, studentId)
	ret0, _ := ret[0].(*dtos.Result)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockIStudentDeleteUseCaseMockRecorder) Execute(ctx, studentId any) *MockIStudentDeleteUseCaseExecuteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIStudentDeleteUseCase)(nil).Execute), ctx, studentId)
	return &MockIStudentDeleteUseCaseExecuteCall{Call: call}
}

// MockIStudentDeleteUseCaseExecuteCall wrap *gomock.Call
type MockIStudentDeleteUseCaseExecuteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStudentDeleteUseCaseExecuteCall) Return(arg0 *dtos.Result) *MockIStudentDeleteUseCaseExecuteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStudentDeleteUseCaseExecuteCall) Do(f func(context.Context, string) *dtos.Result) *MockIStudentDeleteUseCaseExecuteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStudentDeleteUseCaseExecuteCall) DoAndReturn(f func(context.Context, string) *dtos.Result) *MockIStudentDeleteUseCaseExecuteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
