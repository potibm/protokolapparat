package common

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockValidatable is a test double that implements the Validatable interface.
type mockValidatable struct {
	valid bool
}

func (m mockValidatable) Validate() error {
	if !m.valid {
		return errors.New("mock validation failed")
	}
	return nil
}

func TestNewCreateEvent(t *testing.T) {
	item := mockValidatable{valid: true}
	evt := NewCreateEvent(item)

	assert.Equal(t, SchemaVersion, evt.Version)
	assert.Equal(t, ActionCreate, evt.Action)
	assert.NotZero(t, evt.Timestamp)
	assert.Equal(t, []mockValidatable{item}, evt.Payload)
}

func TestNewUpdateEvent(t *testing.T) {
	item := mockValidatable{valid: true}
	evt := NewUpdateEvent(item)

	assert.Equal(t, SchemaVersion, evt.Version)
	assert.Equal(t, ActionUpdate, evt.Action)
	assert.NotZero(t, evt.Timestamp)
	assert.Equal(t, []mockValidatable{item}, evt.Payload)
}

func TestNewDeleteEvent(t *testing.T) {
	item := mockValidatable{valid: true}
	evt := NewDeleteEvent(item)

	assert.Equal(t, SchemaVersion, evt.Version)
	assert.Equal(t, ActionDelete, evt.Action)
	assert.NotZero(t, evt.Timestamp)
	assert.Equal(t, []mockValidatable{item}, evt.Payload)
}

func TestNewSyncEvent(t *testing.T) {
	items := []mockValidatable{
		{valid: true},
		{valid: true},
	}
	evt := NewSyncEvent(items)

	assert.Equal(t, SchemaVersion, evt.Version)
	assert.Equal(t, ActionSync, evt.Action)
	assert.NotZero(t, evt.Timestamp)
	assert.Equal(t, items, evt.Payload)
}

func TestEvent_Validate(t *testing.T) {
	tests := []struct {
		name    string
		event   Event[mockValidatable]
		wantErr string
	}{
		{
			name:    "empty payload for create",
			event:   Event[mockValidatable]{Action: ActionCreate, Payload: []mockValidatable{}},
			wantErr: "payload cannot be empty for create action",
		},
		{
			name:    "empty payload for update",
			event:   Event[mockValidatable]{Action: ActionUpdate, Payload: []mockValidatable{}},
			wantErr: "payload cannot be empty for update action",
		},
		{
			name:    "empty payload for delete",
			event:   Event[mockValidatable]{Action: ActionDelete, Payload: []mockValidatable{}},
			wantErr: "payload cannot be empty for delete action",
		},
		{
			name:  "empty payload for sync",
			event: Event[mockValidatable]{Action: ActionSync, Payload: []mockValidatable{}},
		},
		{
			name:    "invalid payload entry for create",
			event:   Event[mockValidatable]{Action: ActionCreate, Payload: []mockValidatable{{valid: false}}},
			wantErr: "invalid entry in create payload: mock validation failed",
		},
		{
			name:    "invalid payload entry for sync",
			event:   Event[mockValidatable]{Action: ActionSync, Payload: []mockValidatable{{valid: false}}},
			wantErr: "invalid entry in sync payload: mock validation failed",
		},
		{
			name:  "delete skips payload validation",
			event: Event[mockValidatable]{Action: ActionDelete, Payload: []mockValidatable{{valid: false}}},
		},
		{
			name:  "valid create",
			event: Event[mockValidatable]{Action: ActionCreate, Payload: []mockValidatable{{valid: true}}},
		},
		{
			name:  "valid sync",
			event: Event[mockValidatable]{Action: ActionSync, Payload: []mockValidatable{{valid: true}, {valid: true}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.event.Validate()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
