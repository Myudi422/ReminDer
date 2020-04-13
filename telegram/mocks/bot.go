// Code generated by MockGen. DO NOT EDIT.
// Source: bot.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	tbwrap "github.com/enrico5b1b4/tbwrap"
	gomock "github.com/golang/mock/gomock"
	telebot "gopkg.in/tucnak/telebot.v2"
)

// MockBot is a mock of Bot interface
type MockBot struct {
	ctrl     *gomock.Controller
	recorder *MockBotMockRecorder
}

// MockBotMockRecorder is the mock recorder for MockBot
type MockBotMockRecorder struct {
	mock *MockBot
}

// NewMockBot creates a new mock instance
func NewMockBot(ctrl *gomock.Controller) *MockBot {
	mock := &MockBot{ctrl: ctrl}
	mock.recorder = &MockBotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBot) EXPECT() *MockBotMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockBot) Handle(path string, handler tbwrap.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", path, handler)
}

// Handle indicates an expected call of Handle
func (mr *MockBotMockRecorder) Handle(path, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockBot)(nil).Handle), path, handler)
}

// Respond mocks base method
func (m *MockBot) Respond(callback *telebot.Callback, responseOptional ...*telebot.CallbackResponse) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{callback}
	for _, a := range responseOptional {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Respond", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Respond indicates an expected call of Respond
func (mr *MockBotMockRecorder) Respond(callback interface{}, responseOptional ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{callback}, responseOptional...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Respond", reflect.TypeOf((*MockBot)(nil).Respond), varargs...)
}

// Send mocks base method
func (m *MockBot) Send(to telebot.Recipient, what interface{}, options ...interface{}) (*telebot.Message, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{to, what}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*telebot.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockBotMockRecorder) Send(to, what interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{to, what}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBot)(nil).Send), varargs...)
}

// Start mocks base method
func (m *MockBot) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockBotMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBot)(nil).Start))
}