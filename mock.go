package smskit

import "sync"

type MockClient struct {
	Sent []SentRecord
	mu   sync.Mutex
}

type SentRecord struct {
	From, To, Message string
}

func NewMockClient() *MockClient {
	return &MockClient{}
}
