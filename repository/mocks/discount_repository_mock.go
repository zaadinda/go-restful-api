// Code generated by MockGen. DO NOT EDIT.
// Source: repository/category_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/aronipurwanto/go-restful-api/model/domain"
	gomock "github.com/golang/mock/gomock"
)

type MockDiscountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDiscountRepositoryMockRecorder
}

type MockDiscountRepositoryMockRecorder struct {
	mock *MockDiscountRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockDiscountRepository(ctrl *gomock.Controller) *MockDiscountRepository {
	mock := &MockDiscountRepository{ctrl: ctrl}
	mock.recorder = &MockDiscountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiscountRepository) EXPECT() *MockDiscountRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDiscountRepository) Delete(ctx context.Context, category domain.Discount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDiscountRepositoryMockRecorder) Delete(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDiscountRepository)(nil).Delete), ctx, category)
}

// FindAll mocks base method.
func (m *MockDiscountRepository) FindAll(ctx context.Context) ([]domain.Discount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]domain.Discount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockDiscountRepositoryMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockDiscountRepository)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockDiscountRepository) FindById(ctx context.Context, categoryId uint64) (domain.Discount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, categoryId)
	ret0, _ := ret[0].(domain.Discount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockDiscountRepositoryMockRecorder) FindById(ctx, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockDiscountRepository)(nil).FindById), ctx, categoryId)
}

// Save mocks base method.
func (m *MockDiscountRepository) Save(ctx context.Context, category domain.Discount) (domain.Discount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, category)
	ret0, _ := ret[0].(domain.Discount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockDiscountRepositoryMockRecorder) Save(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDiscountRepository)(nil).Save), ctx, category)
}

// Update mocks base method.
func (m *MockDiscountRepository) Update(ctx context.Context, category domain.Discount) (domain.Discount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, category)
	ret0, _ := ret[0].(domain.Discount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDiscountRepositoryMockRecorder) Update(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDiscountRepository)(nil).Update), ctx, category)
}