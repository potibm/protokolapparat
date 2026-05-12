package common

import (
	"fmt"
	"time"
)

// SchemaVersion is the global version for all protocol events.
const SchemaVersion = 1

// Validatable ensures that payloads can validate their own business rules.
type Validatable interface {
	Validate() error
}

// Event is the generic envelope for all protocol messages.
// T represents the domain payload (e.g., schedule.Entry or news.Entry).
type Event[T Validatable] struct {
	Version   int        `json:"v"`
	Action    ActionType `json:"action"`
	Timestamp int64      `json:"timestamp"`
	Payload   []T        `json:"payload"`
}

// NewCreateEvent creates an event for a new item.
func NewCreateEvent[T Validatable](item T) Event[T] {
	return newEvent([]T{item}, ActionCreate)
}

// NewUpdateEvent creates an event for an updated item.
func NewUpdateEvent[T Validatable](item T) Event[T] {
	return newEvent([]T{item}, ActionUpdate)
}

// NewDeleteEvent creates an event to delete an item.
// Note: The item should only contain the ID.
func NewDeleteEvent[T Validatable](item T) Event[T] {
	return newEvent([]T{item}, ActionDelete)
}

// NewSyncEvent creates an event for a full synchronization.
func NewSyncEvent[T Validatable](items []T) Event[T] {
	return newEvent(items, ActionSync)
}

// newEvent is the internal helper for building the envelope.
func newEvent[T Validatable](items []T, action ActionType) Event[T] {
	return Event[T]{
		Version:   SchemaVersion,
		Action:    action,
		Timestamp: time.Now().Unix(),
		Payload:   items,
	}
}

// Validate checks the event action and triggers the payload's validation.
func (e Event[T]) Validate() error {
	if len(e.Payload) == 0 && e.Action != ActionSync {
		return fmt.Errorf("payload cannot be empty for %s action", e.Action)
	}

	if e.Action == ActionDelete {
		return nil
	}

	for _, item := range e.Payload {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("invalid entry in %s payload: %w", e.Action, err)
		}
	}

	return nil
}
