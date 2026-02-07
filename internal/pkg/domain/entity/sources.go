package entity

type Step int64

const (
	UnknownStep           Step = 0
	CreateSourceStep      Step = 1
	SourceNameRuStep      Step = 2
	SourceNameENGStep     Step = 3
	SourceAuthorsRUStep   Step = 4
	SourceYearStep        Step = 5
	SourceDescriptionStep Step = 6
	SourceDownloadURLStep Step = 7
	SourceReadyToSend     Step = 8 //Вот тут обновление CreatedAt
	SourceIsSentStep      Step = 9
)

// GetNextStep ...
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
		return SourceDownloadURLStep
	case SourceDownloadURLStep:
		return SourceReadyToSend
	case SourceReadyToSend:
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
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	Step        Step       `json:"step"`
	Type        SourceType `json:"type"`
	NameRU      string     `json:"name_ru"`
	NameENG     string     `json:"name_eng"`
	AuthorRU    string     `json:"author_ru"`
	Year        int64      `json:"year"`
	Description string     `json:"description"`
	DownloadURL string     `json:"download_url"`
	CreatedAt   string     `json:"created_at"`
	IsSent      int64      `json:"isSent"`
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

// GetDescription ...
func (s *Source) GetDescription() string {
	if s == nil {
		return ""
	}
	return s.Description
}

// GetDownloadURL ...
func (s *Source) GetDownloadURL() string {
	if s == nil {
		return ""
	}
	return s.DownloadURL
}

// GetCreatedAt ...
func (s *Source) GetCreatedAt() string {
	if s == nil {
		return ""
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
