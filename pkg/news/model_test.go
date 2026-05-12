package news

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntry_Validate(t *testing.T) {
	tests := []struct {
		name    string
		entry   Entry
		wantErr string
	}{
		{
			name:    "missing ID",
			entry:   Entry{Title: "T", Body: "B"},
			wantErr: "entry ID is required",
		},
		{
			name:    "missing title",
			entry:   Entry{ID: 1, Body: "B"},
			wantErr: "entry title cannot be empty",
		},
		{
			name:    "missing body",
			entry:   Entry{ID: 1, Title: "T"},
			wantErr: "entry body cannot be empty",
		},
		{
			name:  "valid",
			entry: Entry{ID: 1, Title: "T", Body: "B"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.entry.Validate()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
