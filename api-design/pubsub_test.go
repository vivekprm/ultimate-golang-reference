package pubsub_test

import (
	"log"
	"testing"
)

// publisher is an interface to allow this package to mock the
// pubsub package support.
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the
// pubsub package.
type mock struct{}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {
	log.Printf("MockImpl: Publishing key: %s with value: %v", key, v)
	return nil
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {
	log.Printf("MockImpl: Subscribing to key: %s", key)
	return nil
}

func TestPubsub(t *testing.T) {
	// Create a mock publisher
	mockPublisher := &mock{}

	// Test publishing a message
	err := mockPublisher.Publish("testKey", "testValue")
	if err != nil {
		t.Errorf("Failed to publish: %v", err)
	}

	// Test subscribing to a key
	err = mockPublisher.Subscribe("testKey")
	if err != nil {
		t.Errorf("Failed to subscribe: %v", err)
	}
}
