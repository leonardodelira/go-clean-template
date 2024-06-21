// Code generated by MockGen. DO NOT EDIT.
// Source: ./services.go
//
// Generated by this command:
//
//	mockgen -source=./services.go -package=mockups -destination=../../../mocks/mockups/services/services.go
//

// Package mockups is a generated GoMock package.
package mockups

import (
	context "context"
	domain "leonardodelira/go-clean-template/internal/core/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTranslationService is a mock of TranslationService interface.
type MockTranslationService struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationServiceMockRecorder
}

// MockTranslationServiceMockRecorder is the mock recorder for MockTranslationService.
type MockTranslationServiceMockRecorder struct {
	mock *MockTranslationService
}

// NewMockTranslationService creates a new mock instance.
func NewMockTranslationService(ctrl *gomock.Controller) *MockTranslationService {
	mock := &MockTranslationService{ctrl: ctrl}
	mock.recorder = &MockTranslationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslationService) EXPECT() *MockTranslationServiceMockRecorder {
	return m.recorder
}

// DoTranslation mocks base method.
func (m *MockTranslationService) DoTranslation(ctx context.Context, input domain.TranslationInput) (*domain.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoTranslation", ctx, input)
	ret0, _ := ret[0].(*domain.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoTranslation indicates an expected call of DoTranslation.
func (mr *MockTranslationServiceMockRecorder) DoTranslation(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoTranslation", reflect.TypeOf((*MockTranslationService)(nil).DoTranslation), ctx, input)
}

// GetTranslation mocks base method.
func (m *MockTranslationService) GetTranslation(ctx context.Context) ([]domain.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTranslation", ctx)
	ret0, _ := ret[0].([]domain.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTranslation indicates an expected call of GetTranslation.
func (mr *MockTranslationServiceMockRecorder) GetTranslation(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTranslation", reflect.TypeOf((*MockTranslationService)(nil).GetTranslation), ctx)
}
