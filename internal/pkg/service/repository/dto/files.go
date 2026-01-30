package dto

import "TG_Bot_Admin/internal/pkg/domain/entity"

const (
	FilesTableName = "files"

	FilesColumnID       = "id"
	FilesColumnSourceID = "source_id"
	FilesColumnFileData = "file_data"
)

var FilesColumns = []string{
	FilesColumnID,
	FilesColumnSourceID,
	FilesColumnFileData,
}

// File ...
type File struct {
	ID       int64  `json:"id"`
	SourceID int64  `json:"source_id"`
	FileData []byte `json:"file_data"`
}

// GetID ...
func (f *File) GetID() int64 {
	if f == nil {
		return 0
	}
	return f.ID
}

// GetSourceID ...
func (f *File) GetSourceID() int64 {
	if f == nil {
		return 0
	}
	return f.SourceID
}

// GetFileData ...
func (f *File) GetFileData() []byte {
	if f == nil {
		return nil
	}
	return f.FileData
}

// Entity ...
func (f *File) Entity() *entity.File {
	if f == nil {
		return nil
	}

	return &entity.File{
		ID:       f.GetID(),
		SourceID: f.GetSourceID(),
		FileData: f.GetFileData(),
	}
}

// Files ...
type Files []*File

// Entity ...
func (f Files) Entity() []*entity.File { return ToEntitySlice[[]*entity.File](f) }

// FileDtoFromEntity ...
func FileDtoFromEntity(e *entity.File) *File {
	if e == nil {
		return nil
	}
	return &File{
		ID:       e.GetID(),
		SourceID: e.GetSourceID(),
		FileData: e.GetFileData(),
	}
}
