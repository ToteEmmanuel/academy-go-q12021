// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/repository/pokerepository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockPokeRepository is a mock of PokeRepository interface.
type MockPokeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPokeRepositoryMockRecorder
}

// MockPokeRepositoryMockRecorder is the mock recorder for MockPokeRepository.
type MockPokeRepositoryMockRecorder struct {
	mock *MockPokeRepository
}

// NewMockPokeRepository creates a new mock instance.
func NewMockPokeRepository(ctrl *gomock.Controller) *MockPokeRepository {
	mock := &MockPokeRepository{ctrl: ctrl}
	mock.recorder = &MockPokeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPokeRepository) EXPECT() *MockPokeRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockPokeRepository) FindAll() ([]*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockPokeRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockPokeRepository)(nil).FindAll))
}

// FindAllWorkers mocks base method.
func (m *MockPokeRepository) FindAllWorkers(query string, items, worker int) ([]*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllWorkers", query, items, worker)
	ret0, _ := ret[0].([]*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllWorkers indicates an expected call of FindAllWorkers.
func (mr *MockPokeRepositoryMockRecorder) FindAllWorkers(query, items, worker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllWorkers", reflect.TypeOf((*MockPokeRepository)(nil).FindAllWorkers), query, items, worker)
}

// FindById mocks base method.
func (m *MockPokeRepository) FindById(id int) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockPokeRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockPokeRepository)(nil).FindById), id)
}

// Save mocks base method.
func (m *MockPokeRepository) Save(arg0 *model.Pokemon) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockPokeRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPokeRepository)(nil).Save), arg0)
}