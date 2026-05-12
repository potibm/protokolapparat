// Package schedule provides the ScheduleEntry model and its associated event types.
package schedule

import (
	"errors"
	"fmt"
	"time"
)

// Entry represents a single entry in a schedule or timetable.
type Entry struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ExternalURL string    `json:"external_url"`
	StartTime   string    `json:"start_time"` // RFC3339
	EndTime     string    `json:"end_time"`   // RFC3339
	Hidden      bool      `json:"hidden"`
	Category    *Category `json:"category,omitempty"`
	Location    *Location `json:"location,omitempty"`
}

// Validate checks that the entry has a valid ID, title, and RFC3339 times.
func (e Entry) Validate() error {
	if e.ID <= 0 {
		return errors.New("entry ID is required")
	}
	if e.Title == "" {
		return errors.New("entry title cannot be empty")
	}

	if _, err := time.Parse(time.RFC3339, e.StartTime); err != nil {
		return fmt.Errorf("invalid start_time format: %w", err)
	}
	if e.EndTime != "" {
		if _, err := time.Parse(time.RFC3339, e.EndTime); err != nil {
			return fmt.Errorf("invalid end_time format: %w", err)
		}
	}

	if e.Category != nil {
		if err := e.Category.Validate(); err != nil {
			return fmt.Errorf("invalid category: %w", err)
		}
	}

	if e.Location != nil {
		if err := e.Location.Validate(); err != nil {
			return fmt.Errorf("invalid location: %w", err)
		}
	}

	return nil
}

// Category groups schedule entries by name and color.
type Category struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// Validate ensures the category has a non-empty name and color.
func (c Category) Validate() error {
	if c.Name == "" {
		return errors.New("category name cannot be empty")
	}
	if c.Color == "" {
		return errors.New("category color cannot be empty")
	}
	return nil
}

// Location represents a physical venue for a schedule entry.
type Location struct {
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
}

// Validate ensures the location has a non-empty name.
func (l Location) Validate() error {
	if l.Name == "" {
		return errors.New("location name cannot be empty")
	}
	return nil
}
