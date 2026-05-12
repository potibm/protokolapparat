package news

import (
	"errors"
)

type News struct {
    ID          int64        `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	IsUrgent    bool   `json:"is_urgent"`
	ExternalURL string `json:"external_url,omitempty"`
	IsHidden    bool   `json:"is_hidden"`
}

func (e News) Validate() error {
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
