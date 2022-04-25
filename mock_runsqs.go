// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/github.com/asecurityteam/runsqs/v2/interface.go

// Package stat is a generated GoMock package.
package stat

import (
	context "context"
	reflect "reflect"

	runsqs "github.com/asecurityteam/runsqs/v2"
	sqs "github.com/aws/aws-sdk-go/service/sqs"
	gomock "github.com/golang/mock/gomock"
)

// MockSQSConsumer is a mock of SQSConsumer interface.
type MockSQSConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockSQSConsumerMockRecorder
}

// MockSQSConsumerMockRecorder is the mock recorder for MockSQSConsumer.
type MockSQSConsumerMockRecorder struct {
	mock *MockSQSConsumer
}

// NewMockSQSConsumer creates a new mock instance.
func NewMockSQSConsumer(ctrl *gomock.Controller) *MockSQSConsumer {
	mock := &MockSQSConsumer{ctrl: ctrl}
	mock.recorder = &MockSQSConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQSConsumer) EXPECT() *MockSQSConsumerMockRecorder {
	return m.recorder
}

// GetSQSMessageConsumer mocks base method.
func (m *MockSQSConsumer) GetSQSMessageConsumer() runsqs.SQSMessageConsumer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSQSMessageConsumer")
	ret0, _ := ret[0].(runsqs.SQSMessageConsumer)
	return ret0
}

// GetSQSMessageConsumer indicates an expected call of GetSQSMessageConsumer.
func (mr *MockSQSConsumerMockRecorder) GetSQSMessageConsumer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSQSMessageConsumer", reflect.TypeOf((*MockSQSConsumer)(nil).GetSQSMessageConsumer))
}

// StartConsuming mocks base method.
func (m *MockSQSConsumer) StartConsuming(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartConsuming", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartConsuming indicates an expected call of StartConsuming.
func (mr *MockSQSConsumerMockRecorder) StartConsuming(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConsuming", reflect.TypeOf((*MockSQSConsumer)(nil).StartConsuming), ctx)
}

// StopConsuming mocks base method.
func (m *MockSQSConsumer) StopConsuming(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopConsuming", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopConsuming indicates an expected call of StopConsuming.
func (mr *MockSQSConsumerMockRecorder) StopConsuming(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopConsuming", reflect.TypeOf((*MockSQSConsumer)(nil).StopConsuming), ctx)
}

// MockSQSMessageConsumer is a mock of SQSMessageConsumer interface.
type MockSQSMessageConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockSQSMessageConsumerMockRecorder
}

// MockSQSMessageConsumerMockRecorder is the mock recorder for MockSQSMessageConsumer.
type MockSQSMessageConsumerMockRecorder struct {
	mock *MockSQSMessageConsumer
}

// NewMockSQSMessageConsumer creates a new mock instance.
func NewMockSQSMessageConsumer(ctrl *gomock.Controller) *MockSQSMessageConsumer {
	mock := &MockSQSMessageConsumer{ctrl: ctrl}
	mock.recorder = &MockSQSMessageConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQSMessageConsumer) EXPECT() *MockSQSMessageConsumerMockRecorder {
	return m.recorder
}

// ConsumeMessage mocks base method.
func (m *MockSQSMessageConsumer) ConsumeMessage(ctx context.Context, message *sqs.Message) runsqs.SQSMessageConsumerError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeMessage", ctx, message)
	ret0, _ := ret[0].(runsqs.SQSMessageConsumerError)
	return ret0
}

// ConsumeMessage indicates an expected call of ConsumeMessage.
func (mr *MockSQSMessageConsumerMockRecorder) ConsumeMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeMessage", reflect.TypeOf((*MockSQSMessageConsumer)(nil).ConsumeMessage), ctx, message)
}

// DeadLetter mocks base method.
func (m *MockSQSMessageConsumer) DeadLetter(ctx context.Context, message *sqs.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeadLetter", ctx, message)
}

// DeadLetter indicates an expected call of DeadLetter.
func (mr *MockSQSMessageConsumerMockRecorder) DeadLetter(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeadLetter", reflect.TypeOf((*MockSQSMessageConsumer)(nil).DeadLetter), ctx, message)
}

// MockSQSMessageConsumerError is a mock of SQSMessageConsumerError interface.
type MockSQSMessageConsumerError struct {
	ctrl     *gomock.Controller
	recorder *MockSQSMessageConsumerErrorMockRecorder
}

// MockSQSMessageConsumerErrorMockRecorder is the mock recorder for MockSQSMessageConsumerError.
type MockSQSMessageConsumerErrorMockRecorder struct {
	mock *MockSQSMessageConsumerError
}

// NewMockSQSMessageConsumerError creates a new mock instance.
func NewMockSQSMessageConsumerError(ctrl *gomock.Controller) *MockSQSMessageConsumerError {
	mock := &MockSQSMessageConsumerError{ctrl: ctrl}
	mock.recorder = &MockSQSMessageConsumerErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQSMessageConsumerError) EXPECT() *MockSQSMessageConsumerErrorMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockSQSMessageConsumerError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockSQSMessageConsumerErrorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockSQSMessageConsumerError)(nil).Error))
}

// IsRetryable mocks base method.
func (m *MockSQSMessageConsumerError) IsRetryable() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRetryable")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRetryable indicates an expected call of IsRetryable.
func (mr *MockSQSMessageConsumerErrorMockRecorder) IsRetryable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRetryable", reflect.TypeOf((*MockSQSMessageConsumerError)(nil).IsRetryable))
}

// RetryAfter mocks base method.
func (m *MockSQSMessageConsumerError) RetryAfter() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetryAfter")
	ret0, _ := ret[0].(int64)
	return ret0
}

// RetryAfter indicates an expected call of RetryAfter.
func (mr *MockSQSMessageConsumerErrorMockRecorder) RetryAfter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryAfter", reflect.TypeOf((*MockSQSMessageConsumerError)(nil).RetryAfter))
}

// MockSQSProducer is a mock of SQSProducer interface.
type MockSQSProducer struct {
	ctrl     *gomock.Controller
	recorder *MockSQSProducerMockRecorder
}

// MockSQSProducerMockRecorder is the mock recorder for MockSQSProducer.
type MockSQSProducerMockRecorder struct {
	mock *MockSQSProducer
}

// NewMockSQSProducer creates a new mock instance.
func NewMockSQSProducer(ctrl *gomock.Controller) *MockSQSProducer {
	mock := &MockSQSProducer{ctrl: ctrl}
	mock.recorder = &MockSQSProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQSProducer) EXPECT() *MockSQSProducerMockRecorder {
	return m.recorder
}

// ProduceMessage mocks base method.
func (m *MockSQSProducer) ProduceMessage(ctx context.Context, messageInput *sqs.SendMessageInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceMessage", ctx, messageInput)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceMessage indicates an expected call of ProduceMessage.
func (mr *MockSQSProducerMockRecorder) ProduceMessage(ctx, messageInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceMessage", reflect.TypeOf((*MockSQSProducer)(nil).ProduceMessage), ctx, messageInput)
}

// QueueURL mocks base method.
func (m *MockSQSProducer) QueueURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// QueueURL indicates an expected call of QueueURL.
func (mr *MockSQSProducerMockRecorder) QueueURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueURL", reflect.TypeOf((*MockSQSProducer)(nil).QueueURL))
}
