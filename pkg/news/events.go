package news

import (
	"errors"
	"fmt"

	"github.com/potibm/protokolapparat/pkg/common"
)

const SchemaVersion = 1

type NewsSyncEvent struct {
    Version   int                     `json:"v"`
    Action    common.ActionType       `json:"action"`
    Timestamp int64                   `json:"timestamp"`
    Payload   []News      `json:"payload"`
}

type NewsEvent struct {
    Version   int                     `json:"v"`
    Action    common.ActionType       `json:"action"`
    Timestamp int64                   `json:"timestamp"`
    Payload   News           `json:"payload"`
}

func (e NewsSyncEvent) Validate() error {
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

func NewScheduleCreateEvent(entry News) NewsEvent {
    return newScheduleEvent(entry, common.ActionCreate)
}

func NewScheduleUpdateEvent(entry News) NewsEvent {
    return newScheduleEvent(entry, common.ActionUpdate)
}

func NewScheduleDeleteEvent(entryID int64) NewsEvent {
    return newScheduleEvent(News{ID: entryID}, common.ActionDelete)
}

func newScheduleEvent(entry News, action common.ActionType) NewsEvent {
    return NewsEvent{
        Version:   SchemaVersion,
        Action:    action,
        Timestamp: common.NowUnix(),
        Payload:   entry,
    }
}

func (e NewsEvent) Validate() error {
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

func NewScheduleSyncEvent(entries []News) NewsSyncEvent {
    return NewsSyncEvent{
        Version:   SchemaVersion,
        Action:    common.ActionSync,
        Timestamp: common.NowUnix(),
        Payload:   entries,
    }
}