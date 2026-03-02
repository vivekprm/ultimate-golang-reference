// Package pubsub simulates a package that provides
// publication/subscription type services.
package pubsub

import "log"

// PubSub provides access to a queue system.
type PubSub struct {
	key string
}

// New creates a pubsub value for use.
func New(key string) *PubSub {
	ps := PubSub{
		key: key,
	}

	return &ps
}

// Publish sends the data for the specified key.
func (ps *PubSub) Publish(key string, v interface{}) error {
	log.Printf("Publishing %v to %s\n", v, key)
	return nil
}

// Subscribe requests the data for the specified key.
func (ps *PubSub) Subscribe(key string) error {
	log.Printf("Subscribing to %s\n", key)
	return nil
}
