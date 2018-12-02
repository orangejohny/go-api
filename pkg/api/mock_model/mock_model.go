// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/orangejohny/SBWeb/pkg/model (interfaces: SM,DB)

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/orangejohny/SBWeb/pkg/model"
)

// MockSM is a mock of SM interface
type MockSM struct {
	ctrl     *gomock.Controller
	recorder *MockSMMockRecorder
}

// MockSMMockRecorder is the mock recorder for MockSM
type MockSMMockRecorder struct {
	mock *MockSM
}

// NewMockSM creates a new mock instance
func NewMockSM(ctrl *gomock.Controller) *MockSM {
	mock := &MockSM{ctrl: ctrl}
	mock.recorder = &MockSMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSM) EXPECT() *MockSMMockRecorder {
	return m.recorder
}

// CheckSession mocks base method
func (m *MockSM) CheckSession(arg0 *model.SessionID) (*model.Session, error) {
	ret := m.ctrl.Call(m, "CheckSession", arg0)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession
func (mr *MockSMMockRecorder) CheckSession(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockSM)(nil).CheckSession), arg0)
}

// CreateSession mocks base method
func (m *MockSM) CreateSession(arg0 *model.Session, arg1 bool) (*model.SessionID, error) {
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(*model.SessionID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockSMMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSM)(nil).CreateSession), arg0, arg1)
}

// DeleteSession mocks base method
func (m *MockSM) DeleteSession(arg0 *model.SessionID) error {
	ret := m.ctrl.Call(m, "DeleteSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession
func (mr *MockSMMockRecorder) DeleteSession(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSM)(nil).DeleteSession), arg0)
}

// MockDB is a mock of DB interface
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// EditAd mocks base method
func (m *MockDB) EditAd(arg0 *model.AdItem) (int64, error) {
	ret := m.ctrl.Call(m, "EditAd", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditAd indicates an expected call of EditAd
func (mr *MockDBMockRecorder) EditAd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAd", reflect.TypeOf((*MockDB)(nil).EditAd), arg0)
}

// EditUser mocks base method
func (m *MockDB) EditUser(arg0 *model.User) (int64, error) {
	ret := m.ctrl.Call(m, "EditUser", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditUser indicates an expected call of EditUser
func (mr *MockDBMockRecorder) EditUser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockDB)(nil).EditUser), arg0)
}

// GetAd mocks base method
func (m *MockDB) GetAd(arg0 int64) (*model.AdItem, error) {
	ret := m.ctrl.Call(m, "GetAd", arg0)
	ret0, _ := ret[0].(*model.AdItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAd indicates an expected call of GetAd
func (mr *MockDBMockRecorder) GetAd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAd", reflect.TypeOf((*MockDB)(nil).GetAd), arg0)
}

// GetAds mocks base method
func (m *MockDB) GetAds(arg0 *model.SearchParams) ([]*model.AdItem, error) {
	ret := m.ctrl.Call(m, "GetAds", arg0)
	ret0, _ := ret[0].([]*model.AdItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAds indicates an expected call of GetAds
func (mr *MockDBMockRecorder) GetAds(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAds", reflect.TypeOf((*MockDB)(nil).GetAds), arg0)
}

// GetAdsOfUser mocks base method
func (m *MockDB) GetAdsOfUser(arg0 int64) ([]*model.AdItem, error) {
	ret := m.ctrl.Call(m, "GetAdsOfUser", arg0)
	ret0, _ := ret[0].([]*model.AdItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdsOfUser indicates an expected call of GetAdsOfUser
func (mr *MockDBMockRecorder) GetAdsOfUser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdsOfUser", reflect.TypeOf((*MockDB)(nil).GetAdsOfUser), arg0)
}

// GetUserWithEmail mocks base method
func (m *MockDB) GetUserWithEmail(arg0 string) (*model.User, error) {
	ret := m.ctrl.Call(m, "GetUserWithEmail", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWithEmail indicates an expected call of GetUserWithEmail
func (mr *MockDBMockRecorder) GetUserWithEmail(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWithEmail", reflect.TypeOf((*MockDB)(nil).GetUserWithEmail), arg0)
}

// GetUserWithID mocks base method
func (m *MockDB) GetUserWithID(arg0 int64) (*model.User, error) {
	ret := m.ctrl.Call(m, "GetUserWithID", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWithID indicates an expected call of GetUserWithID
func (mr *MockDBMockRecorder) GetUserWithID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWithID", reflect.TypeOf((*MockDB)(nil).GetUserWithID), arg0)
}

// NewAd mocks base method
func (m *MockDB) NewAd(arg0 *model.AdItem) (int64, error) {
	ret := m.ctrl.Call(m, "NewAd", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAd indicates an expected call of NewAd
func (mr *MockDBMockRecorder) NewAd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAd", reflect.TypeOf((*MockDB)(nil).NewAd), arg0)
}

// NewUser mocks base method
func (m *MockDB) NewUser(arg0 *model.User) (int64, error) {
	ret := m.ctrl.Call(m, "NewUser", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUser indicates an expected call of NewUser
func (mr *MockDBMockRecorder) NewUser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUser", reflect.TypeOf((*MockDB)(nil).NewUser), arg0)
}

// RemoveAd mocks base method
func (m *MockDB) RemoveAd(arg0 int64) (int64, error) {
	ret := m.ctrl.Call(m, "RemoveAd", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAd indicates an expected call of RemoveAd
func (mr *MockDBMockRecorder) RemoveAd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAd", reflect.TypeOf((*MockDB)(nil).RemoveAd), arg0)
}

// RemoveUser mocks base method
func (m *MockDB) RemoveUser(arg0 int64) (int64, error) {
	ret := m.ctrl.Call(m, "RemoveUser", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveUser indicates an expected call of RemoveUser
func (mr *MockDBMockRecorder) RemoveUser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockDB)(nil).RemoveUser), arg0)
}
