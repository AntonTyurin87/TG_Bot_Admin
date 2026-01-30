package entity

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Step int64

const (
	UnknownStep           Step = 0
	CreateSourceStep      Step = 1
	SourceNameRuStep      Step = 2
	SourceNameENGStep     Step = 3
	SourceAuthorsRUStep   Step = 4
	SourceYearStep        Step = 5
	SourceDescriptionStep Step = 6
	SourceLoadFileStep    Step = 7 //Вот тут обновление CreatedAt
	SourceIsSentStep      Step = 8
)

func (s Step) GetNextStep() Step {
	switch s {
	case CreateSourceStep:
		return SourceNameRuStep
	case SourceNameRuStep:
		return SourceNameENGStep
	case SourceNameENGStep:
		return SourceAuthorsRUStep
	case SourceAuthorsRUStep:
		return SourceYearStep
	case SourceYearStep:
		return SourceDescriptionStep
	case SourceDescriptionStep:
		return SourceLoadFileStep
	case SourceLoadFileStep:
		return SourceIsSentStep
	default:
		return UnknownStep
	}
}

type SourceType string

const (
	UnknownSourceType  SourceType = "Unknown"
	BookSourceType     SourceType = "Книга"
	ArticleSourceType  SourceType = "Статья"
	FragmentSourceType SourceType = "Фрагмент"
)

// Source ...
type Source struct {
	ID          int64                `json:"id"`
	UserID      int64                `json:"user_id"`
	Step        Step                 `json:"step"`
	Type        SourceType           `json:"type"`
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
func (s *Source) GetStep() Step {
	if s == nil {
		return 0
	}
	return s.Step
}

// GetType ...
func (s *Source) GetType() SourceType {
	if s == nil {
		return UnknownSourceType
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
