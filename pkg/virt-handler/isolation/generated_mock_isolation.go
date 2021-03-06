// Automatically generated by MockGen. DO NOT EDIT!
// Source: isolation.go

package isolation

import (
	gomock "github.com/golang/mock/gomock"

	v1 "kubevirt.io/client-go/api/v1"
)

// Mock of PodIsolationDetector interface
type MockPodIsolationDetector struct {
	ctrl     *gomock.Controller
	recorder *_MockPodIsolationDetectorRecorder
}

// Recorder for MockPodIsolationDetector (not exported)
type _MockPodIsolationDetectorRecorder struct {
	mock *MockPodIsolationDetector
}

func NewMockPodIsolationDetector(ctrl *gomock.Controller) *MockPodIsolationDetector {
	mock := &MockPodIsolationDetector{ctrl: ctrl}
	mock.recorder = &_MockPodIsolationDetectorRecorder{mock}
	return mock
}

func (_m *MockPodIsolationDetector) EXPECT() *_MockPodIsolationDetectorRecorder {
	return _m.recorder
}

func (_m *MockPodIsolationDetector) Detect(vm *v1.VirtualMachineInstance) (IsolationResult, error) {
	ret := _m.ctrl.Call(_m, "Detect", vm)
	ret0, _ := ret[0].(IsolationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPodIsolationDetectorRecorder) Detect(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Detect", arg0)
}

func (_m *MockPodIsolationDetector) DetectForSocket(vm *v1.VirtualMachineInstance, socket string) (IsolationResult, error) {
	ret := _m.ctrl.Call(_m, "DetectForSocket", vm, socket)
	ret0, _ := ret[0].(IsolationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPodIsolationDetectorRecorder) DetectForSocket(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DetectForSocket", arg0, arg1)
}

func (_m *MockPodIsolationDetector) Whitelist(controller []string) PodIsolationDetector {
	ret := _m.ctrl.Call(_m, "Whitelist", controller)
	ret0, _ := ret[0].(PodIsolationDetector)
	return ret0
}

func (_mr *_MockPodIsolationDetectorRecorder) Whitelist(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Whitelist", arg0)
}

func (_m *MockPodIsolationDetector) AdjustResources(vm *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "AdjustResources", vm)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPodIsolationDetectorRecorder) AdjustResources(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AdjustResources", arg0)
}

// Mock of IsolationResult interface
type MockIsolationResult struct {
	ctrl     *gomock.Controller
	recorder *_MockIsolationResultRecorder
}

// Recorder for MockIsolationResult (not exported)
type _MockIsolationResultRecorder struct {
	mock *MockIsolationResult
}

func NewMockIsolationResult(ctrl *gomock.Controller) *MockIsolationResult {
	mock := &MockIsolationResult{ctrl: ctrl}
	mock.recorder = &_MockIsolationResultRecorder{mock}
	return mock
}

func (_m *MockIsolationResult) EXPECT() *_MockIsolationResultRecorder {
	return _m.recorder
}

func (_m *MockIsolationResult) Slice() string {
	ret := _m.ctrl.Call(_m, "Slice")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) Slice() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Slice")
}

func (_m *MockIsolationResult) Pid() int {
	ret := _m.ctrl.Call(_m, "Pid")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) Pid() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Pid")
}

func (_m *MockIsolationResult) PIDNamespace() string {
	ret := _m.ctrl.Call(_m, "PIDNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) PIDNamespace() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PIDNamespace")
}

func (_m *MockIsolationResult) MountRoot() string {
	ret := _m.ctrl.Call(_m, "MountRoot")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) MountRoot() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MountRoot")
}

func (_m *MockIsolationResult) MountInfoRoot() (*MountInfo, error) {
	ret := _m.ctrl.Call(_m, "MountInfoRoot")
	ret0, _ := ret[0].(*MountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIsolationResultRecorder) MountInfoRoot() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MountInfoRoot")
}

func (_m *MockIsolationResult) MountNamespace() string {
	ret := _m.ctrl.Call(_m, "MountNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) MountNamespace() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MountNamespace")
}

func (_m *MockIsolationResult) NetNamespace() string {
	ret := _m.ctrl.Call(_m, "NetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) NetNamespace() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetNamespace")
}

func (_m *MockIsolationResult) DoNetNS(_param0 func() error) error {
	ret := _m.ctrl.Call(_m, "DoNetNS", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) DoNetNS(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoNetNS", arg0)
}
