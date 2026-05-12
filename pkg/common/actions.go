// Package common provides shared types and utilities used across domain packages.
package common

// ActionType represents the kind of operation performed on an entity.
type ActionType string

const (
	// ActionCreate indicates a new entity was created.
	ActionCreate ActionType = "create"
	// ActionUpdate indicates an existing entity was modified.
	ActionUpdate ActionType = "update"
	// ActionDelete indicates an entity was removed.
	ActionDelete ActionType = "delete"
	// ActionSync indicates a full-sync payload.
	ActionSync ActionType = "sync"
)
