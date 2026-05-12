package common

type ActionType string

const (
    ActionCreate ActionType = "create"
    ActionUpdate ActionType = "update"
    ActionDelete ActionType = "delete"
    ActionSync   ActionType = "sync"
)