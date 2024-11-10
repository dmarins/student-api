// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repositories/student_repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/repositories/student_repository.go -destination=internal/domain/mocks/student_repository.go -typed=true -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "github.com/dmarins/student-api/internal/domain/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockIStudentRepository is a mock of IStudentRepository interface.
type MockIStudentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIStudentRepositoryMockRecorder
	isgomock struct{}
}

// MockIStudentRepositoryMockRecorder is the mock recorder for MockIStudentRepository.
type MockIStudentRepositoryMockRecorder struct {
	mock *MockIStudentRepository
}

// NewMockIStudentRepository creates a new mock instance.
func NewMockIStudentRepository(ctrl *gomock.Controller) *MockIStudentRepository {
	mock := &MockIStudentRepository{ctrl: ctrl}
	mock.recorder = &MockIStudentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStudentRepository) EXPECT() *MockIStudentRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIStudentRepository) Add(ctx context.Context, student *entities.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, student)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockIStudentRepositoryMockRecorder) Add(ctx, student any) *MockIStudentRepositoryAddCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIStudentRepository)(nil).Add), ctx, student)
	return &MockIStudentRepositoryAddCall{Call: call}
}

// MockIStudentRepositoryAddCall wrap *gomock.Call
type MockIStudentRepositoryAddCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStudentRepositoryAddCall) Return(arg0 error) *MockIStudentRepositoryAddCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStudentRepositoryAddCall) Do(f func(context.Context, *entities.Student) error) *MockIStudentRepositoryAddCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStudentRepositoryAddCall) DoAndReturn(f func(context.Context, *entities.Student) error) *MockIStudentRepositoryAddCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockIStudentRepository) Delete(ctx context.Context, studentId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, studentId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIStudentRepositoryMockRecorder) Delete(ctx, studentId any) *MockIStudentRepositoryDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIStudentRepository)(nil).Delete), ctx, studentId)
	return &MockIStudentRepositoryDeleteCall{Call: call}
}

// MockIStudentRepositoryDeleteCall wrap *gomock.Call
type MockIStudentRepositoryDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStudentRepositoryDeleteCall) Return(arg0 error) *MockIStudentRepositoryDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStudentRepositoryDeleteCall) Do(f func(context.Context, string) error) *MockIStudentRepositoryDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStudentRepositoryDeleteCall) DoAndReturn(f func(context.Context, string) error) *MockIStudentRepositoryDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ExistsByName mocks base method.
func (m *MockIStudentRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByName", ctx, name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByName indicates an expected call of ExistsByName.
func (mr *MockIStudentRepositoryMockRecorder) ExistsByName(ctx, name any) *MockIStudentRepositoryExistsByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByName", reflect.TypeOf((*MockIStudentRepository)(nil).ExistsByName), ctx, name)
	return &MockIStudentRepositoryExistsByNameCall{Call: call}
}

// MockIStudentRepositoryExistsByNameCall wrap *gomock.Call
type MockIStudentRepositoryExistsByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStudentRepositoryExistsByNameCall) Return(arg0 bool, arg1 error) *MockIStudentRepositoryExistsByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStudentRepositoryExistsByNameCall) Do(f func(context.Context, string) (bool, error)) *MockIStudentRepositoryExistsByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStudentRepositoryExistsByNameCall) DoAndReturn(f func(context.Context, string) (bool, error)) *MockIStudentRepositoryExistsByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FindById mocks base method.
func (m *MockIStudentRepository) FindById(ctx context.Context, studentId string) (*entities.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, studentId)
	ret0, _ := ret[0].(*entities.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockIStudentRepositoryMockRecorder) FindById(ctx, studentId any) *MockIStudentRepositoryFindByIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockIStudentRepository)(nil).FindById), ctx, studentId)
	return &MockIStudentRepositoryFindByIdCall{Call: call}
}

// MockIStudentRepositoryFindByIdCall wrap *gomock.Call
type MockIStudentRepositoryFindByIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStudentRepositoryFindByIdCall) Return(arg0 *entities.Student, arg1 error) *MockIStudentRepositoryFindByIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStudentRepositoryFindByIdCall) Do(f func(context.Context, string) (*entities.Student, error)) *MockIStudentRepositoryFindByIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStudentRepositoryFindByIdCall) DoAndReturn(f func(context.Context, string) (*entities.Student, error)) *MockIStudentRepositoryFindByIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m *MockIStudentRepository) Update(ctx context.Context, student *entities.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, student)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIStudentRepositoryMockRecorder) Update(ctx, student any) *MockIStudentRepositoryUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIStudentRepository)(nil).Update), ctx, student)
	return &MockIStudentRepositoryUpdateCall{Call: call}
}

// MockIStudentRepositoryUpdateCall wrap *gomock.Call
type MockIStudentRepositoryUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStudentRepositoryUpdateCall) Return(arg0 error) *MockIStudentRepositoryUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStudentRepositoryUpdateCall) Do(f func(context.Context, *entities.Student) error) *MockIStudentRepositoryUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStudentRepositoryUpdateCall) DoAndReturn(f func(context.Context, *entities.Student) error) *MockIStudentRepositoryUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
