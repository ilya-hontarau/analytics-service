// Code generated by MockGen. DO NOT EDIT.
// Source: internal/analytics/service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	analytics "github.com/illfate/analytics-service/internal/analytics"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// InsertEvents mocks base method.
func (m *MockRepository) InsertEvents(ctx context.Context, events []analytics.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertEvents", ctx, events)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertEvents indicates an expected call of InsertEvents.
func (mr *MockRepositoryMockRecorder) InsertEvents(ctx, events interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEvents", reflect.TypeOf((*MockRepository)(nil).InsertEvents), ctx, events)
}
