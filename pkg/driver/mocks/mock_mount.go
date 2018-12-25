// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/kubernetes/pkg/util/mount (interfaces: Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mount "k8s.io/kubernetes/pkg/util/mount"
	os "os"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CleanSubPaths mocks base method
func (m *MockInterface) CleanSubPaths(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanSubPaths", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanSubPaths indicates an expected call of CleanSubPaths
func (mr *MockInterfaceMockRecorder) CleanSubPaths(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanSubPaths", reflect.TypeOf((*MockInterface)(nil).CleanSubPaths), arg0, arg1)
}

// DeviceOpened mocks base method
func (m *MockInterface) DeviceOpened(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeviceOpened", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeviceOpened indicates an expected call of DeviceOpened
func (mr *MockInterfaceMockRecorder) DeviceOpened(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeviceOpened", reflect.TypeOf((*MockInterface)(nil).DeviceOpened), arg0)
}

// EvalHostSymlinks mocks base method
func (m *MockInterface) EvalHostSymlinks(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvalHostSymlinks", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvalHostSymlinks indicates an expected call of EvalHostSymlinks
func (mr *MockInterfaceMockRecorder) EvalHostSymlinks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalHostSymlinks", reflect.TypeOf((*MockInterface)(nil).EvalHostSymlinks), arg0)
}

// ExistsPath mocks base method
func (m *MockInterface) ExistsPath(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsPath", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsPath indicates an expected call of ExistsPath
func (mr *MockInterfaceMockRecorder) ExistsPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsPath", reflect.TypeOf((*MockInterface)(nil).ExistsPath), arg0)
}

// GetDeviceNameFromMount mocks base method
func (m *MockInterface) GetDeviceNameFromMount(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceNameFromMount", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceNameFromMount indicates an expected call of GetDeviceNameFromMount
func (mr *MockInterfaceMockRecorder) GetDeviceNameFromMount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceNameFromMount", reflect.TypeOf((*MockInterface)(nil).GetDeviceNameFromMount), arg0, arg1)
}

// GetFSGroup mocks base method
func (m *MockInterface) GetFSGroup(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFSGroup", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFSGroup indicates an expected call of GetFSGroup
func (mr *MockInterfaceMockRecorder) GetFSGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFSGroup", reflect.TypeOf((*MockInterface)(nil).GetFSGroup), arg0)
}

// GetFileType mocks base method
func (m *MockInterface) GetFileType(arg0 string) (mount.FileType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileType", arg0)
	ret0, _ := ret[0].(mount.FileType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileType indicates an expected call of GetFileType
func (mr *MockInterfaceMockRecorder) GetFileType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileType", reflect.TypeOf((*MockInterface)(nil).GetFileType), arg0)
}

// GetMode mocks base method
func (m *MockInterface) GetMode(arg0 string) (os.FileMode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMode", arg0)
	ret0, _ := ret[0].(os.FileMode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMode indicates an expected call of GetMode
func (mr *MockInterfaceMockRecorder) GetMode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMode", reflect.TypeOf((*MockInterface)(nil).GetMode), arg0)
}

// GetMountRefs mocks base method
func (m *MockInterface) GetMountRefs(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMountRefs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMountRefs indicates an expected call of GetMountRefs
func (mr *MockInterfaceMockRecorder) GetMountRefs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMountRefs", reflect.TypeOf((*MockInterface)(nil).GetMountRefs), arg0)
}

// GetSELinuxSupport mocks base method
func (m *MockInterface) GetSELinuxSupport(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSELinuxSupport", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSELinuxSupport indicates an expected call of GetSELinuxSupport
func (mr *MockInterfaceMockRecorder) GetSELinuxSupport(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSELinuxSupport", reflect.TypeOf((*MockInterface)(nil).GetSELinuxSupport), arg0)
}

// IsLikelyNotMountPoint mocks base method
func (m *MockInterface) IsLikelyNotMountPoint(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLikelyNotMountPoint", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLikelyNotMountPoint indicates an expected call of IsLikelyNotMountPoint
func (mr *MockInterfaceMockRecorder) IsLikelyNotMountPoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLikelyNotMountPoint", reflect.TypeOf((*MockInterface)(nil).IsLikelyNotMountPoint), arg0)
}

// IsMountPointMatch mocks base method
func (m *MockInterface) IsMountPointMatch(arg0 mount.MountPoint, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMountPointMatch", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMountPointMatch indicates an expected call of IsMountPointMatch
func (mr *MockInterfaceMockRecorder) IsMountPointMatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMountPointMatch", reflect.TypeOf((*MockInterface)(nil).IsMountPointMatch), arg0, arg1)
}

// IsNotMountPoint mocks base method
func (m *MockInterface) IsNotMountPoint(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNotMountPoint", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsNotMountPoint indicates an expected call of IsNotMountPoint
func (mr *MockInterfaceMockRecorder) IsNotMountPoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotMountPoint", reflect.TypeOf((*MockInterface)(nil).IsNotMountPoint), arg0)
}

// List mocks base method
func (m *MockInterface) List() ([]mount.MountPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]mount.MountPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockInterfaceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List))
}

// MakeDir mocks base method
func (m *MockInterface) MakeDir(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDir", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeDir indicates an expected call of MakeDir
func (mr *MockInterfaceMockRecorder) MakeDir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDir", reflect.TypeOf((*MockInterface)(nil).MakeDir), arg0)
}

// MakeFile mocks base method
func (m *MockInterface) MakeFile(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeFile indicates an expected call of MakeFile
func (mr *MockInterfaceMockRecorder) MakeFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeFile", reflect.TypeOf((*MockInterface)(nil).MakeFile), arg0)
}

// MakeRShared mocks base method
func (m *MockInterface) MakeRShared(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRShared", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeRShared indicates an expected call of MakeRShared
func (mr *MockInterfaceMockRecorder) MakeRShared(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRShared", reflect.TypeOf((*MockInterface)(nil).MakeRShared), arg0)
}

// Mount mocks base method
func (m *MockInterface) Mount(arg0, arg1, arg2 string, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount
func (mr *MockInterfaceMockRecorder) Mount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockInterface)(nil).Mount), arg0, arg1, arg2, arg3)
}

// PathIsDevice mocks base method
func (m *MockInterface) PathIsDevice(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PathIsDevice", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PathIsDevice indicates an expected call of PathIsDevice
func (mr *MockInterfaceMockRecorder) PathIsDevice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PathIsDevice", reflect.TypeOf((*MockInterface)(nil).PathIsDevice), arg0)
}

// PrepareSafeSubpath mocks base method
func (m *MockInterface) PrepareSafeSubpath(arg0 mount.Subpath) (string, func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareSafeSubpath", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(func())
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PrepareSafeSubpath indicates an expected call of PrepareSafeSubpath
func (mr *MockInterfaceMockRecorder) PrepareSafeSubpath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareSafeSubpath", reflect.TypeOf((*MockInterface)(nil).PrepareSafeSubpath), arg0)
}

// SafeMakeDir mocks base method
func (m *MockInterface) SafeMakeDir(arg0, arg1 string, arg2 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SafeMakeDir", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SafeMakeDir indicates an expected call of SafeMakeDir
func (mr *MockInterfaceMockRecorder) SafeMakeDir(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeMakeDir", reflect.TypeOf((*MockInterface)(nil).SafeMakeDir), arg0, arg1, arg2)
}

// Unmount mocks base method
func (m *MockInterface) Unmount(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount
func (mr *MockInterfaceMockRecorder) Unmount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockInterface)(nil).Unmount), arg0)
}
