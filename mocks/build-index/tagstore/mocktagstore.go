// Copyright (c) 2016-2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/kraken/build-index/tagstore (interfaces: Store,FileStore)

// Package mocktagstore is a generated GoMock package.
package mocktagstore

import (
	gomock "github.com/golang/mock/gomock"
	core "github.com/uber/kraken/core"
	base "github.com/uber/kraken/lib/store/base"
	metadata "github.com/uber/kraken/lib/store/metadata"
	io "io"
	reflect "reflect"
	time "time"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockStore) Get(arg0 string) (core.Digest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(core.Digest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), arg0)
}

// Put mocks base method
func (m *MockStore) Put(arg0 string, arg1 core.Digest, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockStoreMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockStore)(nil).Put), arg0, arg1, arg2)
}

// MockFileStore is a mock of FileStore interface
type MockFileStore struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoreMockRecorder
}

// MockFileStoreMockRecorder is the mock recorder for MockFileStore
type MockFileStoreMockRecorder struct {
	mock *MockFileStore
}

// NewMockFileStore creates a new mock instance
func NewMockFileStore(ctrl *gomock.Controller) *MockFileStore {
	mock := &MockFileStore{ctrl: ctrl}
	mock.recorder = &MockFileStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileStore) EXPECT() *MockFileStoreMockRecorder {
	return m.recorder
}

// CreateCacheFile mocks base method
func (m *MockFileStore) CreateCacheFile(arg0 string, arg1 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCacheFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCacheFile indicates an expected call of CreateCacheFile
func (mr *MockFileStoreMockRecorder) CreateCacheFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCacheFile", reflect.TypeOf((*MockFileStore)(nil).CreateCacheFile), arg0, arg1)
}

// GetCacheFileReader mocks base method
func (m *MockFileStore) GetCacheFileReader(arg0 string) (base.FileReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCacheFileReader", arg0)
	ret0, _ := ret[0].(base.FileReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheFileReader indicates an expected call of GetCacheFileReader
func (mr *MockFileStoreMockRecorder) GetCacheFileReader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheFileReader", reflect.TypeOf((*MockFileStore)(nil).GetCacheFileReader), arg0)
}

// SetCacheFileMetadata mocks base method
func (m *MockFileStore) SetCacheFileMetadata(arg0 string, arg1 metadata.Metadata) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCacheFileMetadata", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCacheFileMetadata indicates an expected call of SetCacheFileMetadata
func (mr *MockFileStoreMockRecorder) SetCacheFileMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCacheFileMetadata", reflect.TypeOf((*MockFileStore)(nil).SetCacheFileMetadata), arg0, arg1)
}
