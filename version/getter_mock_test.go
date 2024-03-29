// Code generated by MockGen. DO NOT EDIT.
// Source: internal/version/getter.go
//
// Generated by this command:
//
//	mockgen --source internal/version/getter.go --destination internal/version/getter_mock_test.go --package version_test --typed
//

// Package version_test is a generated GoMock package.
package version_test

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGetter is a mock of Getter interface.
type MockGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGetterMockRecorder
}

// MockGetterMockRecorder is the mock recorder for MockGetter.
type MockGetterMockRecorder struct {
	mock *MockGetter
}

// NewMockGetter creates a new mock instance.
func NewMockGetter(ctrl *gomock.Controller) *MockGetter {
	mock := &MockGetter{ctrl: ctrl}
	mock.recorder = &MockGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetter) EXPECT() *MockGetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetter) Get() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockGetterMockRecorder) Get() *MockGetterGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetter)(nil).Get))
	return &MockGetterGetCall{Call: call}
}

// MockGetterGetCall wrap *gomock.Call
type MockGetterGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGetterGetCall) Return(arg0 string) *MockGetterGetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGetterGetCall) Do(f func() string) *MockGetterGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGetterGetCall) DoAndReturn(f func() string) *MockGetterGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
