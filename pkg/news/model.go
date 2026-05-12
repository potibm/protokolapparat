// Package news provides the News model and its associated event types.
package news

import (
	"errors"
)

// Entry represents a news or announcement item.
type Entry struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	IsUrgent    bool   `json:"is_urgent"`
	ExternalURL string `json:"external_url,omitempty"`
	IsHidden    bool   `json:"is_hidden"`
}

// Validate checks that the News entry has a valid ID, title, and body.
func (e Entry) Validate() error {
	if e.ID <= 0 {
		return errors.New("entry ID is required")
	}
	if e.Title == "" {
		return errors.New("entry title cannot be empty")
	}
	if e.Body == "" {
		return errors.New("entry body cannot be empty")
	}

	return nil
}
