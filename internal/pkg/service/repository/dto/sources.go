package dto

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"

	"github.com/golang/protobuf/ptypes/timestamp"
)

const (
	SourcesTableName = "sources"

	SourcesColumnID          = "id"
	SourcesColumnUserID      = "user_id"
	SourcesColumnStep        = "step"
	SourcesColumnType        = "type"
	SourcesColumnNameRU      = "name_ru"
	SourcesColumnNameENG     = "name_eng"
	SourcesColumnAuthorRU    = "author_ru"
	SourcesColumnYear        = "year"
	SourcesColumnRegions     = "regions"
	SourcesColumnTimePeriods = "time_periods"
	SourcesColumnDescription = "description"
	SourcesColumnFileFormat  = "file_format"
	SourcesColumnCreatedAt   = "created_at"
	SourcesColumnIsSent      = "isSent"
)

var SourcesColumns = []string{
	SourcesColumnID,
	SourcesColumnUserID,
	SourcesColumnStep,
	SourcesColumnType,
	SourcesColumnNameRU,
	SourcesColumnNameENG,
	SourcesColumnAuthorRU,
	SourcesColumnYear,
	SourcesColumnRegions,
	SourcesColumnTimePeriods,
	SourcesColumnDescription,
	SourcesColumnFileFormat,
	SourcesColumnCreatedAt,
	SourcesColumnIsSent,
}

// Source ...
type Source struct {
	ID          int64                `json:"id"`
	UserID      int64                `json:"user_id"`
	Step        int64                `json:"step"`
	Type        string               `json:"type"`
	NameRU      string               `json:"name_ru"`
	NameENG     string               `json:"name_eng"`
	AuthorRU    string               `json:"author_ru"`
	Year        int64                `json:"year"`
	Regions     string               `json:"regions"`
	TimePeriods string               `json:"time_periods"`
	Description string               `json:"description"`
	FileFormat  string               `json:"file_format"`
	CreatedAt   *timestamp.Timestamp `json:"created_at"`
	IsSent      int64                `json:"isSent"`
}

// GetID ...
func (s *Source) GetID() int64 {
	if s == nil {
		return 0
	}
	return s.ID
}

// GetUserID ...
func (s *Source) GetUserID() int64 {
	if s == nil {
		return 0
	}
	return s.UserID
}

// GetStep ...
func (s *Source) GetStep() int64 {
	if s == nil {
		return 0
	}
	return s.Step
}

// GetType ...
func (s *Source) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

// GetNameRU ...
func (s *Source) GetNameRU() string {
	if s == nil {
		return ""
	}
	return s.NameRU
}

// GetNameENG ...
func (s *Source) GetNameENG() string {
	if s == nil {
		return ""
	}
	return s.NameENG
}

// GetAuthorRU ...
func (s *Source) GetAuthorRU() string {
	if s == nil {
		return ""
	}
	return s.AuthorRU
}

// GetYear ...
func (s *Source) GetYear() int64 {
	if s == nil {
		return 0
	}
	return s.Year
}

// GetRegions ...
func (s *Source) GetRegions() string {
	if s == nil {
		return ""
	}
	return s.Regions
}

// GetTimePeriods ...
func (s *Source) GetTimePeriods() string {
	if s == nil {
		return ""
	}
	return s.TimePeriods
}

// GetDescription ...
func (s *Source) GetDescription() string {
	if s == nil {
		return ""
	}
	return s.Description
}

// GetFileFormat ...
func (s *Source) GetFileFormat() string {
	if s == nil {
		return ""
	}
	return s.FileFormat
}

// GetCreatedAt ...
func (s *Source) GetCreatedAt() *timestamp.Timestamp {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

// GetIsSent ...
func (s *Source) GetIsSent() int64 {
	if s == nil {
		return 0
	}
	return s.IsSent
}

// Entity ...
func (s *Source) Entity() *entity.Source {
	if s == nil {
		return nil
	}

	return &entity.Source{
		ID:          s.GetID(),
		UserID:      s.GetUserID(),
		Type:        entity.SourceType(s.GetType()),
		Step:        entity.Step(s.GetStep()),
		NameRU:      s.GetNameRU(),
		NameENG:     s.GetNameENG(),
		AuthorRU:    s.GetAuthorRU(),
		Year:        s.GetYear(),
		Regions:     s.GetRegions(),
		TimePeriods: s.GetTimePeriods(),
		Description: s.GetDescription(),
		FileFormat:  s.GetFileFormat(),
		CreatedAt:   s.GetCreatedAt(),
		IsSent:      s.GetIsSent(),
	}
}

// Sources ...
type Sources []*Source

// Entity ...
func (s Sources) Entity() []*entity.Source {
	return ToEntitySlice[[]*entity.Source](s)
}

// SourceDtoFromEntity ...
func SourceDtoFromEntity(e *entity.Source) *Source {
	if e == nil {
		return nil
	}
	return &Source{
		ID:          e.GetID(),
		UserID:      e.GetUserID(),
		Step:        int64(e.GetStep()),
		Type:        string(e.GetType()),
		NameRU:      e.GetNameRU(),
		NameENG:     e.GetNameENG(),
		AuthorRU:    e.GetAuthorRU(),
		Year:        e.GetYear(),
		Regions:     e.GetRegions(),
		TimePeriods: e.GetTimePeriods(),
		Description: e.GetDescription(),
		FileFormat:  e.GetFileFormat(),
		CreatedAt:   e.GetCreatedAt(),
		IsSent:      e.GetIsSent(),
	}
}
