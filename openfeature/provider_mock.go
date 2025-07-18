//go:build testtools

// Code generated by MockGen. DO NOT EDIT.
// Source: openfeature/provider.go
//
// Generated by this command:
//
//	mockgen -source=openfeature/provider.go -destination=openfeature/provider_mock.go -package=openfeature -build_constraint=testtools
//

// Package openfeature is a generated GoMock package.
package openfeature

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFeatureProvider is a mock of FeatureProvider interface.
type MockFeatureProvider struct {
	ctrl     *gomock.Controller
	recorder *MockFeatureProviderMockRecorder
	isgomock struct{}
}

// MockFeatureProviderMockRecorder is the mock recorder for MockFeatureProvider.
type MockFeatureProviderMockRecorder struct {
	mock *MockFeatureProvider
}

// NewMockFeatureProvider creates a new mock instance.
func NewMockFeatureProvider(ctrl *gomock.Controller) *MockFeatureProvider {
	mock := &MockFeatureProvider{ctrl: ctrl}
	mock.recorder = &MockFeatureProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeatureProvider) EXPECT() *MockFeatureProviderMockRecorder {
	return m.recorder
}

// BooleanEvaluation mocks base method.
func (m *MockFeatureProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, flatCtx FlattenedContext) BoolResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BooleanEvaluation", ctx, flag, defaultValue, flatCtx)
	ret0, _ := ret[0].(BoolResolutionDetail)
	return ret0
}

// BooleanEvaluation indicates an expected call of BooleanEvaluation.
func (mr *MockFeatureProviderMockRecorder) BooleanEvaluation(ctx, flag, defaultValue, flatCtx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BooleanEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).BooleanEvaluation), ctx, flag, defaultValue, flatCtx)
}

// FloatEvaluation mocks base method.
func (m *MockFeatureProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, flatCtx FlattenedContext) FloatResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FloatEvaluation", ctx, flag, defaultValue, flatCtx)
	ret0, _ := ret[0].(FloatResolutionDetail)
	return ret0
}

// FloatEvaluation indicates an expected call of FloatEvaluation.
func (mr *MockFeatureProviderMockRecorder) FloatEvaluation(ctx, flag, defaultValue, flatCtx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FloatEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).FloatEvaluation), ctx, flag, defaultValue, flatCtx)
}

// Hooks mocks base method.
func (m *MockFeatureProvider) Hooks() []Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockFeatureProviderMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockFeatureProvider)(nil).Hooks))
}

// IntEvaluation mocks base method.
func (m *MockFeatureProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, flatCtx FlattenedContext) IntResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntEvaluation", ctx, flag, defaultValue, flatCtx)
	ret0, _ := ret[0].(IntResolutionDetail)
	return ret0
}

// IntEvaluation indicates an expected call of IntEvaluation.
func (mr *MockFeatureProviderMockRecorder) IntEvaluation(ctx, flag, defaultValue, flatCtx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).IntEvaluation), ctx, flag, defaultValue, flatCtx)
}

// Metadata mocks base method.
func (m *MockFeatureProvider) Metadata() Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].(Metadata)
	return ret0
}

// Metadata indicates an expected call of Metadata.
func (mr *MockFeatureProviderMockRecorder) Metadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockFeatureProvider)(nil).Metadata))
}

// ObjectEvaluation mocks base method.
func (m *MockFeatureProvider) ObjectEvaluation(ctx context.Context, flag string, defaultValue any, flatCtx FlattenedContext) InterfaceResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectEvaluation", ctx, flag, defaultValue, flatCtx)
	ret0, _ := ret[0].(InterfaceResolutionDetail)
	return ret0
}

// ObjectEvaluation indicates an expected call of ObjectEvaluation.
func (mr *MockFeatureProviderMockRecorder) ObjectEvaluation(ctx, flag, defaultValue, flatCtx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).ObjectEvaluation), ctx, flag, defaultValue, flatCtx)
}

// StringEvaluation mocks base method.
func (m *MockFeatureProvider) StringEvaluation(ctx context.Context, flag, defaultValue string, flatCtx FlattenedContext) StringResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringEvaluation", ctx, flag, defaultValue, flatCtx)
	ret0, _ := ret[0].(StringResolutionDetail)
	return ret0
}

// StringEvaluation indicates an expected call of StringEvaluation.
func (mr *MockFeatureProviderMockRecorder) StringEvaluation(ctx, flag, defaultValue, flatCtx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).StringEvaluation), ctx, flag, defaultValue, flatCtx)
}

// MockStateHandler is a mock of StateHandler interface.
type MockStateHandler struct {
	ctrl     *gomock.Controller
	recorder *MockStateHandlerMockRecorder
	isgomock struct{}
}

// MockStateHandlerMockRecorder is the mock recorder for MockStateHandler.
type MockStateHandlerMockRecorder struct {
	mock *MockStateHandler
}

// NewMockStateHandler creates a new mock instance.
func NewMockStateHandler(ctrl *gomock.Controller) *MockStateHandler {
	mock := &MockStateHandler{ctrl: ctrl}
	mock.recorder = &MockStateHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateHandler) EXPECT() *MockStateHandlerMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockStateHandler) Init(evaluationContext EvaluationContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", evaluationContext)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockStateHandlerMockRecorder) Init(evaluationContext any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockStateHandler)(nil).Init), evaluationContext)
}

// Shutdown mocks base method.
func (m *MockStateHandler) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockStateHandlerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockStateHandler)(nil).Shutdown))
}

// MockTracker is a mock of Tracker interface.
type MockTracker struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerMockRecorder
	isgomock struct{}
}

// MockTrackerMockRecorder is the mock recorder for MockTracker.
type MockTrackerMockRecorder struct {
	mock *MockTracker
}

// NewMockTracker creates a new mock instance.
func NewMockTracker(ctrl *gomock.Controller) *MockTracker {
	mock := &MockTracker{ctrl: ctrl}
	mock.recorder = &MockTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTracker) EXPECT() *MockTrackerMockRecorder {
	return m.recorder
}

// Track mocks base method.
func (m *MockTracker) Track(ctx context.Context, trackingEventName string, evaluationContext EvaluationContext, details TrackingEventDetails) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Track", ctx, trackingEventName, evaluationContext, details)
}

// Track indicates an expected call of Track.
func (mr *MockTrackerMockRecorder) Track(ctx, trackingEventName, evaluationContext, details any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Track", reflect.TypeOf((*MockTracker)(nil).Track), ctx, trackingEventName, evaluationContext, details)
}

// MockEventHandler is a mock of EventHandler interface.
type MockEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerMockRecorder
	isgomock struct{}
}

// MockEventHandlerMockRecorder is the mock recorder for MockEventHandler.
type MockEventHandlerMockRecorder struct {
	mock *MockEventHandler
}

// NewMockEventHandler creates a new mock instance.
func NewMockEventHandler(ctrl *gomock.Controller) *MockEventHandler {
	mock := &MockEventHandler{ctrl: ctrl}
	mock.recorder = &MockEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandler) EXPECT() *MockEventHandlerMockRecorder {
	return m.recorder
}

// EventChannel mocks base method.
func (m *MockEventHandler) EventChannel() <-chan Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventChannel")
	ret0, _ := ret[0].(<-chan Event)
	return ret0
}

// EventChannel indicates an expected call of EventChannel.
func (mr *MockEventHandlerMockRecorder) EventChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventChannel", reflect.TypeOf((*MockEventHandler)(nil).EventChannel))
}
