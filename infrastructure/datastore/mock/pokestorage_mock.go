// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructure/datastore/pokestorage.go

// Package mock_datastore is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockPokeStorage is a mock of PokeStorage interface.
type MockPokeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockPokeStorageMockRecorder
}

// MockPokeStorageMockRecorder is the mock recorder for MockPokeStorage.
type MockPokeStorageMockRecorder struct {
	mock *MockPokeStorage
}

// NewMockPokeStorage creates a new mock instance.
func NewMockPokeStorage(ctrl *gomock.Controller) *MockPokeStorage {
	mock := &MockPokeStorage{ctrl: ctrl}
	mock.recorder = &MockPokeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPokeStorage) EXPECT() *MockPokeStorageMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockPokeStorage) FindAll() []*model.Pokemon {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*model.Pokemon)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockPokeStorageMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockPokeStorage)(nil).FindAll))
}

// FindAllWorkers mocks base method.
func (m *MockPokeStorage) FindAllWorkers(typeStr string, items, itemsPerWorker int) ([]*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllWorkers", typeStr, items, itemsPerWorker)
	ret0, _ := ret[0].([]*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllWorkers indicates an expected call of FindAllWorkers.
func (mr *MockPokeStorageMockRecorder) FindAllWorkers(typeStr, items, itemsPerWorker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllWorkers", reflect.TypeOf((*MockPokeStorage)(nil).FindAllWorkers), typeStr, items, itemsPerWorker)
}

// FindById mocks base method.
func (m *MockPokeStorage) FindById(id int) *model.Pokemon {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*model.Pokemon)
	return ret0
}

// FindById indicates an expected call of FindById.
func (mr *MockPokeStorageMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockPokeStorage)(nil).FindById), id)
}

// Save mocks base method.
func (m *MockPokeStorage) Save(pokemon *model.Pokemon) (*model.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", pokemon)
	ret0, _ := ret[0].(*model.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockPokeStorageMockRecorder) Save(pokemon interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPokeStorage)(nil).Save), pokemon)
}
