// Copyright 2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Source: dependencies.go in package engine
// Automatically generated by MockGen. DO NOT EDIT!

package engine

import (
	io "io"

	gomock "code.google.com/p/gomock/gomock"
)

// Mock of downloader interface
type Mockdownloader struct {
	ctrl     *gomock.Controller
	recorder *_MockdownloaderRecorder
}

// Recorder for Mockdownloader (not exported)
type _MockdownloaderRecorder struct {
	mock *Mockdownloader
}

func NewMockdownloader(ctrl *gomock.Controller) *Mockdownloader {
	mock := &Mockdownloader{ctrl: ctrl}
	mock.recorder = &_MockdownloaderRecorder{mock}
	return mock
}

func (_m *Mockdownloader) EXPECT() *_MockdownloaderRecorder {
	return _m.recorder
}

func (_m *Mockdownloader) IsAgentCached() bool {
	ret := _m.ctrl.Call(_m, "IsAgentCached")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdownloaderRecorder) IsAgentCached() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsAgentCached")
}

func (_m *Mockdownloader) IsAgentLatest() bool {
	ret := _m.ctrl.Call(_m, "IsAgentLatest")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdownloaderRecorder) IsAgentLatest() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsAgentLatest")
}

func (_m *Mockdownloader) DownloadAgent() error {
	ret := _m.ctrl.Call(_m, "DownloadAgent")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdownloaderRecorder) DownloadAgent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DownloadAgent")
}

func (_m *Mockdownloader) LoadCachedAgent() (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "LoadCachedAgent")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdownloaderRecorder) LoadCachedAgent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoadCachedAgent")
}

func (_m *Mockdownloader) LoadDesiredAgent() (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "LoadDesiredAgent")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdownloaderRecorder) LoadDesiredAgent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoadDesiredAgent")
}

// Mock of dockerClient interface
type MockdockerClient struct {
	ctrl     *gomock.Controller
	recorder *_MockdockerClientRecorder
}

// Recorder for MockdockerClient (not exported)
type _MockdockerClientRecorder struct {
	mock *MockdockerClient
}

func NewMockdockerClient(ctrl *gomock.Controller) *MockdockerClient {
	mock := &MockdockerClient{ctrl: ctrl}
	mock.recorder = &_MockdockerClientRecorder{mock}
	return mock
}

func (_m *MockdockerClient) EXPECT() *_MockdockerClientRecorder {
	return _m.recorder
}

func (_m *MockdockerClient) IsAgentImageLoaded() (bool, error) {
	ret := _m.ctrl.Call(_m, "IsAgentImageLoaded")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdockerClientRecorder) IsAgentImageLoaded() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsAgentImageLoaded")
}

func (_m *MockdockerClient) LoadImage(image io.Reader) error {
	ret := _m.ctrl.Call(_m, "LoadImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdockerClientRecorder) LoadImage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoadImage", arg0)
}

func (_m *MockdockerClient) RemoveExistingAgentContainer() error {
	ret := _m.ctrl.Call(_m, "RemoveExistingAgentContainer")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdockerClientRecorder) RemoveExistingAgentContainer() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveExistingAgentContainer")
}

func (_m *MockdockerClient) StartAgent() (int, error) {
	ret := _m.ctrl.Call(_m, "StartAgent")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdockerClientRecorder) StartAgent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartAgent")
}

func (_m *MockdockerClient) StopAgent() error {
	ret := _m.ctrl.Call(_m, "StopAgent")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdockerClientRecorder) StopAgent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopAgent")
}
