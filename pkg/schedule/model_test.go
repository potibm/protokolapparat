package schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntry_Validate(t *testing.T) {
	validEntry := Entry{
		ID:        1,
		Title:     "Meeting",
		StartTime: "2024-01-01T10:00:00Z",
	}

	tests := []struct {
		name    string
		entry   Entry
		wantErr string
	}{
		{
			name:    "missing ID",
			entry:   Entry{Title: "T", StartTime: "2024-01-01T10:00:00Z"},
			wantErr: "entry ID is required",
		},
		{
			name:    "missing title",
			entry:   Entry{ID: 1, StartTime: "2024-01-01T10:00:00Z"},
			wantErr: "entry title cannot be empty",
		},
		{
			name:    "invalid start time",
			entry:   Entry{ID: 1, Title: "T", StartTime: "not-a-time"},
			wantErr: "invalid start_time format: parsing time \"not-a-time\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"not-a-time\" as \"2006\"",
		},
		{
			name:    "invalid end time",
			entry:   Entry{ID: 1, Title: "T", StartTime: "2024-01-01T10:00:00Z", EndTime: "bad"},
			wantErr: "invalid end_time format: parsing time \"bad\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"bad\" as \"2006\"",
		},
		{
			name:    "invalid category",
			entry:   Entry{ID: 1, Title: "T", StartTime: "2024-01-01T10:00:00Z", Category: &Category{Name: ""}},
			wantErr: "invalid category: category name cannot be empty",
		},
		{
			name:    "invalid location",
			entry:   Entry{ID: 1, Title: "T", StartTime: "2024-01-01T10:00:00Z", Location: &Location{Name: ""}},
			wantErr: "invalid location: location name cannot be empty",
		},
		{
			name:  "valid minimal",
			entry: validEntry,
		},
		{
			name: "valid with category and location",
			entry: Entry{
				ID:        1,
				Title:     "Meeting",
				StartTime: "2024-01-01T10:00:00Z",
				EndTime:   "2024-01-01T11:00:00Z",
				Category:  &Category{Name: "Work", Color: "#ff0000"},
				Location:  &Location{Name: "Office"},
			},
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

func TestCategory_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cat     Category
		wantErr string
	}{
		{
			name:    "missing name",
			cat:     Category{Color: "#fff"},
			wantErr: "category name cannot be empty",
		},
		{
			name:    "missing color",
			cat:     Category{Name: "Work"},
			wantErr: "category color cannot be empty",
		},
		{
			name: "valid",
			cat:  Category{Name: "Work", Color: "#ff0000"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cat.Validate()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLocation_Validate(t *testing.T) {
	tests := []struct {
		name    string
		loc     Location
		wantErr string
	}{
		{
			name:    "missing name",
			loc:     Location{Address: "123 Main St"},
			wantErr: "location name cannot be empty",
		},
		{
			name: "valid",
			loc:  Location{Name: "Office"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.loc.Validate()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
