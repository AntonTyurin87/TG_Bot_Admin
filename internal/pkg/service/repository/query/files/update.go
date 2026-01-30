package files

import "TG_Bot_Admin/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Files []*entity.File
}

// GetFiles ...
func (s *Update) GetFiles() []*entity.File {
	if s == nil {
		return nil
	}

	return s.Files
}
