package schedule

import (
	"errors"
	"fmt"

	"github.com/potibm/protokolapparat/pkg/common"
)

const SchemaVersion = 1

type ScheduleSyncEvent struct {
    Version   int                     `json:"v"`
    Action    common.ActionType       `json:"action"`
    Timestamp int64                   `json:"timestamp"`
    Payload   []ScheduleEntry         `json:"payload"`
}

type ScheduleEvent struct {
    Version   int                     `json:"v"`
    Action    common.ActionType       `json:"action"`
    Timestamp int64                   `json:"timestamp"`
    Payload   ScheduleEntry           `json:"payload"`
}

func (e ScheduleSyncEvent) Validate() error {
    if e.Action != common.ActionSync {
        return errors.New("use ScheduleEvent for non-sync actions")
    }

    for _, entry := range e.Payload {   
        if err := entry.Validate(); err != nil {
            return fmt.Errorf("invalid entry in sync payload: %w", err)
        }
    }

    return nil
}

func NewScheduleCreateEvent(entry ScheduleEntry) ScheduleEvent {
    return newScheduleEvent(entry, common.ActionCreate)
}

func NewScheduleUpdateEvent(entry ScheduleEntry) ScheduleEvent {
    return newScheduleEvent(entry, common.ActionUpdate)
}

func NewScheduleDeleteEvent(entryID int64) ScheduleEvent {
    return newScheduleEvent(ScheduleEntry{ID: entryID}, common.ActionDelete)
}

func newScheduleEvent(entry ScheduleEntry, action common.ActionType) ScheduleEvent {
    return ScheduleEvent{
        Version:   SchemaVersion,
        Action:    action,
        Timestamp: common.NowUnix(),
        Payload:   entry,
    }
}

func (e ScheduleEvent) Validate() error {
    switch e.Action {
    case common.ActionDelete:
        if e.Payload.ID <= 0 {
            return errors.New("delete action requires an entry ID")
        }
    case common.ActionCreate, common.ActionUpdate:
        if err := e.Payload.Validate(); err != nil {
            return fmt.Errorf("invalid payload for %s: %w", e.Action, err)
        }
    default:
        return fmt.Errorf("unknown action: %s", e.Action)
    }
    return nil
}

func NewScheduleSyncEvent(entries []ScheduleEntry) ScheduleSyncEvent {
    return ScheduleSyncEvent{
        Version:   SchemaVersion,
        Action:    common.ActionSync,
        Timestamp: common.NowUnix(),
        Payload:   entries,
    }
}